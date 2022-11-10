// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: CommentBuilder)

package mocks

import (
	"reflect"
	"time"

	pegomock "github.com/petergtz/pegomock"
)

type MockCommentBuilder struct {
	fail func(message string, callerSkip ...int)
}

func NewMockCommentBuilder(options ...pegomock.Option) *MockCommentBuilder {
	mock := &MockCommentBuilder{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockCommentBuilder) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockCommentBuilder) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockCommentBuilder) BuildPlanComment(repoRelDir string, workspace string, project string, commentArgs []string) string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommentBuilder().")
	}
	params := []pegomock.Param{repoRelDir, workspace, project, commentArgs}
	result := pegomock.GetGenericMockFrom(mock).Invoke("BuildPlanComment", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockCommentBuilder) BuildApplyComment(repoRelDir string, workspace string, project string, autoMergeDisabled bool) string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommentBuilder().")
	}
	params := []pegomock.Param{repoRelDir, workspace, project, autoMergeDisabled}
	result := pegomock.GetGenericMockFrom(mock).Invoke("BuildApplyComment", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockCommentBuilder) BuildVersionComment(repoRelDir string, workspace string, project string) string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommentBuilder().")
	}
	params := []pegomock.Param{repoRelDir, workspace, project}
	result := pegomock.GetGenericMockFrom(mock).Invoke("BuildVersionComment", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockCommentBuilder) VerifyWasCalledOnce() *VerifierMockCommentBuilder {
	return &VerifierMockCommentBuilder{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockCommentBuilder) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockCommentBuilder {
	return &VerifierMockCommentBuilder{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockCommentBuilder) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockCommentBuilder {
	return &VerifierMockCommentBuilder{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockCommentBuilder) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockCommentBuilder {
	return &VerifierMockCommentBuilder{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockCommentBuilder struct {
	mock                   *MockCommentBuilder
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockCommentBuilder) BuildPlanComment(repoRelDir string, workspace string, project string, commentArgs []string) *MockCommentBuilder_BuildPlanComment_OngoingVerification {
	params := []pegomock.Param{repoRelDir, workspace, project, commentArgs}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "BuildPlanComment", params, verifier.timeout)
	return &MockCommentBuilder_BuildPlanComment_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommentBuilder_BuildPlanComment_OngoingVerification struct {
	mock              *MockCommentBuilder
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommentBuilder_BuildPlanComment_OngoingVerification) GetCapturedArguments() (string, string, string, []string) {
	repoRelDir, workspace, project, commentArgs := c.GetAllCapturedArguments()
	return repoRelDir[len(repoRelDir)-1], workspace[len(workspace)-1], project[len(project)-1], commentArgs[len(commentArgs)-1]
}

func (c *MockCommentBuilder_BuildPlanComment_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 []string, _param3 [][]string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
		_param3 = make([][]string, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.([]string)
		}
	}
	return
}

func (verifier *VerifierMockCommentBuilder) BuildApplyComment(repoRelDir string, workspace string, project string, autoMergeDisabled bool) *MockCommentBuilder_BuildApplyComment_OngoingVerification {
	params := []pegomock.Param{repoRelDir, workspace, project, autoMergeDisabled}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "BuildApplyComment", params, verifier.timeout)
	return &MockCommentBuilder_BuildApplyComment_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommentBuilder_BuildApplyComment_OngoingVerification struct {
	mock              *MockCommentBuilder
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommentBuilder_BuildApplyComment_OngoingVerification) GetCapturedArguments() (string, string, string, bool) {
	repoRelDir, workspace, project, autoMergeDisabled := c.GetAllCapturedArguments()
	return repoRelDir[len(repoRelDir)-1], workspace[len(workspace)-1], project[len(project)-1], autoMergeDisabled[len(autoMergeDisabled)-1]
}

func (c *MockCommentBuilder_BuildApplyComment_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 []string, _param3 []bool) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
		_param3 = make([]bool, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(bool)
		}
	}
	return
}

func (verifier *VerifierMockCommentBuilder) BuildVersionComment(repoRelDir string, workspace string, project string) *MockCommentBuilder_BuildVersionComment_OngoingVerification {
	params := []pegomock.Param{repoRelDir, workspace, project}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "BuildVersionComment", params, verifier.timeout)
	return &MockCommentBuilder_BuildVersionComment_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommentBuilder_BuildVersionComment_OngoingVerification struct {
	mock              *MockCommentBuilder
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommentBuilder_BuildVersionComment_OngoingVerification) GetCapturedArguments() (string, string, string) {
	repoRelDir, workspace, project := c.GetAllCapturedArguments()
	return repoRelDir[len(repoRelDir)-1], workspace[len(workspace)-1], project[len(project)-1]
}

func (c *MockCommentBuilder_BuildVersionComment_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
	}
	return
}
