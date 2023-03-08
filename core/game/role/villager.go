package role

import (
	"uwwolf/game/action"
	"uwwolf/game/contract"
	"uwwolf/game/types"
	"uwwolf/game/vars"
)

type villager struct {
	*role
}

func NewVillager(game contract.Game, playerID types.PlayerID) (contract.Role, error) {
	voteAction, err := action.NewVote(game, &action.VoteActionSetting{
		FactionID: vars.VillagerFactionID,
		PlayerID:  playerID,
		Weight:    1,
	})
	if err != nil {
		return nil, err
	}

	return &villager{
		role: &role{
			id:           vars.VillagerRoleID,
			factionID:    vars.VillagerFactionID,
			phaseID:      vars.DayPhaseID,
			beginRoundID: vars.FirstRound,
			turnID:       vars.VillagerTurnID,
			game:         game,
			player:       game.Player(playerID),
			abilities: []*ability{
				{
					action:      voteAction,
					activeLimit: vars.UnlimitedTimes,
				},
			},
		},
	}, nil
}

// RegisterTurn adds role's turn to the game schedule.
func (v villager) RegisterTurn() {
	v.game.Scheduler().AddSlot(&types.NewTurnSlot{
		PhaseID:      v.phaseID,
		TurnID:       v.turnID,
		BeginRoundID: v.beginRoundID,
		PlayerID:     v.player.ID(),
		RoleID:       v.id,
	})
}