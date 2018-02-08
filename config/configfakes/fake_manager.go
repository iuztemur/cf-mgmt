// This file was generated by counterfeiter
package configfakes

import (
	"sync"

	"github.com/pivotalservices/cf-mgmt/config"
)

type FakeManager struct {
	AddOrgToConfigStub        func(orgConfig *config.OrgConfig) error
	addOrgToConfigMutex       sync.RWMutex
	addOrgToConfigArgsForCall []struct {
		orgConfig *config.OrgConfig
	}
	addOrgToConfigReturns struct {
		result1 error
	}
	AddSpaceToConfigStub        func(spaceConfig *config.SpaceConfig) error
	addSpaceToConfigMutex       sync.RWMutex
	addSpaceToConfigArgsForCall []struct {
		spaceConfig *config.SpaceConfig
	}
	addSpaceToConfigReturns struct {
		result1 error
	}
	AddSecurityGroupToSpaceStub        func(orgName, spaceName string, securityGroupDefinition []byte) error
	addSecurityGroupToSpaceMutex       sync.RWMutex
	addSecurityGroupToSpaceArgsForCall []struct {
		orgName                 string
		spaceName               string
		securityGroupDefinition []byte
	}
	addSecurityGroupToSpaceReturns struct {
		result1 error
	}
	AddSecurityGroupStub        func(securityGroupName string, securityGroupDefinition []byte) error
	addSecurityGroupMutex       sync.RWMutex
	addSecurityGroupArgsForCall []struct {
		securityGroupName       string
		securityGroupDefinition []byte
	}
	addSecurityGroupReturns struct {
		result1 error
	}
	CreateConfigIfNotExistsStub        func(uaaOrigin string) error
	createConfigIfNotExistsMutex       sync.RWMutex
	createConfigIfNotExistsArgsForCall []struct {
		uaaOrigin string
	}
	createConfigIfNotExistsReturns struct {
		result1 error
	}
	DeleteConfigIfExistsStub        func() error
	deleteConfigIfExistsMutex       sync.RWMutex
	deleteConfigIfExistsArgsForCall []struct{}
	deleteConfigIfExistsReturns     struct {
		result1 error
	}
	SaveOrgSpacesStub        func(spaces *config.Spaces) error
	saveOrgSpacesMutex       sync.RWMutex
	saveOrgSpacesArgsForCall []struct {
		spaces *config.Spaces
	}
	saveOrgSpacesReturns struct {
		result1 error
	}
	SaveSpaceConfigStub        func(spaceConfig *config.SpaceConfig) error
	saveSpaceConfigMutex       sync.RWMutex
	saveSpaceConfigArgsForCall []struct {
		spaceConfig *config.SpaceConfig
	}
	saveSpaceConfigReturns struct {
		result1 error
	}
	SaveOrgConfigStub        func(orgConfig *config.OrgConfig) error
	saveOrgConfigMutex       sync.RWMutex
	saveOrgConfigArgsForCall []struct {
		orgConfig *config.OrgConfig
	}
	saveOrgConfigReturns struct {
		result1 error
	}
	DeleteOrgConfigStub        func(orgName string) error
	deleteOrgConfigMutex       sync.RWMutex
	deleteOrgConfigArgsForCall []struct {
		orgName string
	}
	deleteOrgConfigReturns struct {
		result1 error
	}
	DeleteSpaceConfigStub        func(orgName, spaceName string) error
	deleteSpaceConfigMutex       sync.RWMutex
	deleteSpaceConfigArgsForCall []struct {
		orgName   string
		spaceName string
	}
	deleteSpaceConfigReturns struct {
		result1 error
	}
	SaveOrgsStub        func(*config.Orgs) error
	saveOrgsMutex       sync.RWMutex
	saveOrgsArgsForCall []struct {
		arg1 *config.Orgs
	}
	saveOrgsReturns struct {
		result1 error
	}
	OrgsStub        func() (*config.Orgs, error)
	orgsMutex       sync.RWMutex
	orgsArgsForCall []struct{}
	orgsReturns     struct {
		result1 *config.Orgs
		result2 error
	}
	OrgSpacesStub        func(orgName string) (*config.Spaces, error)
	orgSpacesMutex       sync.RWMutex
	orgSpacesArgsForCall []struct {
		orgName string
	}
	orgSpacesReturns struct {
		result1 *config.Spaces
		result2 error
	}
	SpacesStub        func() ([]config.Spaces, error)
	spacesMutex       sync.RWMutex
	spacesArgsForCall []struct{}
	spacesReturns     struct {
		result1 []config.Spaces
		result2 error
	}
	GetOrgConfigsStub        func() ([]config.OrgConfig, error)
	getOrgConfigsMutex       sync.RWMutex
	getOrgConfigsArgsForCall []struct{}
	getOrgConfigsReturns     struct {
		result1 []config.OrgConfig
		result2 error
	}
	GetSpaceConfigsStub        func() ([]config.SpaceConfig, error)
	getSpaceConfigsMutex       sync.RWMutex
	getSpaceConfigsArgsForCall []struct{}
	getSpaceConfigsReturns     struct {
		result1 []config.SpaceConfig
		result2 error
	}
	GetASGConfigsStub        func() ([]config.ASGConfig, error)
	getASGConfigsMutex       sync.RWMutex
	getASGConfigsArgsForCall []struct{}
	getASGConfigsReturns     struct {
		result1 []config.ASGConfig
		result2 error
	}
	GetGlobalConfigStub        func() (config.GlobalConfig, error)
	getGlobalConfigMutex       sync.RWMutex
	getGlobalConfigArgsForCall []struct{}
	getGlobalConfigReturns     struct {
		result1 config.GlobalConfig
		result2 error
	}
	GetSpaceDefaultsStub        func() (*config.SpaceConfig, error)
	getSpaceDefaultsMutex       sync.RWMutex
	getSpaceDefaultsArgsForCall []struct{}
	getSpaceDefaultsReturns     struct {
		result1 *config.SpaceConfig
		result2 error
	}
	GetOrgConfigStub        func(orgName string) (*config.OrgConfig, error)
	getOrgConfigMutex       sync.RWMutex
	getOrgConfigArgsForCall []struct {
		orgName string
	}
	getOrgConfigReturns struct {
		result1 *config.OrgConfig
		result2 error
	}
	GetSpaceConfigStub        func(orgName, spaceName string) (*config.SpaceConfig, error)
	getSpaceConfigMutex       sync.RWMutex
	getSpaceConfigArgsForCall []struct {
		orgName   string
		spaceName string
	}
	getSpaceConfigReturns struct {
		result1 *config.SpaceConfig
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManager) AddOrgToConfig(orgConfig *config.OrgConfig) error {
	fake.addOrgToConfigMutex.Lock()
	fake.addOrgToConfigArgsForCall = append(fake.addOrgToConfigArgsForCall, struct {
		orgConfig *config.OrgConfig
	}{orgConfig})
	fake.recordInvocation("AddOrgToConfig", []interface{}{orgConfig})
	fake.addOrgToConfigMutex.Unlock()
	if fake.AddOrgToConfigStub != nil {
		return fake.AddOrgToConfigStub(orgConfig)
	} else {
		return fake.addOrgToConfigReturns.result1
	}
}

func (fake *FakeManager) AddOrgToConfigCallCount() int {
	fake.addOrgToConfigMutex.RLock()
	defer fake.addOrgToConfigMutex.RUnlock()
	return len(fake.addOrgToConfigArgsForCall)
}

func (fake *FakeManager) AddOrgToConfigArgsForCall(i int) *config.OrgConfig {
	fake.addOrgToConfigMutex.RLock()
	defer fake.addOrgToConfigMutex.RUnlock()
	return fake.addOrgToConfigArgsForCall[i].orgConfig
}

func (fake *FakeManager) AddOrgToConfigReturns(result1 error) {
	fake.AddOrgToConfigStub = nil
	fake.addOrgToConfigReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) AddSpaceToConfig(spaceConfig *config.SpaceConfig) error {
	fake.addSpaceToConfigMutex.Lock()
	fake.addSpaceToConfigArgsForCall = append(fake.addSpaceToConfigArgsForCall, struct {
		spaceConfig *config.SpaceConfig
	}{spaceConfig})
	fake.recordInvocation("AddSpaceToConfig", []interface{}{spaceConfig})
	fake.addSpaceToConfigMutex.Unlock()
	if fake.AddSpaceToConfigStub != nil {
		return fake.AddSpaceToConfigStub(spaceConfig)
	} else {
		return fake.addSpaceToConfigReturns.result1
	}
}

func (fake *FakeManager) AddSpaceToConfigCallCount() int {
	fake.addSpaceToConfigMutex.RLock()
	defer fake.addSpaceToConfigMutex.RUnlock()
	return len(fake.addSpaceToConfigArgsForCall)
}

func (fake *FakeManager) AddSpaceToConfigArgsForCall(i int) *config.SpaceConfig {
	fake.addSpaceToConfigMutex.RLock()
	defer fake.addSpaceToConfigMutex.RUnlock()
	return fake.addSpaceToConfigArgsForCall[i].spaceConfig
}

func (fake *FakeManager) AddSpaceToConfigReturns(result1 error) {
	fake.AddSpaceToConfigStub = nil
	fake.addSpaceToConfigReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) AddSecurityGroupToSpace(orgName string, spaceName string, securityGroupDefinition []byte) error {
	var securityGroupDefinitionCopy []byte
	if securityGroupDefinition != nil {
		securityGroupDefinitionCopy = make([]byte, len(securityGroupDefinition))
		copy(securityGroupDefinitionCopy, securityGroupDefinition)
	}
	fake.addSecurityGroupToSpaceMutex.Lock()
	fake.addSecurityGroupToSpaceArgsForCall = append(fake.addSecurityGroupToSpaceArgsForCall, struct {
		orgName                 string
		spaceName               string
		securityGroupDefinition []byte
	}{orgName, spaceName, securityGroupDefinitionCopy})
	fake.recordInvocation("AddSecurityGroupToSpace", []interface{}{orgName, spaceName, securityGroupDefinitionCopy})
	fake.addSecurityGroupToSpaceMutex.Unlock()
	if fake.AddSecurityGroupToSpaceStub != nil {
		return fake.AddSecurityGroupToSpaceStub(orgName, spaceName, securityGroupDefinition)
	} else {
		return fake.addSecurityGroupToSpaceReturns.result1
	}
}

func (fake *FakeManager) AddSecurityGroupToSpaceCallCount() int {
	fake.addSecurityGroupToSpaceMutex.RLock()
	defer fake.addSecurityGroupToSpaceMutex.RUnlock()
	return len(fake.addSecurityGroupToSpaceArgsForCall)
}

func (fake *FakeManager) AddSecurityGroupToSpaceArgsForCall(i int) (string, string, []byte) {
	fake.addSecurityGroupToSpaceMutex.RLock()
	defer fake.addSecurityGroupToSpaceMutex.RUnlock()
	return fake.addSecurityGroupToSpaceArgsForCall[i].orgName, fake.addSecurityGroupToSpaceArgsForCall[i].spaceName, fake.addSecurityGroupToSpaceArgsForCall[i].securityGroupDefinition
}

func (fake *FakeManager) AddSecurityGroupToSpaceReturns(result1 error) {
	fake.AddSecurityGroupToSpaceStub = nil
	fake.addSecurityGroupToSpaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) AddSecurityGroup(securityGroupName string, securityGroupDefinition []byte) error {
	var securityGroupDefinitionCopy []byte
	if securityGroupDefinition != nil {
		securityGroupDefinitionCopy = make([]byte, len(securityGroupDefinition))
		copy(securityGroupDefinitionCopy, securityGroupDefinition)
	}
	fake.addSecurityGroupMutex.Lock()
	fake.addSecurityGroupArgsForCall = append(fake.addSecurityGroupArgsForCall, struct {
		securityGroupName       string
		securityGroupDefinition []byte
	}{securityGroupName, securityGroupDefinitionCopy})
	fake.recordInvocation("AddSecurityGroup", []interface{}{securityGroupName, securityGroupDefinitionCopy})
	fake.addSecurityGroupMutex.Unlock()
	if fake.AddSecurityGroupStub != nil {
		return fake.AddSecurityGroupStub(securityGroupName, securityGroupDefinition)
	} else {
		return fake.addSecurityGroupReturns.result1
	}
}

func (fake *FakeManager) AddSecurityGroupCallCount() int {
	fake.addSecurityGroupMutex.RLock()
	defer fake.addSecurityGroupMutex.RUnlock()
	return len(fake.addSecurityGroupArgsForCall)
}

func (fake *FakeManager) AddSecurityGroupArgsForCall(i int) (string, []byte) {
	fake.addSecurityGroupMutex.RLock()
	defer fake.addSecurityGroupMutex.RUnlock()
	return fake.addSecurityGroupArgsForCall[i].securityGroupName, fake.addSecurityGroupArgsForCall[i].securityGroupDefinition
}

func (fake *FakeManager) AddSecurityGroupReturns(result1 error) {
	fake.AddSecurityGroupStub = nil
	fake.addSecurityGroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) CreateConfigIfNotExists(uaaOrigin string) error {
	fake.createConfigIfNotExistsMutex.Lock()
	fake.createConfigIfNotExistsArgsForCall = append(fake.createConfigIfNotExistsArgsForCall, struct {
		uaaOrigin string
	}{uaaOrigin})
	fake.recordInvocation("CreateConfigIfNotExists", []interface{}{uaaOrigin})
	fake.createConfigIfNotExistsMutex.Unlock()
	if fake.CreateConfigIfNotExistsStub != nil {
		return fake.CreateConfigIfNotExistsStub(uaaOrigin)
	} else {
		return fake.createConfigIfNotExistsReturns.result1
	}
}

func (fake *FakeManager) CreateConfigIfNotExistsCallCount() int {
	fake.createConfigIfNotExistsMutex.RLock()
	defer fake.createConfigIfNotExistsMutex.RUnlock()
	return len(fake.createConfigIfNotExistsArgsForCall)
}

func (fake *FakeManager) CreateConfigIfNotExistsArgsForCall(i int) string {
	fake.createConfigIfNotExistsMutex.RLock()
	defer fake.createConfigIfNotExistsMutex.RUnlock()
	return fake.createConfigIfNotExistsArgsForCall[i].uaaOrigin
}

func (fake *FakeManager) CreateConfigIfNotExistsReturns(result1 error) {
	fake.CreateConfigIfNotExistsStub = nil
	fake.createConfigIfNotExistsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) DeleteConfigIfExists() error {
	fake.deleteConfigIfExistsMutex.Lock()
	fake.deleteConfigIfExistsArgsForCall = append(fake.deleteConfigIfExistsArgsForCall, struct{}{})
	fake.recordInvocation("DeleteConfigIfExists", []interface{}{})
	fake.deleteConfigIfExistsMutex.Unlock()
	if fake.DeleteConfigIfExistsStub != nil {
		return fake.DeleteConfigIfExistsStub()
	} else {
		return fake.deleteConfigIfExistsReturns.result1
	}
}

func (fake *FakeManager) DeleteConfigIfExistsCallCount() int {
	fake.deleteConfigIfExistsMutex.RLock()
	defer fake.deleteConfigIfExistsMutex.RUnlock()
	return len(fake.deleteConfigIfExistsArgsForCall)
}

func (fake *FakeManager) DeleteConfigIfExistsReturns(result1 error) {
	fake.DeleteConfigIfExistsStub = nil
	fake.deleteConfigIfExistsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) SaveOrgSpaces(spaces *config.Spaces) error {
	fake.saveOrgSpacesMutex.Lock()
	fake.saveOrgSpacesArgsForCall = append(fake.saveOrgSpacesArgsForCall, struct {
		spaces *config.Spaces
	}{spaces})
	fake.recordInvocation("SaveOrgSpaces", []interface{}{spaces})
	fake.saveOrgSpacesMutex.Unlock()
	if fake.SaveOrgSpacesStub != nil {
		return fake.SaveOrgSpacesStub(spaces)
	} else {
		return fake.saveOrgSpacesReturns.result1
	}
}

func (fake *FakeManager) SaveOrgSpacesCallCount() int {
	fake.saveOrgSpacesMutex.RLock()
	defer fake.saveOrgSpacesMutex.RUnlock()
	return len(fake.saveOrgSpacesArgsForCall)
}

func (fake *FakeManager) SaveOrgSpacesArgsForCall(i int) *config.Spaces {
	fake.saveOrgSpacesMutex.RLock()
	defer fake.saveOrgSpacesMutex.RUnlock()
	return fake.saveOrgSpacesArgsForCall[i].spaces
}

func (fake *FakeManager) SaveOrgSpacesReturns(result1 error) {
	fake.SaveOrgSpacesStub = nil
	fake.saveOrgSpacesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) SaveSpaceConfig(spaceConfig *config.SpaceConfig) error {
	fake.saveSpaceConfigMutex.Lock()
	fake.saveSpaceConfigArgsForCall = append(fake.saveSpaceConfigArgsForCall, struct {
		spaceConfig *config.SpaceConfig
	}{spaceConfig})
	fake.recordInvocation("SaveSpaceConfig", []interface{}{spaceConfig})
	fake.saveSpaceConfigMutex.Unlock()
	if fake.SaveSpaceConfigStub != nil {
		return fake.SaveSpaceConfigStub(spaceConfig)
	} else {
		return fake.saveSpaceConfigReturns.result1
	}
}

func (fake *FakeManager) SaveSpaceConfigCallCount() int {
	fake.saveSpaceConfigMutex.RLock()
	defer fake.saveSpaceConfigMutex.RUnlock()
	return len(fake.saveSpaceConfigArgsForCall)
}

func (fake *FakeManager) SaveSpaceConfigArgsForCall(i int) *config.SpaceConfig {
	fake.saveSpaceConfigMutex.RLock()
	defer fake.saveSpaceConfigMutex.RUnlock()
	return fake.saveSpaceConfigArgsForCall[i].spaceConfig
}

func (fake *FakeManager) SaveSpaceConfigReturns(result1 error) {
	fake.SaveSpaceConfigStub = nil
	fake.saveSpaceConfigReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) SaveOrgConfig(orgConfig *config.OrgConfig) error {
	fake.saveOrgConfigMutex.Lock()
	fake.saveOrgConfigArgsForCall = append(fake.saveOrgConfigArgsForCall, struct {
		orgConfig *config.OrgConfig
	}{orgConfig})
	fake.recordInvocation("SaveOrgConfig", []interface{}{orgConfig})
	fake.saveOrgConfigMutex.Unlock()
	if fake.SaveOrgConfigStub != nil {
		return fake.SaveOrgConfigStub(orgConfig)
	} else {
		return fake.saveOrgConfigReturns.result1
	}
}

func (fake *FakeManager) SaveOrgConfigCallCount() int {
	fake.saveOrgConfigMutex.RLock()
	defer fake.saveOrgConfigMutex.RUnlock()
	return len(fake.saveOrgConfigArgsForCall)
}

func (fake *FakeManager) SaveOrgConfigArgsForCall(i int) *config.OrgConfig {
	fake.saveOrgConfigMutex.RLock()
	defer fake.saveOrgConfigMutex.RUnlock()
	return fake.saveOrgConfigArgsForCall[i].orgConfig
}

func (fake *FakeManager) SaveOrgConfigReturns(result1 error) {
	fake.SaveOrgConfigStub = nil
	fake.saveOrgConfigReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) DeleteOrgConfig(orgName string) error {
	fake.deleteOrgConfigMutex.Lock()
	fake.deleteOrgConfigArgsForCall = append(fake.deleteOrgConfigArgsForCall, struct {
		orgName string
	}{orgName})
	fake.recordInvocation("DeleteOrgConfig", []interface{}{orgName})
	fake.deleteOrgConfigMutex.Unlock()
	if fake.DeleteOrgConfigStub != nil {
		return fake.DeleteOrgConfigStub(orgName)
	} else {
		return fake.deleteOrgConfigReturns.result1
	}
}

func (fake *FakeManager) DeleteOrgConfigCallCount() int {
	fake.deleteOrgConfigMutex.RLock()
	defer fake.deleteOrgConfigMutex.RUnlock()
	return len(fake.deleteOrgConfigArgsForCall)
}

func (fake *FakeManager) DeleteOrgConfigArgsForCall(i int) string {
	fake.deleteOrgConfigMutex.RLock()
	defer fake.deleteOrgConfigMutex.RUnlock()
	return fake.deleteOrgConfigArgsForCall[i].orgName
}

func (fake *FakeManager) DeleteOrgConfigReturns(result1 error) {
	fake.DeleteOrgConfigStub = nil
	fake.deleteOrgConfigReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) DeleteSpaceConfig(orgName string, spaceName string) error {
	fake.deleteSpaceConfigMutex.Lock()
	fake.deleteSpaceConfigArgsForCall = append(fake.deleteSpaceConfigArgsForCall, struct {
		orgName   string
		spaceName string
	}{orgName, spaceName})
	fake.recordInvocation("DeleteSpaceConfig", []interface{}{orgName, spaceName})
	fake.deleteSpaceConfigMutex.Unlock()
	if fake.DeleteSpaceConfigStub != nil {
		return fake.DeleteSpaceConfigStub(orgName, spaceName)
	} else {
		return fake.deleteSpaceConfigReturns.result1
	}
}

func (fake *FakeManager) DeleteSpaceConfigCallCount() int {
	fake.deleteSpaceConfigMutex.RLock()
	defer fake.deleteSpaceConfigMutex.RUnlock()
	return len(fake.deleteSpaceConfigArgsForCall)
}

func (fake *FakeManager) DeleteSpaceConfigArgsForCall(i int) (string, string) {
	fake.deleteSpaceConfigMutex.RLock()
	defer fake.deleteSpaceConfigMutex.RUnlock()
	return fake.deleteSpaceConfigArgsForCall[i].orgName, fake.deleteSpaceConfigArgsForCall[i].spaceName
}

func (fake *FakeManager) DeleteSpaceConfigReturns(result1 error) {
	fake.DeleteSpaceConfigStub = nil
	fake.deleteSpaceConfigReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) SaveOrgs(arg1 *config.Orgs) error {
	fake.saveOrgsMutex.Lock()
	fake.saveOrgsArgsForCall = append(fake.saveOrgsArgsForCall, struct {
		arg1 *config.Orgs
	}{arg1})
	fake.recordInvocation("SaveOrgs", []interface{}{arg1})
	fake.saveOrgsMutex.Unlock()
	if fake.SaveOrgsStub != nil {
		return fake.SaveOrgsStub(arg1)
	} else {
		return fake.saveOrgsReturns.result1
	}
}

func (fake *FakeManager) SaveOrgsCallCount() int {
	fake.saveOrgsMutex.RLock()
	defer fake.saveOrgsMutex.RUnlock()
	return len(fake.saveOrgsArgsForCall)
}

func (fake *FakeManager) SaveOrgsArgsForCall(i int) *config.Orgs {
	fake.saveOrgsMutex.RLock()
	defer fake.saveOrgsMutex.RUnlock()
	return fake.saveOrgsArgsForCall[i].arg1
}

func (fake *FakeManager) SaveOrgsReturns(result1 error) {
	fake.SaveOrgsStub = nil
	fake.saveOrgsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) Orgs() (*config.Orgs, error) {
	fake.orgsMutex.Lock()
	fake.orgsArgsForCall = append(fake.orgsArgsForCall, struct{}{})
	fake.recordInvocation("Orgs", []interface{}{})
	fake.orgsMutex.Unlock()
	if fake.OrgsStub != nil {
		return fake.OrgsStub()
	} else {
		return fake.orgsReturns.result1, fake.orgsReturns.result2
	}
}

func (fake *FakeManager) OrgsCallCount() int {
	fake.orgsMutex.RLock()
	defer fake.orgsMutex.RUnlock()
	return len(fake.orgsArgsForCall)
}

func (fake *FakeManager) OrgsReturns(result1 *config.Orgs, result2 error) {
	fake.OrgsStub = nil
	fake.orgsReturns = struct {
		result1 *config.Orgs
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) OrgSpaces(orgName string) (*config.Spaces, error) {
	fake.orgSpacesMutex.Lock()
	fake.orgSpacesArgsForCall = append(fake.orgSpacesArgsForCall, struct {
		orgName string
	}{orgName})
	fake.recordInvocation("OrgSpaces", []interface{}{orgName})
	fake.orgSpacesMutex.Unlock()
	if fake.OrgSpacesStub != nil {
		return fake.OrgSpacesStub(orgName)
	} else {
		return fake.orgSpacesReturns.result1, fake.orgSpacesReturns.result2
	}
}

func (fake *FakeManager) OrgSpacesCallCount() int {
	fake.orgSpacesMutex.RLock()
	defer fake.orgSpacesMutex.RUnlock()
	return len(fake.orgSpacesArgsForCall)
}

func (fake *FakeManager) OrgSpacesArgsForCall(i int) string {
	fake.orgSpacesMutex.RLock()
	defer fake.orgSpacesMutex.RUnlock()
	return fake.orgSpacesArgsForCall[i].orgName
}

func (fake *FakeManager) OrgSpacesReturns(result1 *config.Spaces, result2 error) {
	fake.OrgSpacesStub = nil
	fake.orgSpacesReturns = struct {
		result1 *config.Spaces
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) Spaces() ([]config.Spaces, error) {
	fake.spacesMutex.Lock()
	fake.spacesArgsForCall = append(fake.spacesArgsForCall, struct{}{})
	fake.recordInvocation("Spaces", []interface{}{})
	fake.spacesMutex.Unlock()
	if fake.SpacesStub != nil {
		return fake.SpacesStub()
	} else {
		return fake.spacesReturns.result1, fake.spacesReturns.result2
	}
}

func (fake *FakeManager) SpacesCallCount() int {
	fake.spacesMutex.RLock()
	defer fake.spacesMutex.RUnlock()
	return len(fake.spacesArgsForCall)
}

func (fake *FakeManager) SpacesReturns(result1 []config.Spaces, result2 error) {
	fake.SpacesStub = nil
	fake.spacesReturns = struct {
		result1 []config.Spaces
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) GetOrgConfigs() ([]config.OrgConfig, error) {
	fake.getOrgConfigsMutex.Lock()
	fake.getOrgConfigsArgsForCall = append(fake.getOrgConfigsArgsForCall, struct{}{})
	fake.recordInvocation("GetOrgConfigs", []interface{}{})
	fake.getOrgConfigsMutex.Unlock()
	if fake.GetOrgConfigsStub != nil {
		return fake.GetOrgConfigsStub()
	} else {
		return fake.getOrgConfigsReturns.result1, fake.getOrgConfigsReturns.result2
	}
}

func (fake *FakeManager) GetOrgConfigsCallCount() int {
	fake.getOrgConfigsMutex.RLock()
	defer fake.getOrgConfigsMutex.RUnlock()
	return len(fake.getOrgConfigsArgsForCall)
}

func (fake *FakeManager) GetOrgConfigsReturns(result1 []config.OrgConfig, result2 error) {
	fake.GetOrgConfigsStub = nil
	fake.getOrgConfigsReturns = struct {
		result1 []config.OrgConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) GetSpaceConfigs() ([]config.SpaceConfig, error) {
	fake.getSpaceConfigsMutex.Lock()
	fake.getSpaceConfigsArgsForCall = append(fake.getSpaceConfigsArgsForCall, struct{}{})
	fake.recordInvocation("GetSpaceConfigs", []interface{}{})
	fake.getSpaceConfigsMutex.Unlock()
	if fake.GetSpaceConfigsStub != nil {
		return fake.GetSpaceConfigsStub()
	} else {
		return fake.getSpaceConfigsReturns.result1, fake.getSpaceConfigsReturns.result2
	}
}

func (fake *FakeManager) GetSpaceConfigsCallCount() int {
	fake.getSpaceConfigsMutex.RLock()
	defer fake.getSpaceConfigsMutex.RUnlock()
	return len(fake.getSpaceConfigsArgsForCall)
}

func (fake *FakeManager) GetSpaceConfigsReturns(result1 []config.SpaceConfig, result2 error) {
	fake.GetSpaceConfigsStub = nil
	fake.getSpaceConfigsReturns = struct {
		result1 []config.SpaceConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) GetASGConfigs() ([]config.ASGConfig, error) {
	fake.getASGConfigsMutex.Lock()
	fake.getASGConfigsArgsForCall = append(fake.getASGConfigsArgsForCall, struct{}{})
	fake.recordInvocation("GetASGConfigs", []interface{}{})
	fake.getASGConfigsMutex.Unlock()
	if fake.GetASGConfigsStub != nil {
		return fake.GetASGConfigsStub()
	} else {
		return fake.getASGConfigsReturns.result1, fake.getASGConfigsReturns.result2
	}
}

func (fake *FakeManager) GetASGConfigsCallCount() int {
	fake.getASGConfigsMutex.RLock()
	defer fake.getASGConfigsMutex.RUnlock()
	return len(fake.getASGConfigsArgsForCall)
}

func (fake *FakeManager) GetASGConfigsReturns(result1 []config.ASGConfig, result2 error) {
	fake.GetASGConfigsStub = nil
	fake.getASGConfigsReturns = struct {
		result1 []config.ASGConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) GetGlobalConfig() (config.GlobalConfig, error) {
	fake.getGlobalConfigMutex.Lock()
	fake.getGlobalConfigArgsForCall = append(fake.getGlobalConfigArgsForCall, struct{}{})
	fake.recordInvocation("GetGlobalConfig", []interface{}{})
	fake.getGlobalConfigMutex.Unlock()
	if fake.GetGlobalConfigStub != nil {
		return fake.GetGlobalConfigStub()
	} else {
		return fake.getGlobalConfigReturns.result1, fake.getGlobalConfigReturns.result2
	}
}

func (fake *FakeManager) GetGlobalConfigCallCount() int {
	fake.getGlobalConfigMutex.RLock()
	defer fake.getGlobalConfigMutex.RUnlock()
	return len(fake.getGlobalConfigArgsForCall)
}

func (fake *FakeManager) GetGlobalConfigReturns(result1 config.GlobalConfig, result2 error) {
	fake.GetGlobalConfigStub = nil
	fake.getGlobalConfigReturns = struct {
		result1 config.GlobalConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) GetSpaceDefaults() (*config.SpaceConfig, error) {
	fake.getSpaceDefaultsMutex.Lock()
	fake.getSpaceDefaultsArgsForCall = append(fake.getSpaceDefaultsArgsForCall, struct{}{})
	fake.recordInvocation("GetSpaceDefaults", []interface{}{})
	fake.getSpaceDefaultsMutex.Unlock()
	if fake.GetSpaceDefaultsStub != nil {
		return fake.GetSpaceDefaultsStub()
	} else {
		return fake.getSpaceDefaultsReturns.result1, fake.getSpaceDefaultsReturns.result2
	}
}

func (fake *FakeManager) GetSpaceDefaultsCallCount() int {
	fake.getSpaceDefaultsMutex.RLock()
	defer fake.getSpaceDefaultsMutex.RUnlock()
	return len(fake.getSpaceDefaultsArgsForCall)
}

func (fake *FakeManager) GetSpaceDefaultsReturns(result1 *config.SpaceConfig, result2 error) {
	fake.GetSpaceDefaultsStub = nil
	fake.getSpaceDefaultsReturns = struct {
		result1 *config.SpaceConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) GetOrgConfig(orgName string) (*config.OrgConfig, error) {
	fake.getOrgConfigMutex.Lock()
	fake.getOrgConfigArgsForCall = append(fake.getOrgConfigArgsForCall, struct {
		orgName string
	}{orgName})
	fake.recordInvocation("GetOrgConfig", []interface{}{orgName})
	fake.getOrgConfigMutex.Unlock()
	if fake.GetOrgConfigStub != nil {
		return fake.GetOrgConfigStub(orgName)
	} else {
		return fake.getOrgConfigReturns.result1, fake.getOrgConfigReturns.result2
	}
}

func (fake *FakeManager) GetOrgConfigCallCount() int {
	fake.getOrgConfigMutex.RLock()
	defer fake.getOrgConfigMutex.RUnlock()
	return len(fake.getOrgConfigArgsForCall)
}

func (fake *FakeManager) GetOrgConfigArgsForCall(i int) string {
	fake.getOrgConfigMutex.RLock()
	defer fake.getOrgConfigMutex.RUnlock()
	return fake.getOrgConfigArgsForCall[i].orgName
}

func (fake *FakeManager) GetOrgConfigReturns(result1 *config.OrgConfig, result2 error) {
	fake.GetOrgConfigStub = nil
	fake.getOrgConfigReturns = struct {
		result1 *config.OrgConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) GetSpaceConfig(orgName string, spaceName string) (*config.SpaceConfig, error) {
	fake.getSpaceConfigMutex.Lock()
	fake.getSpaceConfigArgsForCall = append(fake.getSpaceConfigArgsForCall, struct {
		orgName   string
		spaceName string
	}{orgName, spaceName})
	fake.recordInvocation("GetSpaceConfig", []interface{}{orgName, spaceName})
	fake.getSpaceConfigMutex.Unlock()
	if fake.GetSpaceConfigStub != nil {
		return fake.GetSpaceConfigStub(orgName, spaceName)
	} else {
		return fake.getSpaceConfigReturns.result1, fake.getSpaceConfigReturns.result2
	}
}

func (fake *FakeManager) GetSpaceConfigCallCount() int {
	fake.getSpaceConfigMutex.RLock()
	defer fake.getSpaceConfigMutex.RUnlock()
	return len(fake.getSpaceConfigArgsForCall)
}

func (fake *FakeManager) GetSpaceConfigArgsForCall(i int) (string, string) {
	fake.getSpaceConfigMutex.RLock()
	defer fake.getSpaceConfigMutex.RUnlock()
	return fake.getSpaceConfigArgsForCall[i].orgName, fake.getSpaceConfigArgsForCall[i].spaceName
}

func (fake *FakeManager) GetSpaceConfigReturns(result1 *config.SpaceConfig, result2 error) {
	fake.GetSpaceConfigStub = nil
	fake.getSpaceConfigReturns = struct {
		result1 *config.SpaceConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addOrgToConfigMutex.RLock()
	defer fake.addOrgToConfigMutex.RUnlock()
	fake.addSpaceToConfigMutex.RLock()
	defer fake.addSpaceToConfigMutex.RUnlock()
	fake.addSecurityGroupToSpaceMutex.RLock()
	defer fake.addSecurityGroupToSpaceMutex.RUnlock()
	fake.addSecurityGroupMutex.RLock()
	defer fake.addSecurityGroupMutex.RUnlock()
	fake.createConfigIfNotExistsMutex.RLock()
	defer fake.createConfigIfNotExistsMutex.RUnlock()
	fake.deleteConfigIfExistsMutex.RLock()
	defer fake.deleteConfigIfExistsMutex.RUnlock()
	fake.saveOrgSpacesMutex.RLock()
	defer fake.saveOrgSpacesMutex.RUnlock()
	fake.saveSpaceConfigMutex.RLock()
	defer fake.saveSpaceConfigMutex.RUnlock()
	fake.saveOrgConfigMutex.RLock()
	defer fake.saveOrgConfigMutex.RUnlock()
	fake.deleteOrgConfigMutex.RLock()
	defer fake.deleteOrgConfigMutex.RUnlock()
	fake.deleteSpaceConfigMutex.RLock()
	defer fake.deleteSpaceConfigMutex.RUnlock()
	fake.saveOrgsMutex.RLock()
	defer fake.saveOrgsMutex.RUnlock()
	fake.orgsMutex.RLock()
	defer fake.orgsMutex.RUnlock()
	fake.orgSpacesMutex.RLock()
	defer fake.orgSpacesMutex.RUnlock()
	fake.spacesMutex.RLock()
	defer fake.spacesMutex.RUnlock()
	fake.getOrgConfigsMutex.RLock()
	defer fake.getOrgConfigsMutex.RUnlock()
	fake.getSpaceConfigsMutex.RLock()
	defer fake.getSpaceConfigsMutex.RUnlock()
	fake.getASGConfigsMutex.RLock()
	defer fake.getASGConfigsMutex.RUnlock()
	fake.getGlobalConfigMutex.RLock()
	defer fake.getGlobalConfigMutex.RUnlock()
	fake.getSpaceDefaultsMutex.RLock()
	defer fake.getSpaceDefaultsMutex.RUnlock()
	fake.getOrgConfigMutex.RLock()
	defer fake.getOrgConfigMutex.RUnlock()
	fake.getSpaceConfigMutex.RLock()
	defer fake.getSpaceConfigMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeManager) recordInvocation(key string, args []interface{}) {
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

var _ config.Manager = new(FakeManager)
