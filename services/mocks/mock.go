// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/todo_api/models"
	services "github.com/todo_api/services"
)

// MockAuthServ is a mock of AuthServ interface.
type MockAuthServ struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServMockRecorder
}

// MockAuthServMockRecorder is the mock recorder for MockAuthServ.
type MockAuthServMockRecorder struct {
	mock *MockAuthServ
}

// NewMockAuthServ creates a new mock instance.
func NewMockAuthServ(ctrl *gomock.Controller) *MockAuthServ {
	mock := &MockAuthServ{ctrl: ctrl}
	mock.recorder = &MockAuthServMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthServ) EXPECT() *MockAuthServMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthServ) CreateUser(arg0 context.Context, arg1 *models.User) (*services.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(*services.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthServMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthServ)(nil).CreateUser), arg0, arg1)
}

// LoginUser mocks base method.
func (m *MockAuthServ) LoginUser(arg0 context.Context, arg1 *models.AuthUser) (*services.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", arg0, arg1)
	ret0, _ := ret[0].(*services.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockAuthServMockRecorder) LoginUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockAuthServ)(nil).LoginUser), arg0, arg1)
}

// UpdateToken mocks base method.
func (m *MockAuthServ) UpdateToken(arg0 context.Context, arg1 string) (*services.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateToken", arg0, arg1)
	ret0, _ := ret[0].(*services.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateToken indicates an expected call of UpdateToken.
func (mr *MockAuthServMockRecorder) UpdateToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateToken", reflect.TypeOf((*MockAuthServ)(nil).UpdateToken), arg0, arg1)
}

// MockTaskServ is a mock of TaskServ interface.
type MockTaskServ struct {
	ctrl     *gomock.Controller
	recorder *MockTaskServMockRecorder
}

// MockTaskServMockRecorder is the mock recorder for MockTaskServ.
type MockTaskServMockRecorder struct {
	mock *MockTaskServ
}

// NewMockTaskServ creates a new mock instance.
func NewMockTaskServ(ctrl *gomock.Controller) *MockTaskServ {
	mock := &MockTaskServ{ctrl: ctrl}
	mock.recorder = &MockTaskServMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskServ) EXPECT() *MockTaskServMockRecorder {
	return m.recorder
}

// ChangeStatus mocks base method.
func (m *MockTaskServ) ChangeStatus(arg0 context.Context, arg1 int, arg2 bool) (*models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeStatus indicates an expected call of ChangeStatus.
func (mr *MockTaskServMockRecorder) ChangeStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeStatus", reflect.TypeOf((*MockTaskServ)(nil).ChangeStatus), arg0, arg1, arg2)
}

// CreateTask mocks base method.
func (m *MockTaskServ) CreateTask(arg0 context.Context, arg1 *models.Task) (*models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", arg0, arg1)
	ret0, _ := ret[0].(*models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockTaskServMockRecorder) CreateTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockTaskServ)(nil).CreateTask), arg0, arg1)
}

// DeleteTask mocks base method.
func (m *MockTaskServ) DeleteTask(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockTaskServMockRecorder) DeleteTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockTaskServ)(nil).DeleteTask), arg0, arg1)
}

// EditTask mocks base method.
func (m *MockTaskServ) EditTask(arg0 context.Context, arg1 *models.Task) (*models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditTask", arg0, arg1)
	ret0, _ := ret[0].(*models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditTask indicates an expected call of EditTask.
func (mr *MockTaskServMockRecorder) EditTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditTask", reflect.TypeOf((*MockTaskServ)(nil).EditTask), arg0, arg1)
}

// GetAllTasks mocks base method.
func (m *MockTaskServ) GetAllTasks(arg0 context.Context) ([]models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTasks", arg0)
	ret0, _ := ret[0].([]models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTasks indicates an expected call of GetAllTasks.
func (mr *MockTaskServMockRecorder) GetAllTasks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTasks", reflect.TypeOf((*MockTaskServ)(nil).GetAllTasks), arg0)
}

// MockFolderServ is a mock of FolderServ interface.
type MockFolderServ struct {
	ctrl     *gomock.Controller
	recorder *MockFolderServMockRecorder
}

// MockFolderServMockRecorder is the mock recorder for MockFolderServ.
type MockFolderServMockRecorder struct {
	mock *MockFolderServ
}

// NewMockFolderServ creates a new mock instance.
func NewMockFolderServ(ctrl *gomock.Controller) *MockFolderServ {
	mock := &MockFolderServ{ctrl: ctrl}
	mock.recorder = &MockFolderServMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFolderServ) EXPECT() *MockFolderServMockRecorder {
	return m.recorder
}

// ChangeTitle mocks base method.
func (m *MockFolderServ) ChangeTitle(arg0 context.Context, arg1 *models.Folder) (*models.Folder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeTitle", arg0, arg1)
	ret0, _ := ret[0].(*models.Folder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeTitle indicates an expected call of ChangeTitle.
func (mr *MockFolderServMockRecorder) ChangeTitle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeTitle", reflect.TypeOf((*MockFolderServ)(nil).ChangeTitle), arg0, arg1)
}

// CreateFolder mocks base method.
func (m *MockFolderServ) CreateFolder(arg0 context.Context, arg1 *models.Folder) (*models.Folder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFolder", arg0, arg1)
	ret0, _ := ret[0].(*models.Folder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFolder indicates an expected call of CreateFolder.
func (mr *MockFolderServMockRecorder) CreateFolder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFolder", reflect.TypeOf((*MockFolderServ)(nil).CreateFolder), arg0, arg1)
}

// DeleteFolder mocks base method.
func (m *MockFolderServ) DeleteFolder(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFolder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFolder indicates an expected call of DeleteFolder.
func (mr *MockFolderServMockRecorder) DeleteFolder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFolder", reflect.TypeOf((*MockFolderServ)(nil).DeleteFolder), arg0, arg1)
}

// GetAllFolders mocks base method.
func (m *MockFolderServ) GetAllFolders(arg0 context.Context) ([]models.Folder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFolders", arg0)
	ret0, _ := ret[0].([]models.Folder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFolders indicates an expected call of GetAllFolders.
func (mr *MockFolderServMockRecorder) GetAllFolders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFolders", reflect.TypeOf((*MockFolderServ)(nil).GetAllFolders), arg0)
}