// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	domain "clean-arch/domain"

	mock "github.com/stretchr/testify/mock"
)

// NewTrendingStickerUsecase is an autogenerated mock type for the NewTrendingStickerUsecase type
type NewTrendingStickerUsecase struct {
	mock.Mock
}

// GetByID provides a mock function with given fields:
func (_m *NewTrendingStickerUsecase) GetByID() ([]domain.Trendingsticker, error) {
	ret := _m.Called()

	var r0 []domain.Trendingsticker
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.Trendingsticker, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.Trendingsticker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Trendingsticker)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewNewTrendingStickerUsecase creates a new instance of NewTrendingStickerUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNewTrendingStickerUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *NewTrendingStickerUsecase {
	mock := &NewTrendingStickerUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
