// Code generated by mockery v1.1.2. DO NOT EDIT.

package mocks

import (
	distribution "github.com/distribution/distribution/v3"
	mock "github.com/stretchr/testify/mock"
	tag "github.com/argoproj-labs/argocd-image-updater/pkg/tag"
)

// RegistryClient is an autogenerated mock type for the RegistryClient type
type RegistryClient struct {
	mock.Mock
}

func (_m *RegistryClient) TagMetadata(manifest distribution.Manifest) (*tag.TagInfo, error) {
	ret := _m.Called(manifest)

	var r0 *tag.TagInfo
	if rf, ok := ret.Get(0).(func(distribution.Manifest) *tag.TagInfo); ok {
		r0 = rf(manifest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tag.TagInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(distribution.Manifest) error); ok {
		r1 = rf(manifest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Tags provides a mock function with given fields: nameInRepository
func (_m *RegistryClient) Tags() ([]string, error) {
	ret := _m.Called()
	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *RegistryClient) Manifest(tagStr string) (distribution.Manifest, error) {
	ret := _m.Called(tagStr)

	var r0 distribution.Manifest
	if rf, ok := ret.Get(0).(func(string) distribution.Manifest); ok {
		r0 = rf(tagStr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(distribution.Manifest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tagStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *RegistryClient) NewRepository(nameInRepository string) (error){
	ret := _m.Called(nameInRepository)

	var r1 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r1 = rf(nameInRepository)
	} else {
		r1 = ret.Error(0)
	}

	return r1
}
