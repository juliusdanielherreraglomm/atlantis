package testing

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

// DirStructure creates a directory structure in a temporary directory.
// structure describes the dir structure. If the value is another map, then the
// key is the name of a directory. If the value is nil, then the key is the name
// of a file. If val is a string then key is a file name and val is the file's content.
// It returns the path to the temp directory containing the defined
// structure.
// Example usage:
//
//		versionConfig := `
//	 terraform {
//		  required_version = "= 0.12.8"
//	 }
//	 `
//		tmpDir := DirStructure(t, map[string]interface{}{
//			"pulldir": map[string]interface{}{
//				"project1": map[string]interface{}{
//					"main.tf": nil,
//				},
//				"project2": map[string]interface{}{,
//					"main.tf": versionConfig,
//				},
//			},
//		})
func DirStructure(t *testing.T, structure map[string]interface{}) string {
	tmpDir := t.TempDir()
	dirStructureGo(t, tmpDir, structure)
	return tmpDir
}

func sortHash(h map[string]interface{}) []string {
	keys := make([]string, len(h))

	i := 0
	for k := range h {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	return keys
}

func dirStructureGo(t *testing.T, parentDir string, structure map[string]interface{}) {
	for _, key := range sortHash(structure) {
		// If structure[key] is nil then key is a filename and we just create it
		if structure[key] == nil {
			_, err := os.Create(filepath.Join(parentDir, key))
			Ok(t, err)
			continue
		}
		// If structure[key] is another map then key is a dir
		if dirContents, ok := structure[key].(map[string]interface{}); ok {
			subDir := filepath.Join(parentDir, key)
			Ok(t, os.Mkdir(subDir, 0700))
			// Recurse and create contents.
			dirStructureGo(t, subDir, dirContents)
		} else if fileContent, ok := structure[key].(string); ok {
			// If structure[key] is a string then key is a file name and structure[key] is the file's content
			err := ioutil.WriteFile(filepath.Join(parentDir, key), []byte(fileContent), 0600)
			Ok(t, err)
		}
	}
}
