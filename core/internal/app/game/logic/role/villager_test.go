package role

// import (
// 	"fmt"
// 	"testing"
// 	"uwwolf/internal/app/game/logic/types"
// 	"uwwolf/game/vars"
// 	mock_game "uwwolf/mock/game"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/suite"
// )

// type VillagerSuite struct {
// 	suite.Suite
// 	playerID types.PlayerId
// }

// func TestVillagerSuite(t *testing.T) {
// 	suite.Run(t, new(VillagerSuite))
// }

// func (vs *VillagerSuite) SetupSuite() {
// 	vs.playerID = types.PlayerId("1")
// }

// func (vs VillagerSuite) TestNewVillager() {
// 	tests := []struct {
// 		name        string
// 		expectedErr error
// 		setup       func(*mock_game.MockGame, *mock_game.MockPoll)
// 	}{
// 		{
// 			name:        "Failure (Poll does not exist)",
// 			expectedErr: fmt.Errorf("Poll does not exist ¯\\_(ツ)_/¯"),
// 			setup: func(mg *mock_game.MockGame, mp *mock_game.MockPoll) {
// 				mg.EXPECT().Poll(vars.VillagerFactionID).Return(nil)
// 			},
// 		},
// 		{
// 			name: "Ok",
// 			setup: func(mg *mock_game.MockGame, mp *mock_game.MockPoll) {
// 				mg.EXPECT().Poll(vars.VillagerFactionID).Return(mp).Times(2)
// 				mp.EXPECT().AddElectors(vs.playerID)
// 				mp.EXPECT().SetWeight(vs.playerID, uint(1))
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		vs.Run(test.name, func() {
// 			ctrl := gomock.NewController(vs.T())
// 			defer ctrl.Finish()
// 			game := mock_game.NewMockGame(ctrl)
// 			player := mock_game.NewMockPlayer(ctrl)
// 			poll := mock_game.NewMockPoll(ctrl)

// 			game.EXPECT().Player(vs.playerID).Return(player).AnyTimes()
// 			test.setup(game, poll)

// 			v, err := NewVillager(game, vs.playerID)

// 			if test.expectedErr != nil {
// 				vs.Nil(v)
// 				vs.NotNil(err)
// 				vs.Equal(test.expectedErr, err)
// 			} else {
// 				vs.Nil(err)
// 				vs.Equal(vars.VillagerRoleID, v.Id())
// 				vs.Equal(vars.DayPhaseID, v.(*villager).phaseID)
// 				vs.Equal(vars.VillagerFactionID, v.FactionID())
// 				vs.Equal(vars.FirstRound, v.(*villager).beginRoundID)
// 				vs.Equal(player, v.(*villager).player)
// 				vs.Equal(vars.UnlimitedTimes, v.ActiveTimes(0))
// 				vs.Len(v.(*villager).abilities, 1)
// 				vs.Equal(vars.VoteActionID, v.(*villager).abilities[0].action.Id())
// 			}
// 		})
// 	}
// }

// func (vs VillagerSuite) TestOnRevoke() {
// 	ctrl := gomock.NewController(vs.T())
// 	defer ctrl.Finish()
// 	game := mock_game.NewMockGame(ctrl)
// 	player := mock_game.NewMockPlayer(ctrl)
// 	scheduler := mock_game.NewMockScheduler(ctrl)
// 	poll := mock_game.NewMockPoll(ctrl)

// 	// Mock for New fuction
// 	game.EXPECT().Scheduler().Return(scheduler)
// 	game.EXPECT().Player(vs.playerID).Return(player)
// 	game.EXPECT().Poll(vars.VillagerFactionID).Return(poll).Times(2)
// 	poll.EXPECT().AddElectors(vs.playerID)
// 	poll.EXPECT().SetWeight(vs.playerID, uint(1))

// 	game.EXPECT().Poll(vars.VillagerFactionID).Return(poll).Times(2)
// 	game.EXPECT().Poll(vars.WerewolfFactionID).Return(poll)
// 	poll.EXPECT().RemoveElector(vs.playerID)
// 	poll.EXPECT().RemoveCandidate(vs.playerID).Times(2)
// 	player.EXPECT().Id().Return(vs.playerID).Times(4)

// 	v, _ := NewVillager(game, vs.playerID)

// 	scheduler.EXPECT().RemoveSlot(&types.RemovedTurnSlot{
// 		PhaseID:  v.(*villager).phaseID,
// 		PlayerID: vs.playerID,
// 		RoleID:   v.Id(),
// 	})

// 	v.OnRevoke()
// }

// func (vs VillagerSuite) TestOnAssign() {
// 	ctrl := gomock.NewController(vs.T())
// 	defer ctrl.Finish()
// 	game := mock_game.NewMockGame(ctrl)
// 	player := mock_game.NewMockPlayer(ctrl)
// 	scheduler := mock_game.NewMockScheduler(ctrl)
// 	poll := mock_game.NewMockPoll(ctrl)

// 	// Mock for New fuction
// 	game.EXPECT().Scheduler().Return(scheduler)
// 	game.EXPECT().Player(vs.playerID).Return(player)
// 	game.EXPECT().Poll(vars.VillagerFactionID).Return(poll).Times(2)
// 	poll.EXPECT().AddElectors(vs.playerID)
// 	poll.EXPECT().SetWeight(vs.playerID, uint(1))

// 	game.EXPECT().Poll(vars.VillagerFactionID).Return(poll)
// 	game.EXPECT().Poll(vars.WerewolfFactionID).Return(poll)
// 	poll.EXPECT().AddCandidates(vs.playerID).Times(2)
// 	player.EXPECT().Id().Return(vs.playerID).Times(3)

// 	v, _ := NewVillager(game, vs.playerID)

// 	scheduler.EXPECT().AddSlot(&types.NewTurnSlot{
// 		PhaseID:      v.(*villager).phaseID,
// 		TurnID:       v.(*villager).turnID,
// 		BeginRoundID: v.(*villager).beginRoundID,
// 		PlayerID:     vs.playerID,
// 		RoleID:       v.Id(),
// 	})

// 	v.OnAssign()
// }
