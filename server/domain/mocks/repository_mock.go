// Code generated by MockGen. DO NOT EDIT.
// Source: server/domain/repositories/repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	repositories "github.com/script_add_products/server/domain/repositories"
)

// MockIRepository is a mock of IRepository interface.
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository.
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance.
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// GetAllData mocks base method.
func (m *MockIRepository) GetAllData(ctx context.Context) ([]repositories.ProducstEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllData", ctx)
	ret0, _ := ret[0].([]repositories.ProducstEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllData indicates an expected call of GetAllData.
func (mr *MockIRepositoryMockRecorder) GetAllData(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllData", reflect.TypeOf((*MockIRepository)(nil).GetAllData), ctx)
}

// UpdateProduct mocks base method.
func (m *MockIRepository) UpdateProduct(ctx context.Context, id int, shopifyId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", ctx, id, shopifyId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockIRepositoryMockRecorder) UpdateProduct(ctx, id, shopifyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockIRepository)(nil).UpdateProduct), ctx, id, shopifyId)
}
