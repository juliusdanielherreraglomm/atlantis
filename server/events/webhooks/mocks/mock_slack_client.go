// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events/webhooks (interfaces: SlackClient)

package mocks

import (
	"reflect"
	"time"

	pegomock "github.com/petergtz/pegomock"
	webhooks "github.com/runatlantis/atlantis/server/events/webhooks"
)

type MockSlackClient struct {
	fail func(message string, callerSkip ...int)
}

func NewMockSlackClient(options ...pegomock.Option) *MockSlackClient {
	mock := &MockSlackClient{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockSlackClient) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockSlackClient) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockSlackClient) AuthTest() error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSlackClient().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("AuthTest", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockSlackClient) TokenIsSet() bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSlackClient().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("TokenIsSet", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockSlackClient) PostMessage(channel string, applyResult webhooks.ApplyResult) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSlackClient().")
	}
	params := []pegomock.Param{channel, applyResult}
	result := pegomock.GetGenericMockFrom(mock).Invoke("PostMessage", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockSlackClient) VerifyWasCalledOnce() *VerifierMockSlackClient {
	return &VerifierMockSlackClient{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockSlackClient) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockSlackClient {
	return &VerifierMockSlackClient{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockSlackClient) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockSlackClient {
	return &VerifierMockSlackClient{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockSlackClient) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockSlackClient {
	return &VerifierMockSlackClient{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockSlackClient struct {
	mock                   *MockSlackClient
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockSlackClient) AuthTest() *MockSlackClient_AuthTest_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "AuthTest", params, verifier.timeout)
	return &MockSlackClient_AuthTest_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSlackClient_AuthTest_OngoingVerification struct {
	mock              *MockSlackClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSlackClient_AuthTest_OngoingVerification) GetCapturedArguments() {
}

func (c *MockSlackClient_AuthTest_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockSlackClient) TokenIsSet() *MockSlackClient_TokenIsSet_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "TokenIsSet", params, verifier.timeout)
	return &MockSlackClient_TokenIsSet_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSlackClient_TokenIsSet_OngoingVerification struct {
	mock              *MockSlackClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSlackClient_TokenIsSet_OngoingVerification) GetCapturedArguments() {
}

func (c *MockSlackClient_TokenIsSet_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockSlackClient) PostMessage(channel string, applyResult webhooks.ApplyResult) *MockSlackClient_PostMessage_OngoingVerification {
	params := []pegomock.Param{channel, applyResult}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PostMessage", params, verifier.timeout)
	return &MockSlackClient_PostMessage_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSlackClient_PostMessage_OngoingVerification struct {
	mock              *MockSlackClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSlackClient_PostMessage_OngoingVerification) GetCapturedArguments() (string, webhooks.ApplyResult) {
	channel, applyResult := c.GetAllCapturedArguments()
	return channel[len(channel)-1], applyResult[len(applyResult)-1]
}

func (c *MockSlackClient_PostMessage_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []webhooks.ApplyResult) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]webhooks.ApplyResult, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(webhooks.ApplyResult)
		}
	}
	return
}
