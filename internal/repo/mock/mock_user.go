package mock_repo

import (
	reflect "reflect"
	model "wallet/internal/model"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockUserRepoInterface is a mock of UserRepoInterface interface.
type MockUserRepoInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoInterfaceMockRecorder
}

// MockUserRepoInterfaceMockRecorder is the mock recorder for MockUserRepoInterface.
type MockUserRepoInterfaceMockRecorder struct {
	mock *MockUserRepoInterface
}

// NewMockUserRepoInterface creates a new mock instance.
func NewMockUserRepoInterface(ctrl *gomock.Controller) *MockUserRepoInterface {
	mock := &MockUserRepoInterface{ctrl: ctrl}
	mock.recorder = &MockUserRepoInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepoInterface) EXPECT() *MockUserRepoInterfaceMockRecorder {
	return m.recorder
}

// CheckEmailExist mocks base method.
func (m *MockUserRepoInterface) CheckEmailExist(newEmail string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckEmailExist", newEmail)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckEmailExist indicates an expected call of CheckEmailExist.
func (mr *MockUserRepoInterfaceMockRecorder) CheckEmailExist(newEmail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckEmailExist", reflect.TypeOf((*MockUserRepoInterface)(nil).CheckEmailExist), newEmail)
}

// CreateUser mocks base method.
func (m *MockUserRepoInterface) CreateUser(user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepoInterfaceMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepoInterface)(nil).CreateUser), user)
}

// DeleteUser mocks base method.
func (m *MockUserRepoInterface) DeleteUser(userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserRepoInterfaceMockRecorder) DeleteUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepoInterface)(nil).DeleteUser), userID)
}

// GetAllUsers mocks base method.
func (m *MockUserRepoInterface) GetAllUsers(filterEmail, sortOrder string, page, limit int) ([]model.User, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", filterEmail, sortOrder, page, limit)
	ret0, _ := ret[0].([]model.User)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockUserRepoInterfaceMockRecorder) GetAllUsers(filterEmail, sortOrder, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserRepoInterface)(nil).GetAllUsers), filterEmail, sortOrder, page, limit)
}

// GetRoleID mocks base method.
func (m *MockUserRepoInterface) GetRoleID(namerole string) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleID", namerole)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleID indicates an expected call of GetRoleID.
func (mr *MockUserRepoInterfaceMockRecorder) GetRoleID(namerole interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleID", reflect.TypeOf((*MockUserRepoInterface)(nil).GetRoleID), namerole)
}

// GetUser mocks base method.
func (m *MockUserRepoInterface) GetUser(userID uuid.UUID) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", userID)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserRepoInterfaceMockRecorder) GetUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserRepoInterface)(nil).GetUser), userID)
}

// GetUserByEmail mocks base method.
func (m *MockUserRepoInterface) GetUserByEmail(email string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", email)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockUserRepoInterfaceMockRecorder) GetUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUserRepoInterface)(nil).GetUserByEmail), email)
}

// GetUserByID mocks base method.
func (m *MockUserRepoInterface) GetUserByID(id string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserRepoInterfaceMockRecorder) GetUserByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserRepoInterface)(nil).GetUserByID), id)
}
