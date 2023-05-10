// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/game/logic/contract/moderator.go

// Package mock_game_logic is a generated GoMock package.
package mock_game_logic

import (
	reflect "reflect"
	contract "uwwolf/internal/app/game/logic/contract"
	types "uwwolf/internal/app/game/logic/types"

	gomock "github.com/golang/mock/gomock"
)

// MockModerator is a mock of Moderator interface.
type MockModerator struct {
	ctrl     *gomock.Controller
	recorder *MockModeratorMockRecorder
}

// MockModeratorMockRecorder is the mock recorder for MockModerator.
type MockModeratorMockRecorder struct {
	mock *MockModerator
}

// NewMockModerator creates a new mock instance.
func NewMockModerator(ctrl *gomock.Controller) *MockModerator {
	mock := &MockModerator{ctrl: ctrl}
	mock.recorder = &MockModeratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModerator) EXPECT() *MockModeratorMockRecorder {
	return m.recorder
}

// FinishGame mocks base method.
func (m *MockModerator) FinishGame() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinishGame")
	ret0, _ := ret[0].(bool)
	return ret0
}

// FinishGame indicates an expected call of FinishGame.
func (mr *MockModeratorMockRecorder) FinishGame() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinishGame", reflect.TypeOf((*MockModerator)(nil).FinishGame))
}

// GameID mocks base method.
func (m *MockModerator) GameID() types.GameID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GameID")
	ret0, _ := ret[0].(types.GameID)
	return ret0
}

// GameID indicates an expected call of GameID.
func (mr *MockModeratorMockRecorder) GameID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GameID", reflect.TypeOf((*MockModerator)(nil).GameID))
}

// GameStatus mocks base method.
func (m *MockModerator) GameStatus() types.GameStatusID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GameStatus")
	ret0, _ := ret[0].(types.GameStatusID)
	return ret0
}

// GameStatus indicates an expected call of GameStatus.
func (mr *MockModeratorMockRecorder) GameStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GameStatus", reflect.TypeOf((*MockModerator)(nil).GameStatus))
}

// OnPhaseChanged mocks base method.
func (m *MockModerator) OnPhaseChanged(fn func(contract.Moderator)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnPhaseChanged", fn)
}

// OnPhaseChanged indicates an expected call of OnPhaseChanged.
func (mr *MockModeratorMockRecorder) OnPhaseChanged(fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnPhaseChanged", reflect.TypeOf((*MockModerator)(nil).OnPhaseChanged), fn)
}

// Player mocks base method.
func (m *MockModerator) Player(ID types.PlayerId) contract.Player {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Player", ID)
	ret0, _ := ret[0].(contract.Player)
	return ret0
}

// Player indicates an expected call of Player.
func (mr *MockModeratorMockRecorder) Player(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Player", reflect.TypeOf((*MockModerator)(nil).Player), ID)
}

// RequestPlay mocks base method.
func (m *MockModerator) RequestPlay(playerID types.PlayerId, req *types.ActivateAbilityRequest) *types.ActionResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestPlay", playerID, req)
	ret0, _ := ret[0].(*types.ActionResponse)
	return ret0
}

// RequestPlay indicates an expected call of RequestPlay.
func (mr *MockModeratorMockRecorder) RequestPlay(playerID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestPlay", reflect.TypeOf((*MockModerator)(nil).RequestPlay), playerID, req)
}

// Scheduler mocks base method.
func (m *MockModerator) Scheduler() contract.Scheduler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scheduler")
	ret0, _ := ret[0].(contract.Scheduler)
	return ret0
}

// Scheduler indicates an expected call of Scheduler.
func (mr *MockModeratorMockRecorder) Scheduler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scheduler", reflect.TypeOf((*MockModerator)(nil).Scheduler))
}

// StartGame mocks base method.
func (m *MockModerator) StartGame() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartGame")
	ret0, _ := ret[0].(int64)
	return ret0
}

// StartGame indicates an expected call of StartGame.
func (mr *MockModeratorMockRecorder) StartGame() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartGame", reflect.TypeOf((*MockModerator)(nil).StartGame))
}