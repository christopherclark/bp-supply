// This file was generated by counterfeiter
package servicebuilderfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/actors/servicebuilder"
	"code.cloudfoundry.org/cli/cf/models"
)

type FakeServiceBuilder struct {
	GetAllServicesStub        func() ([]models.ServiceOffering, error)
	getAllServicesMutex       sync.RWMutex
	getAllServicesArgsForCall []struct{}
	getAllServicesReturns     struct {
		result1 []models.ServiceOffering
		result2 error
	}
	GetAllServicesWithPlansStub        func() ([]models.ServiceOffering, error)
	getAllServicesWithPlansMutex       sync.RWMutex
	getAllServicesWithPlansArgsForCall []struct{}
	getAllServicesWithPlansReturns     struct {
		result1 []models.ServiceOffering
		result2 error
	}
	GetServiceByNameWithPlansStub        func(string) (models.ServiceOffering, error)
	getServiceByNameWithPlansMutex       sync.RWMutex
	getServiceByNameWithPlansArgsForCall []struct {
		arg1 string
	}
	getServiceByNameWithPlansReturns struct {
		result1 models.ServiceOffering
		result2 error
	}
	GetServiceByNameWithPlansWithOrgNamesStub        func(string) (models.ServiceOffering, error)
	getServiceByNameWithPlansWithOrgNamesMutex       sync.RWMutex
	getServiceByNameWithPlansWithOrgNamesArgsForCall []struct {
		arg1 string
	}
	getServiceByNameWithPlansWithOrgNamesReturns struct {
		result1 models.ServiceOffering
		result2 error
	}
	GetServiceByNameForSpaceStub        func(string, string) (models.ServiceOffering, error)
	getServiceByNameForSpaceMutex       sync.RWMutex
	getServiceByNameForSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getServiceByNameForSpaceReturns struct {
		result1 models.ServiceOffering
		result2 error
	}
	GetServiceByNameForSpaceWithPlansStub        func(string, string) (models.ServiceOffering, error)
	getServiceByNameForSpaceWithPlansMutex       sync.RWMutex
	getServiceByNameForSpaceWithPlansArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getServiceByNameForSpaceWithPlansReturns struct {
		result1 models.ServiceOffering
		result2 error
	}
	GetServicesByNameForSpaceWithPlansStub        func(string, string) (models.ServiceOfferings, error)
	getServicesByNameForSpaceWithPlansMutex       sync.RWMutex
	getServicesByNameForSpaceWithPlansArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getServicesByNameForSpaceWithPlansReturns struct {
		result1 models.ServiceOfferings
		result2 error
	}
	GetServiceByNameForOrgStub        func(string, string) (models.ServiceOffering, error)
	getServiceByNameForOrgMutex       sync.RWMutex
	getServiceByNameForOrgArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getServiceByNameForOrgReturns struct {
		result1 models.ServiceOffering
		result2 error
	}
	GetServicesForManyBrokersStub        func([]string) ([]models.ServiceOffering, error)
	getServicesForManyBrokersMutex       sync.RWMutex
	getServicesForManyBrokersArgsForCall []struct {
		arg1 []string
	}
	getServicesForManyBrokersReturns struct {
		result1 []models.ServiceOffering
		result2 error
	}
	GetServicesForBrokerStub        func(string) ([]models.ServiceOffering, error)
	getServicesForBrokerMutex       sync.RWMutex
	getServicesForBrokerArgsForCall []struct {
		arg1 string
	}
	getServicesForBrokerReturns struct {
		result1 []models.ServiceOffering
		result2 error
	}
	GetServicesForSpaceStub        func(string) ([]models.ServiceOffering, error)
	getServicesForSpaceMutex       sync.RWMutex
	getServicesForSpaceArgsForCall []struct {
		arg1 string
	}
	getServicesForSpaceReturns struct {
		result1 []models.ServiceOffering
		result2 error
	}
	GetServicesForSpaceWithPlansStub        func(string) ([]models.ServiceOffering, error)
	getServicesForSpaceWithPlansMutex       sync.RWMutex
	getServicesForSpaceWithPlansArgsForCall []struct {
		arg1 string
	}
	getServicesForSpaceWithPlansReturns struct {
		result1 []models.ServiceOffering
		result2 error
	}
	GetServiceVisibleToOrgStub        func(string, string) (models.ServiceOffering, error)
	getServiceVisibleToOrgMutex       sync.RWMutex
	getServiceVisibleToOrgArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getServiceVisibleToOrgReturns struct {
		result1 models.ServiceOffering
		result2 error
	}
	GetServicesVisibleToOrgStub        func(string) ([]models.ServiceOffering, error)
	getServicesVisibleToOrgMutex       sync.RWMutex
	getServicesVisibleToOrgArgsForCall []struct {
		arg1 string
	}
	getServicesVisibleToOrgReturns struct {
		result1 []models.ServiceOffering
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceBuilder) GetAllServices() ([]models.ServiceOffering, error) {
	fake.getAllServicesMutex.Lock()
	fake.getAllServicesArgsForCall = append(fake.getAllServicesArgsForCall, struct{}{})
	fake.recordInvocation("GetAllServices", []interface{}{})
	fake.getAllServicesMutex.Unlock()
	if fake.GetAllServicesStub != nil {
		return fake.GetAllServicesStub()
	} else {
		return fake.getAllServicesReturns.result1, fake.getAllServicesReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetAllServicesCallCount() int {
	fake.getAllServicesMutex.RLock()
	defer fake.getAllServicesMutex.RUnlock()
	return len(fake.getAllServicesArgsForCall)
}

func (fake *FakeServiceBuilder) GetAllServicesReturns(result1 []models.ServiceOffering, result2 error) {
	fake.GetAllServicesStub = nil
	fake.getAllServicesReturns = struct {
		result1 []models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) GetAllServicesWithPlans() ([]models.ServiceOffering, error) {
	fake.getAllServicesWithPlansMutex.Lock()
	fake.getAllServicesWithPlansArgsForCall = append(fake.getAllServicesWithPlansArgsForCall, struct{}{})
	fake.recordInvocation("GetAllServicesWithPlans", []interface{}{})
	fake.getAllServicesWithPlansMutex.Unlock()
	if fake.GetAllServicesWithPlansStub != nil {
		return fake.GetAllServicesWithPlansStub()
	} else {
		return fake.getAllServicesWithPlansReturns.result1, fake.getAllServicesWithPlansReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetAllServicesWithPlansCallCount() int {
	fake.getAllServicesWithPlansMutex.RLock()
	defer fake.getAllServicesWithPlansMutex.RUnlock()
	return len(fake.getAllServicesWithPlansArgsForCall)
}

func (fake *FakeServiceBuilder) GetAllServicesWithPlansReturns(result1 []models.ServiceOffering, result2 error) {
	fake.GetAllServicesWithPlansStub = nil
	fake.getAllServicesWithPlansReturns = struct {
		result1 []models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) GetServiceByNameWithPlans(arg1 string) (models.ServiceOffering, error) {
	fake.getServiceByNameWithPlansMutex.Lock()
	fake.getServiceByNameWithPlansArgsForCall = append(fake.getServiceByNameWithPlansArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetServiceByNameWithPlans", []interface{}{arg1})
	fake.getServiceByNameWithPlansMutex.Unlock()
	if fake.GetServiceByNameWithPlansStub != nil {
		return fake.GetServiceByNameWithPlansStub(arg1)
	} else {
		return fake.getServiceByNameWithPlansReturns.result1, fake.getServiceByNameWithPlansReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetServiceByNameWithPlansCallCount() int {
	fake.getServiceByNameWithPlansMutex.RLock()
	defer fake.getServiceByNameWithPlansMutex.RUnlock()
	return len(fake.getServiceByNameWithPlansArgsForCall)
}

func (fake *FakeServiceBuilder) GetServiceByNameWithPlansArgsForCall(i int) string {
	fake.getServiceByNameWithPlansMutex.RLock()
	defer fake.getServiceByNameWithPlansMutex.RUnlock()
	return fake.getServiceByNameWithPlansArgsForCall[i].arg1
}

func (fake *FakeServiceBuilder) GetServiceByNameWithPlansReturns(result1 models.ServiceOffering, result2 error) {
	fake.GetServiceByNameWithPlansStub = nil
	fake.getServiceByNameWithPlansReturns = struct {
		result1 models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) GetServiceByNameWithPlansWithOrgNames(arg1 string) (models.ServiceOffering, error) {
	fake.getServiceByNameWithPlansWithOrgNamesMutex.Lock()
	fake.getServiceByNameWithPlansWithOrgNamesArgsForCall = append(fake.getServiceByNameWithPlansWithOrgNamesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetServiceByNameWithPlansWithOrgNames", []interface{}{arg1})
	fake.getServiceByNameWithPlansWithOrgNamesMutex.Unlock()
	if fake.GetServiceByNameWithPlansWithOrgNamesStub != nil {
		return fake.GetServiceByNameWithPlansWithOrgNamesStub(arg1)
	} else {
		return fake.getServiceByNameWithPlansWithOrgNamesReturns.result1, fake.getServiceByNameWithPlansWithOrgNamesReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetServiceByNameWithPlansWithOrgNamesCallCount() int {
	fake.getServiceByNameWithPlansWithOrgNamesMutex.RLock()
	defer fake.getServiceByNameWithPlansWithOrgNamesMutex.RUnlock()
	return len(fake.getServiceByNameWithPlansWithOrgNamesArgsForCall)
}

func (fake *FakeServiceBuilder) GetServiceByNameWithPlansWithOrgNamesArgsForCall(i int) string {
	fake.getServiceByNameWithPlansWithOrgNamesMutex.RLock()
	defer fake.getServiceByNameWithPlansWithOrgNamesMutex.RUnlock()
	return fake.getServiceByNameWithPlansWithOrgNamesArgsForCall[i].arg1
}

func (fake *FakeServiceBuilder) GetServiceByNameWithPlansWithOrgNamesReturns(result1 models.ServiceOffering, result2 error) {
	fake.GetServiceByNameWithPlansWithOrgNamesStub = nil
	fake.getServiceByNameWithPlansWithOrgNamesReturns = struct {
		result1 models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) GetServiceByNameForSpace(arg1 string, arg2 string) (models.ServiceOffering, error) {
	fake.getServiceByNameForSpaceMutex.Lock()
	fake.getServiceByNameForSpaceArgsForCall = append(fake.getServiceByNameForSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetServiceByNameForSpace", []interface{}{arg1, arg2})
	fake.getServiceByNameForSpaceMutex.Unlock()
	if fake.GetServiceByNameForSpaceStub != nil {
		return fake.GetServiceByNameForSpaceStub(arg1, arg2)
	} else {
		return fake.getServiceByNameForSpaceReturns.result1, fake.getServiceByNameForSpaceReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetServiceByNameForSpaceCallCount() int {
	fake.getServiceByNameForSpaceMutex.RLock()
	defer fake.getServiceByNameForSpaceMutex.RUnlock()
	return len(fake.getServiceByNameForSpaceArgsForCall)
}

func (fake *FakeServiceBuilder) GetServiceByNameForSpaceArgsForCall(i int) (string, string) {
	fake.getServiceByNameForSpaceMutex.RLock()
	defer fake.getServiceByNameForSpaceMutex.RUnlock()
	return fake.getServiceByNameForSpaceArgsForCall[i].arg1, fake.getServiceByNameForSpaceArgsForCall[i].arg2
}

func (fake *FakeServiceBuilder) GetServiceByNameForSpaceReturns(result1 models.ServiceOffering, result2 error) {
	fake.GetServiceByNameForSpaceStub = nil
	fake.getServiceByNameForSpaceReturns = struct {
		result1 models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) GetServiceByNameForSpaceWithPlans(arg1 string, arg2 string) (models.ServiceOffering, error) {
	fake.getServiceByNameForSpaceWithPlansMutex.Lock()
	fake.getServiceByNameForSpaceWithPlansArgsForCall = append(fake.getServiceByNameForSpaceWithPlansArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetServiceByNameForSpaceWithPlans", []interface{}{arg1, arg2})
	fake.getServiceByNameForSpaceWithPlansMutex.Unlock()
	if fake.GetServiceByNameForSpaceWithPlansStub != nil {
		return fake.GetServiceByNameForSpaceWithPlansStub(arg1, arg2)
	} else {
		return fake.getServiceByNameForSpaceWithPlansReturns.result1, fake.getServiceByNameForSpaceWithPlansReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetServiceByNameForSpaceWithPlansCallCount() int {
	fake.getServiceByNameForSpaceWithPlansMutex.RLock()
	defer fake.getServiceByNameForSpaceWithPlansMutex.RUnlock()
	return len(fake.getServiceByNameForSpaceWithPlansArgsForCall)
}

func (fake *FakeServiceBuilder) GetServiceByNameForSpaceWithPlansArgsForCall(i int) (string, string) {
	fake.getServiceByNameForSpaceWithPlansMutex.RLock()
	defer fake.getServiceByNameForSpaceWithPlansMutex.RUnlock()
	return fake.getServiceByNameForSpaceWithPlansArgsForCall[i].arg1, fake.getServiceByNameForSpaceWithPlansArgsForCall[i].arg2
}

func (fake *FakeServiceBuilder) GetServiceByNameForSpaceWithPlansReturns(result1 models.ServiceOffering, result2 error) {
	fake.GetServiceByNameForSpaceWithPlansStub = nil
	fake.getServiceByNameForSpaceWithPlansReturns = struct {
		result1 models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) GetServicesByNameForSpaceWithPlans(arg1 string, arg2 string) (models.ServiceOfferings, error) {
	fake.getServicesByNameForSpaceWithPlansMutex.Lock()
	fake.getServicesByNameForSpaceWithPlansArgsForCall = append(fake.getServicesByNameForSpaceWithPlansArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetServicesByNameForSpaceWithPlans", []interface{}{arg1, arg2})
	fake.getServicesByNameForSpaceWithPlansMutex.Unlock()
	if fake.GetServicesByNameForSpaceWithPlansStub != nil {
		return fake.GetServicesByNameForSpaceWithPlansStub(arg1, arg2)
	} else {
		return fake.getServicesByNameForSpaceWithPlansReturns.result1, fake.getServicesByNameForSpaceWithPlansReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetServicesByNameForSpaceWithPlansCallCount() int {
	fake.getServicesByNameForSpaceWithPlansMutex.RLock()
	defer fake.getServicesByNameForSpaceWithPlansMutex.RUnlock()
	return len(fake.getServicesByNameForSpaceWithPlansArgsForCall)
}

func (fake *FakeServiceBuilder) GetServicesByNameForSpaceWithPlansArgsForCall(i int) (string, string) {
	fake.getServicesByNameForSpaceWithPlansMutex.RLock()
	defer fake.getServicesByNameForSpaceWithPlansMutex.RUnlock()
	return fake.getServicesByNameForSpaceWithPlansArgsForCall[i].arg1, fake.getServicesByNameForSpaceWithPlansArgsForCall[i].arg2
}

func (fake *FakeServiceBuilder) GetServicesByNameForSpaceWithPlansReturns(result1 models.ServiceOfferings, result2 error) {
	fake.GetServicesByNameForSpaceWithPlansStub = nil
	fake.getServicesByNameForSpaceWithPlansReturns = struct {
		result1 models.ServiceOfferings
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) GetServiceByNameForOrg(arg1 string, arg2 string) (models.ServiceOffering, error) {
	fake.getServiceByNameForOrgMutex.Lock()
	fake.getServiceByNameForOrgArgsForCall = append(fake.getServiceByNameForOrgArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetServiceByNameForOrg", []interface{}{arg1, arg2})
	fake.getServiceByNameForOrgMutex.Unlock()
	if fake.GetServiceByNameForOrgStub != nil {
		return fake.GetServiceByNameForOrgStub(arg1, arg2)
	} else {
		return fake.getServiceByNameForOrgReturns.result1, fake.getServiceByNameForOrgReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetServiceByNameForOrgCallCount() int {
	fake.getServiceByNameForOrgMutex.RLock()
	defer fake.getServiceByNameForOrgMutex.RUnlock()
	return len(fake.getServiceByNameForOrgArgsForCall)
}

func (fake *FakeServiceBuilder) GetServiceByNameForOrgArgsForCall(i int) (string, string) {
	fake.getServiceByNameForOrgMutex.RLock()
	defer fake.getServiceByNameForOrgMutex.RUnlock()
	return fake.getServiceByNameForOrgArgsForCall[i].arg1, fake.getServiceByNameForOrgArgsForCall[i].arg2
}

func (fake *FakeServiceBuilder) GetServiceByNameForOrgReturns(result1 models.ServiceOffering, result2 error) {
	fake.GetServiceByNameForOrgStub = nil
	fake.getServiceByNameForOrgReturns = struct {
		result1 models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) GetServicesForManyBrokers(arg1 []string) ([]models.ServiceOffering, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.getServicesForManyBrokersMutex.Lock()
	fake.getServicesForManyBrokersArgsForCall = append(fake.getServicesForManyBrokersArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	fake.recordInvocation("GetServicesForManyBrokers", []interface{}{arg1Copy})
	fake.getServicesForManyBrokersMutex.Unlock()
	if fake.GetServicesForManyBrokersStub != nil {
		return fake.GetServicesForManyBrokersStub(arg1)
	} else {
		return fake.getServicesForManyBrokersReturns.result1, fake.getServicesForManyBrokersReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetServicesForManyBrokersCallCount() int {
	fake.getServicesForManyBrokersMutex.RLock()
	defer fake.getServicesForManyBrokersMutex.RUnlock()
	return len(fake.getServicesForManyBrokersArgsForCall)
}

func (fake *FakeServiceBuilder) GetServicesForManyBrokersArgsForCall(i int) []string {
	fake.getServicesForManyBrokersMutex.RLock()
	defer fake.getServicesForManyBrokersMutex.RUnlock()
	return fake.getServicesForManyBrokersArgsForCall[i].arg1
}

func (fake *FakeServiceBuilder) GetServicesForManyBrokersReturns(result1 []models.ServiceOffering, result2 error) {
	fake.GetServicesForManyBrokersStub = nil
	fake.getServicesForManyBrokersReturns = struct {
		result1 []models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) GetServicesForBroker(arg1 string) ([]models.ServiceOffering, error) {
	fake.getServicesForBrokerMutex.Lock()
	fake.getServicesForBrokerArgsForCall = append(fake.getServicesForBrokerArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetServicesForBroker", []interface{}{arg1})
	fake.getServicesForBrokerMutex.Unlock()
	if fake.GetServicesForBrokerStub != nil {
		return fake.GetServicesForBrokerStub(arg1)
	} else {
		return fake.getServicesForBrokerReturns.result1, fake.getServicesForBrokerReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetServicesForBrokerCallCount() int {
	fake.getServicesForBrokerMutex.RLock()
	defer fake.getServicesForBrokerMutex.RUnlock()
	return len(fake.getServicesForBrokerArgsForCall)
}

func (fake *FakeServiceBuilder) GetServicesForBrokerArgsForCall(i int) string {
	fake.getServicesForBrokerMutex.RLock()
	defer fake.getServicesForBrokerMutex.RUnlock()
	return fake.getServicesForBrokerArgsForCall[i].arg1
}

func (fake *FakeServiceBuilder) GetServicesForBrokerReturns(result1 []models.ServiceOffering, result2 error) {
	fake.GetServicesForBrokerStub = nil
	fake.getServicesForBrokerReturns = struct {
		result1 []models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) GetServicesForSpace(arg1 string) ([]models.ServiceOffering, error) {
	fake.getServicesForSpaceMutex.Lock()
	fake.getServicesForSpaceArgsForCall = append(fake.getServicesForSpaceArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetServicesForSpace", []interface{}{arg1})
	fake.getServicesForSpaceMutex.Unlock()
	if fake.GetServicesForSpaceStub != nil {
		return fake.GetServicesForSpaceStub(arg1)
	} else {
		return fake.getServicesForSpaceReturns.result1, fake.getServicesForSpaceReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetServicesForSpaceCallCount() int {
	fake.getServicesForSpaceMutex.RLock()
	defer fake.getServicesForSpaceMutex.RUnlock()
	return len(fake.getServicesForSpaceArgsForCall)
}

func (fake *FakeServiceBuilder) GetServicesForSpaceArgsForCall(i int) string {
	fake.getServicesForSpaceMutex.RLock()
	defer fake.getServicesForSpaceMutex.RUnlock()
	return fake.getServicesForSpaceArgsForCall[i].arg1
}

func (fake *FakeServiceBuilder) GetServicesForSpaceReturns(result1 []models.ServiceOffering, result2 error) {
	fake.GetServicesForSpaceStub = nil
	fake.getServicesForSpaceReturns = struct {
		result1 []models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) GetServicesForSpaceWithPlans(arg1 string) ([]models.ServiceOffering, error) {
	fake.getServicesForSpaceWithPlansMutex.Lock()
	fake.getServicesForSpaceWithPlansArgsForCall = append(fake.getServicesForSpaceWithPlansArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetServicesForSpaceWithPlans", []interface{}{arg1})
	fake.getServicesForSpaceWithPlansMutex.Unlock()
	if fake.GetServicesForSpaceWithPlansStub != nil {
		return fake.GetServicesForSpaceWithPlansStub(arg1)
	} else {
		return fake.getServicesForSpaceWithPlansReturns.result1, fake.getServicesForSpaceWithPlansReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetServicesForSpaceWithPlansCallCount() int {
	fake.getServicesForSpaceWithPlansMutex.RLock()
	defer fake.getServicesForSpaceWithPlansMutex.RUnlock()
	return len(fake.getServicesForSpaceWithPlansArgsForCall)
}

func (fake *FakeServiceBuilder) GetServicesForSpaceWithPlansArgsForCall(i int) string {
	fake.getServicesForSpaceWithPlansMutex.RLock()
	defer fake.getServicesForSpaceWithPlansMutex.RUnlock()
	return fake.getServicesForSpaceWithPlansArgsForCall[i].arg1
}

func (fake *FakeServiceBuilder) GetServicesForSpaceWithPlansReturns(result1 []models.ServiceOffering, result2 error) {
	fake.GetServicesForSpaceWithPlansStub = nil
	fake.getServicesForSpaceWithPlansReturns = struct {
		result1 []models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) GetServiceVisibleToOrg(arg1 string, arg2 string) (models.ServiceOffering, error) {
	fake.getServiceVisibleToOrgMutex.Lock()
	fake.getServiceVisibleToOrgArgsForCall = append(fake.getServiceVisibleToOrgArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetServiceVisibleToOrg", []interface{}{arg1, arg2})
	fake.getServiceVisibleToOrgMutex.Unlock()
	if fake.GetServiceVisibleToOrgStub != nil {
		return fake.GetServiceVisibleToOrgStub(arg1, arg2)
	} else {
		return fake.getServiceVisibleToOrgReturns.result1, fake.getServiceVisibleToOrgReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetServiceVisibleToOrgCallCount() int {
	fake.getServiceVisibleToOrgMutex.RLock()
	defer fake.getServiceVisibleToOrgMutex.RUnlock()
	return len(fake.getServiceVisibleToOrgArgsForCall)
}

func (fake *FakeServiceBuilder) GetServiceVisibleToOrgArgsForCall(i int) (string, string) {
	fake.getServiceVisibleToOrgMutex.RLock()
	defer fake.getServiceVisibleToOrgMutex.RUnlock()
	return fake.getServiceVisibleToOrgArgsForCall[i].arg1, fake.getServiceVisibleToOrgArgsForCall[i].arg2
}

func (fake *FakeServiceBuilder) GetServiceVisibleToOrgReturns(result1 models.ServiceOffering, result2 error) {
	fake.GetServiceVisibleToOrgStub = nil
	fake.getServiceVisibleToOrgReturns = struct {
		result1 models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) GetServicesVisibleToOrg(arg1 string) ([]models.ServiceOffering, error) {
	fake.getServicesVisibleToOrgMutex.Lock()
	fake.getServicesVisibleToOrgArgsForCall = append(fake.getServicesVisibleToOrgArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetServicesVisibleToOrg", []interface{}{arg1})
	fake.getServicesVisibleToOrgMutex.Unlock()
	if fake.GetServicesVisibleToOrgStub != nil {
		return fake.GetServicesVisibleToOrgStub(arg1)
	} else {
		return fake.getServicesVisibleToOrgReturns.result1, fake.getServicesVisibleToOrgReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetServicesVisibleToOrgCallCount() int {
	fake.getServicesVisibleToOrgMutex.RLock()
	defer fake.getServicesVisibleToOrgMutex.RUnlock()
	return len(fake.getServicesVisibleToOrgArgsForCall)
}

func (fake *FakeServiceBuilder) GetServicesVisibleToOrgArgsForCall(i int) string {
	fake.getServicesVisibleToOrgMutex.RLock()
	defer fake.getServicesVisibleToOrgMutex.RUnlock()
	return fake.getServicesVisibleToOrgArgsForCall[i].arg1
}

func (fake *FakeServiceBuilder) GetServicesVisibleToOrgReturns(result1 []models.ServiceOffering, result2 error) {
	fake.GetServicesVisibleToOrgStub = nil
	fake.getServicesVisibleToOrgReturns = struct {
		result1 []models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAllServicesMutex.RLock()
	defer fake.getAllServicesMutex.RUnlock()
	fake.getAllServicesWithPlansMutex.RLock()
	defer fake.getAllServicesWithPlansMutex.RUnlock()
	fake.getServiceByNameWithPlansMutex.RLock()
	defer fake.getServiceByNameWithPlansMutex.RUnlock()
	fake.getServiceByNameWithPlansWithOrgNamesMutex.RLock()
	defer fake.getServiceByNameWithPlansWithOrgNamesMutex.RUnlock()
	fake.getServiceByNameForSpaceMutex.RLock()
	defer fake.getServiceByNameForSpaceMutex.RUnlock()
	fake.getServiceByNameForSpaceWithPlansMutex.RLock()
	defer fake.getServiceByNameForSpaceWithPlansMutex.RUnlock()
	fake.getServicesByNameForSpaceWithPlansMutex.RLock()
	defer fake.getServicesByNameForSpaceWithPlansMutex.RUnlock()
	fake.getServiceByNameForOrgMutex.RLock()
	defer fake.getServiceByNameForOrgMutex.RUnlock()
	fake.getServicesForManyBrokersMutex.RLock()
	defer fake.getServicesForManyBrokersMutex.RUnlock()
	fake.getServicesForBrokerMutex.RLock()
	defer fake.getServicesForBrokerMutex.RUnlock()
	fake.getServicesForSpaceMutex.RLock()
	defer fake.getServicesForSpaceMutex.RUnlock()
	fake.getServicesForSpaceWithPlansMutex.RLock()
	defer fake.getServicesForSpaceWithPlansMutex.RUnlock()
	fake.getServiceVisibleToOrgMutex.RLock()
	defer fake.getServiceVisibleToOrgMutex.RUnlock()
	fake.getServicesVisibleToOrgMutex.RLock()
	defer fake.getServicesVisibleToOrgMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeServiceBuilder) recordInvocation(key string, args []interface{}) {
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

var _ servicebuilder.ServiceBuilder = new(FakeServiceBuilder)
