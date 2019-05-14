// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import regources "gitlab.com/tokend/regources/generated"

// KeyValues is an autogenerated mock type for the KeyValues type
type KeyValues struct {
	mock.Mock
}

// Value provides a mock function with given fields: key
func (_m *KeyValues) Value(key string) (*regources.KeyValueEntryValue, error) {
	ret := _m.Called(key)

	var r0 *regources.KeyValueEntryValue
	if rf, ok := ret.Get(0).(func(string) *regources.KeyValueEntryValue); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*regources.KeyValueEntryValue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
