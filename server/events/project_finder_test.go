// Copyright 2017 HootSuite Media Inc.
//
// Licensed under the Apache License, Version 2.0 (the License);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an AS IS BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Modified hereafter by contributors to runatlantis/atlantis.

package events_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/runatlantis/atlantis/server/core/config"
	"github.com/runatlantis/atlantis/server/core/config/valid"
	"github.com/runatlantis/atlantis/server/events"

	"github.com/runatlantis/atlantis/server/logging"
	. "github.com/runatlantis/atlantis/testing"
)

var modifiedRepo = "owner/repo"
var m = events.DefaultProjectFinder{}
var nestedModules1 string
var nestedModules2 string
var topLevelModules string
var envDir string

func setupTmpRepos(t *testing.T) {
	// Create different repo structures for testing.

	// 1. Nested modules directory inside a project
	// non-tf
	// terraform.tfstate
	// terraform.tfstate.backup
	// project1/
	//   main.tf
	//   terraform.tfstate
	//   terraform.tfstate.backup
	//   modules/
	//     main.tf
	var err error
	nestedModules1, err = os.MkdirTemp("", "")
	Ok(t, err)
	err = os.MkdirAll(filepath.Join(nestedModules1, "project1/modules"), 0700)
	Ok(t, err)
	files := []string{
		"non-tf",
		".tflint.hcl",
		"terraform.tfstate.backup",
		"project1/main.tf",
		"project1/terraform.tfstate",
		"project1/terraform.tfstate.backup",
		"project1/modules/main.tf",
	}
	for _, f := range files {
		_, err = os.Create(filepath.Join(nestedModules1, f))
		Ok(t, err)
	}

	// 2. Nested modules dir inside top-level project
	// main.tf
	//  modules/
	//    main.tf
	// We can just re-use part of the previous dir structure.
	nestedModules2 = filepath.Join(nestedModules1, "project1")

	// 3. Top-level modules
	//  modules/
	//    main.tf
	//  project1/
	//    main.tf
	//  project2/
	//    main.tf
	topLevelModules, err = os.MkdirTemp("", "")
	Ok(t, err)
	for _, path := range []string{"modules", "project1", "project2"} {
		err = os.MkdirAll(filepath.Join(topLevelModules, path), 0700)
		Ok(t, err)
		_, err = os.Create(filepath.Join(topLevelModules, path, "main.tf"))
		Ok(t, err)
	}

	// 4. Env/ dir
	// main.tf
	// env/
	//   staging.tfvars
	//   production.tfvars
	//   global-env-config.auto.tfvars.json
	envDir, err = os.MkdirTemp("", "")
	Ok(t, err)
	err = os.MkdirAll(filepath.Join(envDir, "env"), 0700)
	Ok(t, err)
	_, err = os.Create(filepath.Join(envDir, "env/staging.tfvars"))
	Ok(t, err)
	_, err = os.Create(filepath.Join(envDir, "env/production.tfvars"))
	Ok(t, err)
}

func TestDetermineProjects(t *testing.T) {
	noopLogger := logging.NewNoopLogger(t)
	setupTmpRepos(t)

	defaultAutoplanFileList := "**/*.tf,**/*.tfvars,**/*.tfvars.json,**/terragrunt.hcl,**/.terraform.lock.hcl"

	cases := []struct {
		description      string
		files            []string
		expProjectPaths  []string
		repoDir          string
		autoplanFileList string
	}{
		{
			"If no files were modified then should return an empty list",
			nil,
			nil,
			nestedModules1,
			defaultAutoplanFileList,
		},
		{
			"Should ignore non .tf files and return an empty list",
			[]string{"non-tf", "non.tf.suffix"},
			nil,
			nestedModules1,
			defaultAutoplanFileList,
		},
		{
			"Should ignore .tflint.hcl files and return an empty list",
			[]string{".tflint.hcl", "project1/.tflint.hcl"},
			nil,
			nestedModules1,
			defaultAutoplanFileList,
		},
		{
			"Should plan in the parent directory from modules if that dir has a main.tf",
			[]string{"project1/modules/main.tf"},
			[]string{"project1"},
			nestedModules1,
			defaultAutoplanFileList,
		},
		{
			"Should plan in the parent directory from modules if that dir has a main.tf",
			[]string{"modules/main.tf"},
			[]string{"."},
			nestedModules2,
			defaultAutoplanFileList,
		},
		{
			"Should plan in the parent directory from modules when module is in a subdir if that dir has a main.tf",
			[]string{"modules/subdir/main.tf"},
			[]string{"."},
			nestedModules2,
			defaultAutoplanFileList,
		},
		{
			"Should not plan in the parent directory from modules if that dir does not have a main.tf",
			[]string{"modules/main.tf"},
			[]string{},
			topLevelModules,
			defaultAutoplanFileList,
		},
		{
			"Should not plan in the parent directory from modules if that dir does not have a main.tf",
			[]string{"modules/main.tf", "project1/main.tf"},
			[]string{"project1"},
			topLevelModules,
			defaultAutoplanFileList,
		},
		{
			"Should ignore tfstate files and return an empty list",
			[]string{"terraform.tfstate", "terraform.tfstate.backup", "parent/terraform.tfstate", "parent/terraform.tfstate.backup"},
			nil,
			nestedModules1,
			defaultAutoplanFileList,
		},
		{
			"Should return '.' when changed file is at root",
			[]string{"a.tf"},
			[]string{"."},
			nestedModules2,
			defaultAutoplanFileList,
		},
		{
			"Should return directory when changed file is in a dir",
			[]string{"project1/a.tf"},
			[]string{"project1"},
			nestedModules1,
			defaultAutoplanFileList,
		},
		{
			"Should return parent dir when changed file is in an env/ dir",
			[]string{"env/staging.tfvars"},
			[]string{"."},
			envDir,
			defaultAutoplanFileList,
		},
		{
			"Should de-duplicate when multiple files changed in the same dir",
			[]string{"env/staging.tfvars", "main.tf", "other.tf"},
			[]string{"."},
			"",
			defaultAutoplanFileList,
		},
		{
			"Should ignore changes in a dir that was deleted",
			[]string{"wasdeleted/main.tf"},
			[]string{},
			"",
			defaultAutoplanFileList,
		},
		{
			"Should not ignore terragrunt.hcl files",
			[]string{"terragrunt.hcl"},
			[]string{"."},
			nestedModules2,
			defaultAutoplanFileList,
		},
		{
			"Should find terragrunt.hcl file inside a nested directory",
			[]string{"project1/terragrunt.hcl"},
			[]string{"project1"},
			nestedModules1,
			defaultAutoplanFileList,
		},
		{
			"Should find packer files and ignore default tf files",
			[]string{"project1/image.pkr.hcl", "project2/main.tf"},
			[]string{"project1"},
			topLevelModules,
			"**/*.pkr.hcl",
		},
		{
			"Should find yaml files in addition to defaults",
			[]string{"project1/ansible.yml", "project2/main.tf"},
			[]string{"project1", "project2"},
			topLevelModules,
			"**/*.tf,**/*.yml",
		},
		{
			"Should find yaml files unless excluded",
			[]string{"project1/ansible.yml", "project2/config.yml"},
			[]string{"project1"},
			topLevelModules,
			"**/*.yml,!project2/*.yml",
		},
		{
			"Should not ignore .terraform.lock.hcl files",
			[]string{"project1/.terraform.lock.hcl"},
			[]string{"project1"},
			nestedModules1,
			defaultAutoplanFileList,
		},
	}
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			projects := m.DetermineProjects(noopLogger, c.files, modifiedRepo, c.repoDir, c.autoplanFileList)

			// Extract the paths from the projects. We use a slice here instead of a
			// map so we can test whether there are duplicates returned.
			var paths []string
			for _, project := range projects {
				paths = append(paths, project.Path)
				// Check that the project object has the repo set properly.
				Equals(t, modifiedRepo, project.RepoFullName)
			}
			Assert(t, len(c.expProjectPaths) == len(paths),
				"exp %q but found %q", c.expProjectPaths, paths)

			for _, expPath := range c.expProjectPaths {
				found := false
				for _, actPath := range paths {
					if expPath == actPath {
						found = true
						break
					}
				}
				if !found {
					t.Fatalf("exp %q but was not in paths %v", expPath, paths)
				}
			}
		})
	}
}

func TestDefaultProjectFinder_DetermineProjectsViaConfig(t *testing.T) {
	// Create dir structure:
	// main.tf
	// project1/
	//   main.tf
	//   terraform.tfvars.json
	// project2/
	//   main.tf
	//   terraform.tfvars
	// modules/
	//   module/
	//	  main.tf
	tmpDir, cleanup := DirStructure(t, map[string]interface{}{
		"main.tf": nil,
		"project1": map[string]interface{}{
			"main.tf":               nil,
			"terraform.tfvars.json": nil,
		},
		"project2": map[string]interface{}{
			"main.tf":          nil,
			"terraform.tfvars": nil,
		},
		"modules": map[string]interface{}{
			"module": map[string]interface{}{
				"main.tf": nil,
			},
		},
	})
	defer cleanup()

	cases := []struct {
		description  string
		config       valid.RepoCfg
		modified     []string
		expProjPaths []string
	}{
		{
			// When autoplan is disabled, we still return the modified project.
			// If our caller is interested in autoplan enabled projects, they'll
			// need to filter the results.
			description: "autoplan disabled",
			config: valid.RepoCfg{
				Projects: []valid.Project{
					{
						Dir: ".",
						Autoplan: valid.Autoplan{
							Enabled:      false,
							WhenModified: []string{"**/*.tf"},
						},
					},
				},
			},
			modified:     []string{"main.tf"},
			expProjPaths: []string{"."},
		},
		{
			description: "autoplan default",
			config: valid.RepoCfg{
				Projects: []valid.Project{
					{
						Dir: ".",
						Autoplan: valid.Autoplan{
							Enabled:      true,
							WhenModified: []string{"**/*.tf"},
						},
					},
				},
			},
			modified:     []string{"main.tf"},
			expProjPaths: []string{"."},
		},
		{
			description: "parent dir modified",
			config: valid.RepoCfg{
				Projects: []valid.Project{
					{
						Dir: "project",
						Autoplan: valid.Autoplan{
							Enabled:      true,
							WhenModified: []string{"**/*.tf"},
						},
					},
				},
			},
			modified:     []string{"main.tf"},
			expProjPaths: nil,
		},
		{
			description: "parent dir modified matches",
			config: valid.RepoCfg{
				Projects: []valid.Project{
					{
						Dir: "project1",
						Autoplan: valid.Autoplan{
							Enabled:      true,
							WhenModified: []string{"../**/*.tf"},
						},
					},
				},
			},
			modified:     []string{"main.tf"},
			expProjPaths: []string{"project1"},
		},
		{
			description: "dir deleted",
			config: valid.RepoCfg{
				Projects: []valid.Project{
					{
						Dir: "project3",
						Autoplan: valid.Autoplan{
							Enabled:      true,
							WhenModified: []string{"*.tf"},
						},
					},
				},
			},
			modified:     []string{"project3/main.tf"},
			expProjPaths: nil,
		},
		{
			description: "multiple projects",
			config: valid.RepoCfg{
				Projects: []valid.Project{
					{
						Dir: ".",
						Autoplan: valid.Autoplan{
							Enabled:      true,
							WhenModified: []string{"*.tf"},
						},
					},
					{
						Dir: "project1",
						Autoplan: valid.Autoplan{
							Enabled:      true,
							WhenModified: []string{"../modules/module/*.tf", "**/*.tf"},
						},
					},
					{
						Dir: "project2",
						Autoplan: valid.Autoplan{
							Enabled:      true,
							WhenModified: []string{"**/*.tf"},
						},
					},
				},
			},
			modified:     []string{"main.tf", "modules/module/another.tf", "project2/nontf.txt"},
			expProjPaths: []string{".", "project1"},
		},
		{
			description: ".tfvars file modified",
			config: valid.RepoCfg{
				Projects: []valid.Project{
					{
						Dir: "project2",
						Autoplan: valid.Autoplan{
							Enabled:      true,
							WhenModified: []string{"*.tf*"},
						},
					},
				},
			},
			modified:     []string{"project2/terraform.tfvars"},
			expProjPaths: []string{"project2"},
		},
		{
			description: "file excluded",
			config: valid.RepoCfg{
				Projects: []valid.Project{
					{
						Dir: "project1",
						Autoplan: valid.Autoplan{
							Enabled:      true,
							WhenModified: []string{"*.tf", "!exclude-me.tf"},
						},
					},
				},
			},
			modified:     []string{"project1/exclude-me.tf"},
			expProjPaths: nil,
		},
		{
			description: "some files excluded and others included",
			config: valid.RepoCfg{
				Projects: []valid.Project{
					{
						Dir: "project1",
						Autoplan: valid.Autoplan{
							Enabled:      true,
							WhenModified: []string{"*.tf", "!exclude-me.tf"},
						},
					},
				},
			},
			modified:     []string{"project1/exclude-me.tf", "project1/include-me.tf"},
			expProjPaths: []string{"project1"},
		},
		{
			description: "multiple dirs excluded",
			config: valid.RepoCfg{
				Projects: []valid.Project{
					{
						Dir: "project1",
						Autoplan: valid.Autoplan{
							Enabled:      true,
							WhenModified: []string{"**/*.tf", "!subdir1/*", "!subdir2/*"},
						},
					},
				},
			},
			modified:     []string{"project1/subdir1/main.tf", "project1/subdir2/main.tf"},
			expProjPaths: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			pf := events.DefaultProjectFinder{}
			projects, err := pf.DetermineProjectsViaConfig(logging.NewNoopLogger(t), c.modified, c.config, tmpDir)
			Ok(t, err)
			Equals(t, len(c.expProjPaths), len(projects))
			for i, proj := range projects {
				Equals(t, c.expProjPaths[i], proj.Dir)
			}
		})
	}
}

func createTerragruntFiles(t *testing.T, files []string, destroyFlagSet bool) (string, func()) {
	content := ""

	if destroyFlagSet {
		content = "# ATLANTIS_PLEASE_DESTROY_STACK\n"
	}

	dirs := make(map[string]interface{})

	for _, file := range files {
		currentDirPtr := dirs

		folders := strings.Split(file, string(os.PathSeparator))

		for i, folder := range folders {
			if i == len(folders)-1 { // last item in folders is the filename
				currentDirPtr[folder] = content
				continue
			}

			if _, ok := currentDirPtr[folder]; !ok {
				currentDirPtr[folder] = make(map[string]interface{})
			}

			currentDirPtr = currentDirPtr[folder].(map[string]interface{})
		}
	}

	return DirStructure(t, dirs)
}

func createDirEnv(t *testing.T, files []string, atlantisYaml string, destroyFlagSet bool) (valid.RepoCfg, string, func()) {
	tmpDir, cleanup := createTerragruntFiles(t, files, destroyFlagSet)

	if atlantisYaml != "" {
		err := os.WriteFile(filepath.Join(tmpDir, config.AtlantisYAMLFilename), []byte(atlantisYaml), 0600)
		Ok(t, err)
	}
	r := config.ParserValidator{}
	var globalCfg = valid.NewGlobalCfg(true, false, false)
	config, err := r.ParseRepoCfg(tmpDir, globalCfg, "")
	Ok(t, err)

	return config, tmpDir, cleanup
}

func TestDefaultProjectFinder_DetermineProjectsViaConfig_FindProjectNo(t *testing.T) {
	noopLogger := logging.NewNoopLogger(t)
	cases := []struct {
		description  string
		AtlantisYAML string
		modified     []string
	}{
		{
			description: "test FindProjectNo",
			AtlantisYAML: `
version: 3
projects:
- autoplan:
    enabled: true
    when_modified:
    - '*.hcl'
    - '*.tf*'
  dir: dkprto/icawtopr/foo
- autoplan:
    enabled: true
    when_modified:
    - '*.hcl'
    - '*.tf*'
    - ../foo/terragrunt.hcl
  dir: dkprto/icawtopr/foo-attachment
`,
			modified: []string{"dkprto/icawtopr/foo/terragrunt.hcl", "dkprto/icawtopr/foo-attachment/terragrunt.hcl"},
		},
		{
			description: "test FindProjectNo with reverse atlantis config",
			AtlantisYAML: `
version: 3
projects:
- autoplan:
    enabled: true
    when_modified:
    - '*.hcl'
    - '*.tf*'
    - ../foo/terragrunt.hcl
  dir: dkprto/icawtopr/foo-attachment
- autoplan:
    enabled: true
    when_modified:
    - '*.hcl'
    - '*.tf*'
  dir: dkprto/icawtopr/foo
`,
			modified: []string{"dkprto/icawtopr/foo/terragrunt.hcl", "dkprto/icawtopr/foo-attachment/terragrunt.hcl"},
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			config, tmpDir, cleanup := createDirEnv(t, c.modified, c.AtlantisYAML, false)
			defer cleanup()

			pf := events.DefaultProjectFinder{}

			projects, err := pf.DetermineProjectsViaConfig(noopLogger, c.modified, config, tmpDir)
			Ok(t, err)

			for _, file := range c.modified {
				index, err := events.FindProjectNo(projects, file)
				Ok(t, err)
				Equals(t, projects[index].Dir, filepath.Dir(file))
			}
		})
	}
}

func TestDefaultProjectFinder_DetermineProjectsViaConfig_TestDependencyTracking(t *testing.T) {
	noopLogger := logging.NewNoopLogger(t)
	cases := []struct {
		description    string
		AtlantisYAML   string
		modified       []string
		expProjPaths   []string
		destroyFlagSet bool
		expError       string
	}{
		{
			description: "test dependency ordering",
			AtlantisYAML: `
version: 3
projects:
- autoplan:
    enabled: true
    when_modified:
    - '*.hcl'
    - '*.tf*'
    - ../4/terragrunt.hcl
  dir: dependency-test/1
- autoplan:
    enabled: true
    when_modified:
    - '*.hcl'
    - '*.tf*'
    - ../3/terragrunt.hcl
  dir: dependency-test/2
- autoplan:
    enabled: true
    when_modified:
    - '*.hcl'
    - '*.tf*'
    - ../1/terragrunt.hcl
  dir: dependency-test/3
- autoplan:
    enabled: true
    when_modified:
    - '*.hcl'
    - '*.tf*'
  dir: dependency-test/4
`,
			modified:       []string{"dependency-test/1/terragrunt.hcl", "dependency-test/2/terragrunt.hcl", "dependency-test/3/terragrunt.hcl", "dependency-test/4/terragrunt.hcl"},
			expProjPaths:   []string{"dependency-test/4", "dependency-test/1", "dependency-test/3", "dependency-test/2"},
			destroyFlagSet: false,
			expError:       "",
		},
		{
			description: "test reverted dependency ordering for destroy workflows",
			AtlantisYAML: `
version: 3
projects:
- autoplan:
    enabled: true
    when_modified:
    - '*.hcl'
    - '*.tf*'
    - ../4/terragrunt.hcl
  dir: dependency-test/1
- autoplan:
    enabled: true
    when_modified:
    - '*.hcl'
    - '*.tf*'
    - ../3/terragrunt.hcl
  dir: dependency-test/2
- autoplan:
    enabled: true
    when_modified:
    - '*.hcl'
    - '*.tf*'
    - ../1/terragrunt.hcl
  dir: dependency-test/3
- autoplan:
    enabled: true
    when_modified:
    - '*.hcl'
    - '*.tf*'
  dir: dependency-test/4
`,
			modified:       []string{"dependency-test/1/terragrunt.hcl", "dependency-test/2/terragrunt.hcl", "dependency-test/3/terragrunt.hcl", "dependency-test/4/terragrunt.hcl"},
			expProjPaths:   []string{"dependency-test/2", "dependency-test/3", "dependency-test/1", "dependency-test/4"},
			destroyFlagSet: true,
			expError:       "",
		},
		{
			description: "test dependency loop",
			AtlantisYAML: `
version: 3
projects:
- autoplan:
    enabled: true
    when_modified:
    - '*.hcl'
    - '*.tf*'
    - ../1/terragrunt.hcl
  dir: dependency-test/2
- autoplan:
    enabled: true
    when_modified:
    - '*.hcl'
    - '*.tf*'
    - ../2/terragrunt.hcl
  dir: dependency-test/3
- autoplan:
    enabled: true
    when_modified:
    - '*.hcl'
    - '*.tf*'
    - ../3/terragrunt.hcl
  dir: dependency-test/1
`,
			modified:       []string{"dependency-test/1/terragrunt.hcl", "dependency-test/2/terragrunt.hcl", "dependency-test/3/terragrunt.hcl"},
			expProjPaths:   []string{},
			destroyFlagSet: false,
			expError:       "topological sort failed",
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			config, tmpDir, cleanup := createDirEnv(t, c.modified, c.AtlantisYAML, c.destroyFlagSet)
			defer cleanup()

			pf := events.DefaultProjectFinder{}

			projects, err := pf.DetermineProjectsViaConfig(noopLogger, c.modified, config, tmpDir)
			if c.expError == "" {
				Ok(t, err)
				Equals(t, len(c.expProjPaths), len(projects))
				for i, proj := range projects {
					Equals(t, c.expProjPaths[i], proj.Dir)
				}
			} else {
				ErrContains(t, c.expError, err)
			}
		})
	}
}
