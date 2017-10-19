// Code generated by counterfeiter. DO NOT EDIT.
package wrapperfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/api/uaa"
	"code.cloudfoundry.org/cli/api/uaa/wrapper"
)

type FakeUAAClient struct {
	RefreshAccessTokenStub        func(refreshToken string) (uaa.RefreshedTokens, error)
	refreshAccessTokenMutex       sync.RWMutex
	refreshAccessTokenArgsForCall []struct {
		refreshToken string
	}
	refreshAccessTokenReturns struct {
		result1 uaa.RefreshedTokens
		result2 error
	}
	refreshAccessTokenReturnsOnCall map[int]struct {
		result1 uaa.RefreshedTokens
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUAAClient) RefreshAccessToken(refreshToken string) (uaa.RefreshedTokens, error) {
	fake.refreshAccessTokenMutex.Lock()
	ret, specificReturn := fake.refreshAccessTokenReturnsOnCall[len(fake.refreshAccessTokenArgsForCall)]
	fake.refreshAccessTokenArgsForCall = append(fake.refreshAccessTokenArgsForCall, struct {
		refreshToken string
	}{refreshToken})
	fake.recordInvocation("RefreshAccessToken", []interface{}{refreshToken})
	fake.refreshAccessTokenMutex.Unlock()
	if fake.RefreshAccessTokenStub != nil {
		return fake.RefreshAccessTokenStub(refreshToken)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.refreshAccessTokenReturns.result1, fake.refreshAccessTokenReturns.result2
}

func (fake *FakeUAAClient) RefreshAccessTokenCallCount() int {
	fake.refreshAccessTokenMutex.RLock()
	defer fake.refreshAccessTokenMutex.RUnlock()
	return len(fake.refreshAccessTokenArgsForCall)
}

func (fake *FakeUAAClient) RefreshAccessTokenArgsForCall(i int) string {
	fake.refreshAccessTokenMutex.RLock()
	defer fake.refreshAccessTokenMutex.RUnlock()
	return fake.refreshAccessTokenArgsForCall[i].refreshToken
}

func (fake *FakeUAAClient) RefreshAccessTokenReturns(result1 uaa.RefreshedTokens, result2 error) {
	fake.RefreshAccessTokenStub = nil
	fake.refreshAccessTokenReturns = struct {
		result1 uaa.RefreshedTokens
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) RefreshAccessTokenReturnsOnCall(i int, result1 uaa.RefreshedTokens, result2 error) {
	fake.RefreshAccessTokenStub = nil
	if fake.refreshAccessTokenReturnsOnCall == nil {
		fake.refreshAccessTokenReturnsOnCall = make(map[int]struct {
			result1 uaa.RefreshedTokens
			result2 error
		})
	}
	fake.refreshAccessTokenReturnsOnCall[i] = struct {
		result1 uaa.RefreshedTokens
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.refreshAccessTokenMutex.RLock()
	defer fake.refreshAccessTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUAAClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ wrapper.UAAClient = new(FakeUAAClient)
