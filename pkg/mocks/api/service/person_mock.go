// Code generated by MockGen. DO NOT EDIT.
// Source: internal/adapters/api/service/person.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	dto "people_service/internal/domain/dto"
	model "people_service/internal/domain/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPersonService is a mock of PersonService interface.
type MockPersonService struct {
	ctrl     *gomock.Controller
	recorder *MockPersonServiceMockRecorder
}

// MockPersonServiceMockRecorder is the mock recorder for MockPersonService.
type MockPersonServiceMockRecorder struct {
	mock *MockPersonService
}

// NewMockPersonService creates a new mock instance.
func NewMockPersonService(ctrl *gomock.Controller) *MockPersonService {
	mock := &MockPersonService{ctrl: ctrl}
	mock.recorder = &MockPersonServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersonService) EXPECT() *MockPersonServiceMockRecorder {
	return m.recorder
}

// AddPerson mocks base method.
func (m *MockPersonService) AddPerson(ctx context.Context, data *dto.AddPersonRawDTO) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPerson", ctx, data)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPerson indicates an expected call of AddPerson.
func (mr *MockPersonServiceMockRecorder) AddPerson(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPerson", reflect.TypeOf((*MockPersonService)(nil).AddPerson), ctx, data)
}

// DeletePerson mocks base method.
func (m *MockPersonService) DeletePerson(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePerson", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePerson indicates an expected call of DeletePerson.
func (mr *MockPersonServiceMockRecorder) DeletePerson(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePerson", reflect.TypeOf((*MockPersonService)(nil).DeletePerson), ctx, id)
}

// GetPerson mocks base method.
func (m *MockPersonService) GetPerson(ctx context.Context, id int64) (*model.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPerson", ctx, id)
	ret0, _ := ret[0].(*model.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPerson indicates an expected call of GetPerson.
func (mr *MockPersonServiceMockRecorder) GetPerson(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPerson", reflect.TypeOf((*MockPersonService)(nil).GetPerson), ctx, id)
}

// GetPersons mocks base method.
func (m *MockPersonService) GetPersons(ctx context.Context, data *dto.PersonsGetDTO) ([]model.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPersons", ctx, data)
	ret0, _ := ret[0].([]model.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPersons indicates an expected call of GetPersons.
func (mr *MockPersonServiceMockRecorder) GetPersons(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPersons", reflect.TypeOf((*MockPersonService)(nil).GetPersons), ctx, data)
}

// UpdatePerson mocks base method.
func (m *MockPersonService) UpdatePerson(ctx context.Context, data *dto.UpdatePersonDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePerson", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePerson indicates an expected call of UpdatePerson.
func (mr *MockPersonServiceMockRecorder) UpdatePerson(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePerson", reflect.TypeOf((*MockPersonService)(nil).UpdatePerson), ctx, data)
}
