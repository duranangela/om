// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type JobsService struct {
	JobsStub        func(string) (map[string]string, error)
	jobsMutex       sync.RWMutex
	jobsArgsForCall []struct {
		arg1 string
	}
	jobsReturns struct {
		result1 map[string]string
		result2 error
	}
	jobsReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	GetExistingJobConfigStub        func(string, string) (api.JobProperties, error)
	getExistingJobConfigMutex       sync.RWMutex
	getExistingJobConfigArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getExistingJobConfigReturns struct {
		result1 api.JobProperties
		result2 error
	}
	getExistingJobConfigReturnsOnCall map[int]struct {
		result1 api.JobProperties
		result2 error
	}
	ConfigureJobStub        func(string, string, api.JobProperties) error
	configureJobMutex       sync.RWMutex
	configureJobArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 api.JobProperties
	}
	configureJobReturns struct {
		result1 error
	}
	configureJobReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *JobsService) Jobs(arg1 string) (map[string]string, error) {
	fake.jobsMutex.Lock()
	ret, specificReturn := fake.jobsReturnsOnCall[len(fake.jobsArgsForCall)]
	fake.jobsArgsForCall = append(fake.jobsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Jobs", []interface{}{arg1})
	fake.jobsMutex.Unlock()
	if fake.JobsStub != nil {
		return fake.JobsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.jobsReturns.result1, fake.jobsReturns.result2
}

func (fake *JobsService) JobsCallCount() int {
	fake.jobsMutex.RLock()
	defer fake.jobsMutex.RUnlock()
	return len(fake.jobsArgsForCall)
}

func (fake *JobsService) JobsArgsForCall(i int) string {
	fake.jobsMutex.RLock()
	defer fake.jobsMutex.RUnlock()
	return fake.jobsArgsForCall[i].arg1
}

func (fake *JobsService) JobsReturns(result1 map[string]string, result2 error) {
	fake.JobsStub = nil
	fake.jobsReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *JobsService) JobsReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.JobsStub = nil
	if fake.jobsReturnsOnCall == nil {
		fake.jobsReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.jobsReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *JobsService) GetExistingJobConfig(arg1 string, arg2 string) (api.JobProperties, error) {
	fake.getExistingJobConfigMutex.Lock()
	ret, specificReturn := fake.getExistingJobConfigReturnsOnCall[len(fake.getExistingJobConfigArgsForCall)]
	fake.getExistingJobConfigArgsForCall = append(fake.getExistingJobConfigArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetExistingJobConfig", []interface{}{arg1, arg2})
	fake.getExistingJobConfigMutex.Unlock()
	if fake.GetExistingJobConfigStub != nil {
		return fake.GetExistingJobConfigStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getExistingJobConfigReturns.result1, fake.getExistingJobConfigReturns.result2
}

func (fake *JobsService) GetExistingJobConfigCallCount() int {
	fake.getExistingJobConfigMutex.RLock()
	defer fake.getExistingJobConfigMutex.RUnlock()
	return len(fake.getExistingJobConfigArgsForCall)
}

func (fake *JobsService) GetExistingJobConfigArgsForCall(i int) (string, string) {
	fake.getExistingJobConfigMutex.RLock()
	defer fake.getExistingJobConfigMutex.RUnlock()
	return fake.getExistingJobConfigArgsForCall[i].arg1, fake.getExistingJobConfigArgsForCall[i].arg2
}

func (fake *JobsService) GetExistingJobConfigReturns(result1 api.JobProperties, result2 error) {
	fake.GetExistingJobConfigStub = nil
	fake.getExistingJobConfigReturns = struct {
		result1 api.JobProperties
		result2 error
	}{result1, result2}
}

func (fake *JobsService) GetExistingJobConfigReturnsOnCall(i int, result1 api.JobProperties, result2 error) {
	fake.GetExistingJobConfigStub = nil
	if fake.getExistingJobConfigReturnsOnCall == nil {
		fake.getExistingJobConfigReturnsOnCall = make(map[int]struct {
			result1 api.JobProperties
			result2 error
		})
	}
	fake.getExistingJobConfigReturnsOnCall[i] = struct {
		result1 api.JobProperties
		result2 error
	}{result1, result2}
}

func (fake *JobsService) ConfigureJob(arg1 string, arg2 string, arg3 api.JobProperties) error {
	fake.configureJobMutex.Lock()
	ret, specificReturn := fake.configureJobReturnsOnCall[len(fake.configureJobArgsForCall)]
	fake.configureJobArgsForCall = append(fake.configureJobArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 api.JobProperties
	}{arg1, arg2, arg3})
	fake.recordInvocation("ConfigureJob", []interface{}{arg1, arg2, arg3})
	fake.configureJobMutex.Unlock()
	if fake.ConfigureJobStub != nil {
		return fake.ConfigureJobStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.configureJobReturns.result1
}

func (fake *JobsService) ConfigureJobCallCount() int {
	fake.configureJobMutex.RLock()
	defer fake.configureJobMutex.RUnlock()
	return len(fake.configureJobArgsForCall)
}

func (fake *JobsService) ConfigureJobArgsForCall(i int) (string, string, api.JobProperties) {
	fake.configureJobMutex.RLock()
	defer fake.configureJobMutex.RUnlock()
	return fake.configureJobArgsForCall[i].arg1, fake.configureJobArgsForCall[i].arg2, fake.configureJobArgsForCall[i].arg3
}

func (fake *JobsService) ConfigureJobReturns(result1 error) {
	fake.ConfigureJobStub = nil
	fake.configureJobReturns = struct {
		result1 error
	}{result1}
}

func (fake *JobsService) ConfigureJobReturnsOnCall(i int, result1 error) {
	fake.ConfigureJobStub = nil
	if fake.configureJobReturnsOnCall == nil {
		fake.configureJobReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.configureJobReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *JobsService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.jobsMutex.RLock()
	defer fake.jobsMutex.RUnlock()
	fake.getExistingJobConfigMutex.RLock()
	defer fake.getExistingJobConfigMutex.RUnlock()
	fake.configureJobMutex.RLock()
	defer fake.configureJobMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *JobsService) recordInvocation(key string, args []interface{}) {
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