package action

import (
	"testing"
	"uwwolf/game/types"
	"uwwolf/game/vars"
	gamemock "uwwolf/mock/game"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type PredictSuite struct {
	suite.Suite
	actorID            types.PlayerID
	targetID           types.PlayerID
	predictedFactionID types.FactionID
	predictedRoleID    types.RoleID
}

func TestPredictSuite(t *testing.T) {
	suite.Run(t, new(PredictSuite))
}

func (ps *PredictSuite) SetupSuite() {
	ps.actorID = "1"
	ps.targetID = "2"
	ps.predictedFactionID = vars.WerewolfFactionID
	ps.predictedRoleID = vars.WerewolfRoleID
}

func (ps *PredictSuite) TestNewFactionPredict() {
	ctrl := gomock.NewController(ps.T())
	defer ctrl.Finish()
	game := gamemock.NewMockGame(ctrl)

	pred := NewFactionPredict(game, ps.predictedFactionID).(*predict)

	ps.Equal(vars.PredictActionID, pred.ID())
	ps.Equal(ps.predictedFactionID, pred.FactionID)
	ps.Len(pred.Faction, 0)
	ps.Equal(types.RoleID(0), pred.RoleID)
	ps.Len(pred.Role, 0)
}

func (ps *PredictSuite) TestNewRolePredict() {
	ctrl := gomock.NewController(ps.T())
	defer ctrl.Finish()
	game := gamemock.NewMockGame(ctrl)

	pred := NewRolePredict(game, ps.predictedRoleID).(*predict)

	ps.Equal(vars.PredictActionID, pred.ID())
	ps.Equal(ps.predictedRoleID, pred.RoleID)
	ps.Len(pred.Role, 0)
	ps.Equal(types.FactionID(0), pred.FactionID)
	ps.Len(pred.Faction, 0)
}

func (ps *PredictSuite) TestValidateFactionPredict() {
	tests := []struct {
		name        string
		req         *types.ActionRequest
		expectedErr string
		setup       func(*predict, *gamemock.MockGame)
	}{
		{
			name: "Invalid (Cant predict yourself)",
			req: &types.ActionRequest{
				ActorID:  ps.actorID,
				TargetID: ps.actorID,
			},
			expectedErr: "WTF! You don't know who you are? (╯°□°)╯︵ ┻━┻",
			setup:       func(p *predict, mg *gamemock.MockGame) {},
		},
		{
			name: "Invalid (Cant predict known player)",
			req: &types.ActionRequest{
				ActorID:  ps.actorID,
				TargetID: ps.targetID,
			},
			expectedErr: "You already knew this player ¯\\(º_o)/¯",
			setup: func(p *predict, gm *gamemock.MockGame) {
				p.Faction[ps.targetID] = true
			},
		},
		{
			name: "Invalid (Cant predict non-existent player)",
			req: &types.ActionRequest{
				ActorID:  ps.actorID,
				TargetID: types.PlayerID("-99"),
			},
			expectedErr: "Non-existent player ¯\\_(ツ)_/¯",
			setup: func(p *predict, gm *gamemock.MockGame) {
				gm.EXPECT().Player(types.PlayerID("-99")).Return(nil).Times(1)
			},
		},
		{
			name: "Ok",
			req: &types.ActionRequest{
				ActorID:  ps.actorID,
				TargetID: ps.targetID,
			},
			setup: func(p *predict, gm *gamemock.MockGame) {
				targetPlayer := gamemock.NewMockPlayer(nil)
				gm.EXPECT().Player(ps.targetID).Return(targetPlayer).Times(1)
			},
		},
	}

	for _, test := range tests {
		ps.Run(test.name, func() {
			ctrl := gomock.NewController(ps.T())
			defer ctrl.Finish()
			game := gamemock.NewMockGame(ctrl)

			pred := NewFactionPredict(game, ps.predictedFactionID).(*predict)
			test.setup(pred, game)
			err := pred.validate(test.req)

			if test.expectedErr == "" {
				ps.Nil(err)
			} else {
				ps.Equal(test.expectedErr, err.Error())
			}
		})
	}
}

func (ps *PredictSuite) TestValidateRolePredict() {
	tests := []struct {
		name        string
		req         *types.ActionRequest
		expectedErr string
		setup       func(*predict, *gamemock.MockGame)
	}{
		{
			name: "Invalid (Cant predict yourself)",
			req: &types.ActionRequest{
				ActorID:  ps.actorID,
				TargetID: ps.actorID,
			},
			expectedErr: "WTF! You don't know who you are? (╯°□°)╯︵ ┻━┻",
			setup:       func(p *predict, mg *gamemock.MockGame) {},
		},
		{
			name: "Invalid (Cant predict known player)",
			req: &types.ActionRequest{
				ActorID:  ps.actorID,
				TargetID: ps.targetID,
			},
			expectedErr: "You already knew this player ¯\\(º_o)/¯",
			setup: func(p *predict, gm *gamemock.MockGame) {
				p.Role[ps.targetID] = true
			},
		},
		{
			name: "Invalid (Cant predict non-existent player)",
			req: &types.ActionRequest{
				ActorID:  ps.actorID,
				TargetID: types.PlayerID("-99"),
			},
			expectedErr: "Non-existent player ¯\\_(ツ)_/¯",
			setup: func(p *predict, gm *gamemock.MockGame) {
				gm.EXPECT().Player(types.PlayerID("-99")).Return(nil).Times(1)
			},
		},
		{
			name: "Ok",
			req: &types.ActionRequest{
				ActorID:  ps.actorID,
				TargetID: ps.targetID,
			},
			setup: func(p *predict, gm *gamemock.MockGame) {
				targetPlayer := gamemock.NewMockPlayer(nil)
				gm.EXPECT().Player(ps.targetID).Return(targetPlayer).Times(1)
			},
		},
	}

	for _, test := range tests {
		ps.Run(test.name, func() {
			ctrl := gomock.NewController(ps.T())
			defer ctrl.Finish()
			game := gamemock.NewMockGame(ctrl)

			pred := NewRolePredict(game, ps.predictedRoleID).(*predict)
			test.setup(pred, game)
			err := pred.validate(test.req)

			if test.expectedErr == "" {
				ps.Nil(err)
			} else {
				ps.Equal(test.expectedErr, err.Error())
			}
		})
	}
}

func (ps *PredictSuite) TestPerformFactionPredict() {
	tests := []struct {
		name        string
		req         *types.ActionRequest
		expectedRes *types.ActionResponse
		setup       func(*predict, *gamemock.MockGame, *gamemock.MockPlayer)
	}{
		{
			name: "Ok (Incorrect prediction)",
			req: &types.ActionRequest{
				ActorID:  ps.actorID,
				TargetID: ps.targetID,
			},
			expectedRes: &types.ActionResponse{
				Ok:   true,
				Data: false,
			},
			setup: func(p *predict, mg *gamemock.MockGame, mp *gamemock.MockPlayer) {
				mp.EXPECT().FactionID().Return(vars.VillagerFactionID).Times(1)
				mp.EXPECT().ID().Return(ps.targetID).Times(2)
			},
		},
		{
			name: "Ok (Correct prediction)",
			req: &types.ActionRequest{
				ActorID:  ps.actorID,
				TargetID: ps.targetID,
			},
			expectedRes: &types.ActionResponse{
				Ok:   true,
				Data: true,
			},
			setup: func(p *predict, mg *gamemock.MockGame, mp *gamemock.MockPlayer) {
				mp.EXPECT().FactionID().Return(ps.predictedFactionID).Times(1)
				mp.EXPECT().ID().Return(ps.targetID).Times(2)
			},
		},
	}

	for _, test := range tests {
		ps.Run(test.name, func() {
			ctrl := gomock.NewController(ps.T())
			defer ctrl.Finish()
			game := gamemock.NewMockGame(ctrl)
			targetPlayer := gamemock.NewMockPlayer(ctrl)

			game.EXPECT().Player(ps.targetID).Return(targetPlayer).AnyTimes()

			pred := NewFactionPredict(game, ps.predictedFactionID).(*predict)
			test.setup(pred, game, targetPlayer)
			res := pred.perform(test.req)

			ps.Equal(test.expectedRes, res)
			ps.Equal(test.expectedRes.Data, pred.Faction[ps.targetID])
		})
	}
}

func (ps *PredictSuite) TestPerformRolePredict() {
	tests := []struct {
		name        string
		req         *types.ActionRequest
		expectedRes *types.ActionResponse
		setup       func(*predict, *gamemock.MockGame, *gamemock.MockPlayer)
	}{
		{
			name: "Ok (Incorrect prediction)",
			req: &types.ActionRequest{
				ActorID:  ps.actorID,
				TargetID: ps.targetID,
			},
			expectedRes: &types.ActionResponse{
				Ok:   true,
				Data: false,
			},
			setup: func(p *predict, mg *gamemock.MockGame, mp *gamemock.MockPlayer) {
				mp.EXPECT().RoleIDs().Return([]types.RoleID{}).Times(1)
				mp.EXPECT().ID().Return(ps.targetID).Times(2)
			},
		},
		{
			name: "Ok (Correct prediction)",
			req: &types.ActionRequest{
				ActorID:  ps.actorID,
				TargetID: ps.targetID,
			},
			expectedRes: &types.ActionResponse{
				Ok:   true,
				Data: true,
			},
			setup: func(p *predict, mg *gamemock.MockGame, mp *gamemock.MockPlayer) {
				mp.EXPECT().RoleIDs().Return([]types.RoleID{ps.predictedRoleID}).Times(1)
				mp.EXPECT().ID().Return(ps.targetID).Times(2)
			},
		},
	}

	for _, test := range tests {
		ps.Run(test.name, func() {
			ctrl := gomock.NewController(ps.T())
			defer ctrl.Finish()
			game := gamemock.NewMockGame(ctrl)
			targetPlayer := gamemock.NewMockPlayer(ctrl)

			game.EXPECT().Player(ps.targetID).Return(targetPlayer).AnyTimes()

			pred := NewRolePredict(game, ps.predictedRoleID).(*predict)
			test.setup(pred, game, targetPlayer)
			res := pred.perform(test.req)

			ps.Equal(test.expectedRes, res)
			ps.Equal(test.expectedRes.Data, pred.Role[ps.targetID])
		})
	}
}