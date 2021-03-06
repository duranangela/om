// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	io "io"
	sync "sync"
)

type ProgressBar struct {
	FinishStub        func()
	finishMutex       sync.RWMutex
	finishArgsForCall []struct {
	}
	NewProxyReaderStub        func(io.Reader) io.ReadCloser
	newProxyReaderMutex       sync.RWMutex
	newProxyReaderArgsForCall []struct {
		arg1 io.Reader
	}
	newProxyReaderReturns struct {
		result1 io.ReadCloser
	}
	newProxyReaderReturnsOnCall map[int]struct {
		result1 io.ReadCloser
	}
	ResetStub        func()
	resetMutex       sync.RWMutex
	resetArgsForCall []struct {
	}
	SetTotal64Stub        func(int64)
	setTotal64Mutex       sync.RWMutex
	setTotal64ArgsForCall []struct {
		arg1 int64
	}
	StartStub        func()
	startMutex       sync.RWMutex
	startArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ProgressBar) Finish() {
	fake.finishMutex.Lock()
	fake.finishArgsForCall = append(fake.finishArgsForCall, struct {
	}{})
	fake.recordInvocation("Finish", []interface{}{})
	fake.finishMutex.Unlock()
	if fake.FinishStub != nil {
		fake.FinishStub()
	}
}

func (fake *ProgressBar) FinishCallCount() int {
	fake.finishMutex.RLock()
	defer fake.finishMutex.RUnlock()
	return len(fake.finishArgsForCall)
}

func (fake *ProgressBar) FinishCalls(stub func()) {
	fake.finishMutex.Lock()
	defer fake.finishMutex.Unlock()
	fake.FinishStub = stub
}

func (fake *ProgressBar) NewProxyReader(arg1 io.Reader) io.ReadCloser {
	fake.newProxyReaderMutex.Lock()
	ret, specificReturn := fake.newProxyReaderReturnsOnCall[len(fake.newProxyReaderArgsForCall)]
	fake.newProxyReaderArgsForCall = append(fake.newProxyReaderArgsForCall, struct {
		arg1 io.Reader
	}{arg1})
	fake.recordInvocation("NewProxyReader", []interface{}{arg1})
	fake.newProxyReaderMutex.Unlock()
	if fake.NewProxyReaderStub != nil {
		return fake.NewProxyReaderStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.newProxyReaderReturns
	return fakeReturns.result1
}

func (fake *ProgressBar) NewProxyReaderCallCount() int {
	fake.newProxyReaderMutex.RLock()
	defer fake.newProxyReaderMutex.RUnlock()
	return len(fake.newProxyReaderArgsForCall)
}

func (fake *ProgressBar) NewProxyReaderCalls(stub func(io.Reader) io.ReadCloser) {
	fake.newProxyReaderMutex.Lock()
	defer fake.newProxyReaderMutex.Unlock()
	fake.NewProxyReaderStub = stub
}

func (fake *ProgressBar) NewProxyReaderArgsForCall(i int) io.Reader {
	fake.newProxyReaderMutex.RLock()
	defer fake.newProxyReaderMutex.RUnlock()
	argsForCall := fake.newProxyReaderArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ProgressBar) NewProxyReaderReturns(result1 io.ReadCloser) {
	fake.newProxyReaderMutex.Lock()
	defer fake.newProxyReaderMutex.Unlock()
	fake.NewProxyReaderStub = nil
	fake.newProxyReaderReturns = struct {
		result1 io.ReadCloser
	}{result1}
}

func (fake *ProgressBar) NewProxyReaderReturnsOnCall(i int, result1 io.ReadCloser) {
	fake.newProxyReaderMutex.Lock()
	defer fake.newProxyReaderMutex.Unlock()
	fake.NewProxyReaderStub = nil
	if fake.newProxyReaderReturnsOnCall == nil {
		fake.newProxyReaderReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
		})
	}
	fake.newProxyReaderReturnsOnCall[i] = struct {
		result1 io.ReadCloser
	}{result1}
}

func (fake *ProgressBar) Reset() {
	fake.resetMutex.Lock()
	fake.resetArgsForCall = append(fake.resetArgsForCall, struct {
	}{})
	fake.recordInvocation("Reset", []interface{}{})
	fake.resetMutex.Unlock()
	if fake.ResetStub != nil {
		fake.ResetStub()
	}
}

func (fake *ProgressBar) ResetCallCount() int {
	fake.resetMutex.RLock()
	defer fake.resetMutex.RUnlock()
	return len(fake.resetArgsForCall)
}

func (fake *ProgressBar) ResetCalls(stub func()) {
	fake.resetMutex.Lock()
	defer fake.resetMutex.Unlock()
	fake.ResetStub = stub
}

func (fake *ProgressBar) SetTotal64(arg1 int64) {
	fake.setTotal64Mutex.Lock()
	fake.setTotal64ArgsForCall = append(fake.setTotal64ArgsForCall, struct {
		arg1 int64
	}{arg1})
	fake.recordInvocation("SetTotal64", []interface{}{arg1})
	fake.setTotal64Mutex.Unlock()
	if fake.SetTotal64Stub != nil {
		fake.SetTotal64Stub(arg1)
	}
}

func (fake *ProgressBar) SetTotal64CallCount() int {
	fake.setTotal64Mutex.RLock()
	defer fake.setTotal64Mutex.RUnlock()
	return len(fake.setTotal64ArgsForCall)
}

func (fake *ProgressBar) SetTotal64Calls(stub func(int64)) {
	fake.setTotal64Mutex.Lock()
	defer fake.setTotal64Mutex.Unlock()
	fake.SetTotal64Stub = stub
}

func (fake *ProgressBar) SetTotal64ArgsForCall(i int) int64 {
	fake.setTotal64Mutex.RLock()
	defer fake.setTotal64Mutex.RUnlock()
	argsForCall := fake.setTotal64ArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ProgressBar) Start() {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
	}{})
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		fake.StartStub()
	}
}

func (fake *ProgressBar) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *ProgressBar) StartCalls(stub func()) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *ProgressBar) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.finishMutex.RLock()
	defer fake.finishMutex.RUnlock()
	fake.newProxyReaderMutex.RLock()
	defer fake.newProxyReaderMutex.RUnlock()
	fake.resetMutex.RLock()
	defer fake.resetMutex.RUnlock()
	fake.setTotal64Mutex.RLock()
	defer fake.setTotal64Mutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ProgressBar) recordInvocation(key string, args []interface{}) {
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
