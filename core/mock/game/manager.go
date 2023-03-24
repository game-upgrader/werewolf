// Code generated by MockGen. DO NOT EDIT.
// Source: game/contract/manager.go

// Package mock_game is a generated GoMock package.
package mock_game

import (
	reflect "reflect"
	contract "uwwolf/game/contract"
	types "uwwolf/game/types"

	gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Moderator mocks base method.
func (m *MockManager) Moderator(gameID types.GameID) contract.Moderator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Moderator", gameID)
	ret0, _ := ret[0].(contract.Moderator)
	return ret0
}

// Moderator indicates an expected call of Moderator.
func (mr *MockManagerMockRecorder) Moderator(gameID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Moderator", reflect.TypeOf((*MockManager)(nil).Moderator), gameID)
}

// RegisterGame mocks base method.
func (m *MockManager) RegisterGame(registration *types.GameRegistration) (contract.Moderator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterGame", registration)
	ret0, _ := ret[0].(contract.Moderator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterGame indicates an expected call of RegisterGame.
func (mr *MockManagerMockRecorder) RegisterGame(registration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterGame", reflect.TypeOf((*MockManager)(nil).RegisterGame), registration)
}
