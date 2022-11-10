// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/core/runtime (interfaces: PostWorkflowHookRunner)

package mocks

import (
	"reflect"
	"time"

	pegomock "github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
)

type MockPostWorkflowHookRunner struct {
	fail func(message string, callerSkip ...int)
}

func NewMockPostWorkflowHookRunner(options ...pegomock.Option) *MockPostWorkflowHookRunner {
	mock := &MockPostWorkflowHookRunner{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockPostWorkflowHookRunner) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockPostWorkflowHookRunner) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockPostWorkflowHookRunner) Run(ctx models.WorkflowHookCommandContext, command string, path string) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockPostWorkflowHookRunner().")
	}
	params := []pegomock.Param{ctx, command, path}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Run", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockPostWorkflowHookRunner) VerifyWasCalledOnce() *VerifierMockPostWorkflowHookRunner {
	return &VerifierMockPostWorkflowHookRunner{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockPostWorkflowHookRunner) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockPostWorkflowHookRunner {
	return &VerifierMockPostWorkflowHookRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockPostWorkflowHookRunner) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockPostWorkflowHookRunner {
	return &VerifierMockPostWorkflowHookRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockPostWorkflowHookRunner) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockPostWorkflowHookRunner {
	return &VerifierMockPostWorkflowHookRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockPostWorkflowHookRunner struct {
	mock                   *MockPostWorkflowHookRunner
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockPostWorkflowHookRunner) Run(ctx models.WorkflowHookCommandContext, command string, path string) *MockPostWorkflowHookRunner_Run_OngoingVerification {
	params := []pegomock.Param{ctx, command, path}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Run", params, verifier.timeout)
	return &MockPostWorkflowHookRunner_Run_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockPostWorkflowHookRunner_Run_OngoingVerification struct {
	mock              *MockPostWorkflowHookRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockPostWorkflowHookRunner_Run_OngoingVerification) GetCapturedArguments() (models.WorkflowHookCommandContext, string, string) {
	ctx, command, path := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1], command[len(command)-1], path[len(path)-1]
}

func (c *MockPostWorkflowHookRunner_Run_OngoingVerification) GetAllCapturedArguments() (_param0 []models.WorkflowHookCommandContext, _param1 []string, _param2 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.WorkflowHookCommandContext, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.WorkflowHookCommandContext)
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
