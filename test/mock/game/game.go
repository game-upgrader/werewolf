// Code generated by MockGen. DO NOT EDIT.
// Source: module/game/contract/game.go

// Package game is a generated GoMock package.
package game

import (
	reflect "reflect"
	contract "uwwolf/module/game/contract"
	state "uwwolf/module/game/state"
	types "uwwolf/types"

	gomock "github.com/golang/mock/gomock"
)

// MockGame is a mock of Game interface.
type MockGame struct {
	ctrl     *gomock.Controller
	recorder *MockGameMockRecorder
}

// MockGameMockRecorder is the mock recorder for MockGame.
type MockGameMockRecorder struct {
	mock *MockGame
}

// NewMockGame creates a new mock instance.
func NewMockGame(ctrl *gomock.Controller) *MockGame {
	mock := &MockGame{ctrl: ctrl}
	mock.recorder = &MockGameMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGame) EXPECT() *MockGameMockRecorder {
	return m.recorder
}

// IsStarted mocks base method.
func (m *MockGame) IsStarted() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsStarted")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsStarted indicates an expected call of IsStarted.
func (mr *MockGameMockRecorder) IsStarted() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsStarted", reflect.TypeOf((*MockGame)(nil).IsStarted))
}

// KillPlayer mocks base method.
func (m *MockGame) KillPlayer(playerId types.PlayerId) contract.Player {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KillPlayer", playerId)
	ret0, _ := ret[0].(contract.Player)
	return ret0
}

// KillPlayer indicates an expected call of KillPlayer.
func (mr *MockGameMockRecorder) KillPlayer(playerId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KillPlayer", reflect.TypeOf((*MockGame)(nil).KillPlayer), playerId)
}

// Player mocks base method.
func (m *MockGame) Player(playerId types.PlayerId) contract.Player {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Player", playerId)
	ret0, _ := ret[0].(contract.Player)
	return ret0
}

// Player indicates an expected call of Player.
func (mr *MockGameMockRecorder) Player(playerId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Player", reflect.TypeOf((*MockGame)(nil).Player), playerId)
}

// PlayerIdsWithFaction mocks base method.
func (m *MockGame) PlayerIdsWithFaction(factionId types.FactionId) []types.PlayerId {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlayerIdsWithFaction", factionId)
	ret0, _ := ret[0].([]types.PlayerId)
	return ret0
}

// PlayerIdsWithFaction indicates an expected call of PlayerIdsWithFaction.
func (mr *MockGameMockRecorder) PlayerIdsWithFaction(factionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlayerIdsWithFaction", reflect.TypeOf((*MockGame)(nil).PlayerIdsWithFaction), factionId)
}

// PlayerIdsWithRole mocks base method.
func (m *MockGame) PlayerIdsWithRole(roleId types.RoleId) []types.PlayerId {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlayerIdsWithRole", roleId)
	ret0, _ := ret[0].([]types.PlayerId)
	return ret0
}

// PlayerIdsWithRole indicates an expected call of PlayerIdsWithRole.
func (mr *MockGameMockRecorder) PlayerIdsWithRole(roleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlayerIdsWithRole", reflect.TypeOf((*MockGame)(nil).PlayerIdsWithRole), roleId)
}

// Poll mocks base method.
func (m *MockGame) Poll(factionId types.FactionId) *state.Poll {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Poll", factionId)
	ret0, _ := ret[0].(*state.Poll)
	return ret0
}

// Poll indicates an expected call of Poll.
func (mr *MockGameMockRecorder) Poll(factionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Poll", reflect.TypeOf((*MockGame)(nil).Poll), factionId)
}

// RequestAction mocks base method.
func (m *MockGame) RequestAction(req *types.ActionRequest) *types.ActionResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestAction", req)
	ret0, _ := ret[0].(*types.ActionResponse)
	return ret0
}

// RequestAction indicates an expected call of RequestAction.
func (mr *MockGameMockRecorder) RequestAction(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestAction", reflect.TypeOf((*MockGame)(nil).RequestAction), req)
}

// Round mocks base method.
func (m *MockGame) Round() *state.Round {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Round")
	ret0, _ := ret[0].(*state.Round)
	return ret0
}

// Round indicates an expected call of Round.
func (mr *MockGameMockRecorder) Round() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Round", reflect.TypeOf((*MockGame)(nil).Round))
}

// Start mocks base method.
func (m *MockGame) Start() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockGameMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockGame)(nil).Start))
}
