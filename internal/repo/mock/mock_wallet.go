package mock_repo

import (
	reflect "reflect"
	model "wallet/internal/model"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockWalletRepoInterface is a mock of WalletRepoInterface interface.
type MockWalletRepoInterface struct {
	ctrl     *gomock.Controller
	recorder *MockWalletRepoInterfaceMockRecorder
}

// MockWalletRepoInterfaceMockRecorder is the mock recorder for MockWalletRepoInterface.
type MockWalletRepoInterfaceMockRecorder struct {
	mock *MockWalletRepoInterface
}

// NewMockWalletRepoInterface creates a new mock instance.
func NewMockWalletRepoInterface(ctrl *gomock.Controller) *MockWalletRepoInterface {
	mock := &MockWalletRepoInterface{ctrl: ctrl}
	mock.recorder = &MockWalletRepoInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWalletRepoInterface) EXPECT() *MockWalletRepoInterfaceMockRecorder {
	return m.recorder
}

// AirdropToken mocks base method.
func (m *MockWalletRepoInterface) AirdropToken(airdroptransaction *model.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AirdropToken", airdroptransaction)
	ret0, _ := ret[0].(error)
	return ret0
}

// AirdropToken indicates an expected call of AirdropToken.
func (mr *MockWalletRepoInterfaceMockRecorder) AirdropToken(airdroptransaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AirdropToken", reflect.TypeOf((*MockWalletRepoInterface)(nil).AirdropToken), airdroptransaction)
}

// CheckWalletExist mocks base method.
func (m *MockWalletRepoInterface) CheckWalletExist(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckWalletExist", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckWalletExist indicates an expected call of CheckWalletExist.
func (mr *MockWalletRepoInterfaceMockRecorder) CheckWalletExist(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckWalletExist", reflect.TypeOf((*MockWalletRepoInterface)(nil).CheckWalletExist), name)
}

// CreateWallet mocks base method.
func (m *MockWalletRepoInterface) CreateWallet(newWallet *model.Wallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWallet", newWallet)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWallet indicates an expected call of CreateWallet.
func (mr *MockWalletRepoInterfaceMockRecorder) CreateWallet(newWallet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWallet", reflect.TypeOf((*MockWalletRepoInterface)(nil).CreateWallet), newWallet)
}

// DeleteWallet mocks base method.
func (m *MockWalletRepoInterface) DeleteWallet(userId, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWallet", userId, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWallet indicates an expected call of DeleteWallet.
func (mr *MockWalletRepoInterfaceMockRecorder) DeleteWallet(userId, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWallet", reflect.TypeOf((*MockWalletRepoInterface)(nil).DeleteWallet), userId, name)
}

// GetAllWallet mocks base method.
func (m *MockWalletRepoInterface) GetAllWallet(order, name, userID string, pageSize, page int) ([]model.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllWallet", order, name, userID, pageSize, page)
	ret0, _ := ret[0].([]model.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllWallet indicates an expected call of GetAllWallet.
func (mr *MockWalletRepoInterfaceMockRecorder) GetAllWallet(order, name, userID, pageSize, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllWallet", reflect.TypeOf((*MockWalletRepoInterface)(nil).GetAllWallet), order, name, userID, pageSize, page)
}

// GetOneWallet mocks base method.
func (m *MockWalletRepoInterface) GetOneWallet(name, userID string) ([]model.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneWallet", name, userID)
	ret0, _ := ret[0].([]model.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneWallet indicates an expected call of GetOneWallet.
func (mr *MockWalletRepoInterfaceMockRecorder) GetOneWallet(name, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneWallet", reflect.TypeOf((*MockWalletRepoInterface)(nil).GetOneWallet), name, userID)
}

// GetUserWalletAddress mocks base method.
func (m *MockWalletRepoInterface) GetUserWalletAddress(userid, name string) uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserWalletAddress", userid, name)
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// GetUserWalletAddress indicates an expected call of GetUserWalletAddress.
func (mr *MockWalletRepoInterfaceMockRecorder) GetUserWalletAddress(userid, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserWalletAddress", reflect.TypeOf((*MockWalletRepoInterface)(nil).GetUserWalletAddress), userid, name)
}

// Update mocks base method.
func (m *MockWalletRepoInterface) Update(userid, name, updateName string) ([]model.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userid, name, updateName)
	ret0, _ := ret[0].([]model.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockWalletRepoInterfaceMockRecorder) Update(userid, name, updateName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWalletRepoInterface)(nil).Update), userid, name, updateName)
}
