package game

import (
	"uwwolf/app/game/config"
	"uwwolf/app/game/contract"
	"uwwolf/app/game/types"

	"golang.org/x/exp/slices"
)

type scheduler struct {
	round        types.Round
	beginPhaseID types.PhaseID
	phaseID      types.PhaseID
	turnIndex    int
	phases       map[types.PhaseID][]*types.Turn
}

func NewScheduler(beginPhaseID types.PhaseID) contract.Scheduler {
	return &scheduler{
		round:   config.FirstRound,
		phaseID: beginPhaseID,
		phases:  make(map[types.PhaseID][]*types.Turn),
	}
}

func (s *scheduler) Round() types.Round {
	return s.round
}

func (s *scheduler) PhaseID() types.PhaseID {
	return s.phaseID
}

func (s *scheduler) Phase() []*types.Turn {
	return s.phases[s.phaseID]
}

func (s *scheduler) Turn() *types.Turn {
	if len(s.Phase()) == 0 || s.turnIndex >= len(s.Phase()) {
		return nil
	}

	return s.Phase()[s.turnIndex]
}

func (s *scheduler) IsEmpty(phaseID types.PhaseID) bool {
	if !phaseID.IsUnknown() {
		return len(s.phases[phaseID]) == 0
	}

	for _, p := range s.phases {
		if len(p) != 0 {
			return false
		}
	}

	return true
}

func (s *scheduler) isValidPhaseID(phaseID types.PhaseID) bool {
	return phaseID >= config.NightPhaseID && phaseID <= config.DuskPhaseID
}

func (s *scheduler) existRole(roleID types.RoleID) bool {
	for _, phase := range s.phases {
		if slices.IndexFunc(phase, func(turn *types.Turn) bool {
			return turn.RoleID == roleID
		}) != -1 {
			return true
		}
	}

	return false
}

// Decide which phase and turn index contain new turn. Return -1 in second parameter
// if failed.
func (s *scheduler) calculateTurnIndex(setting *types.TurnSetting) (types.PhaseID, int) {
	turnIndex := -1
	phaseID := setting.PhaseID

	if setting.Position == config.NextPosition {
		phaseID = s.phaseID

		if len(s.Phase()) != 0 {
			turnIndex = s.turnIndex + 1
		} else {
			// Become the first turn if the previous one doesn't exist
			turnIndex = 0
		}
	} else if setting.Position == config.SortedPosition {
		turnIndex = slices.IndexFunc(s.phases[phaseID], func(turn *types.Turn) bool {
			return turn.Priority < setting.Priority
		})

		// Become the first turn if phase is empty or become the last turn
		// if all existed turns have higher priority, respectively
		if turnIndex == -1 {
			turnIndex = len(s.phases[phaseID])
		}
	} else if setting.Position == config.LastPosition {
		turnIndex = len(s.phases[phaseID])
	} else {
		if setting.Position >= 0 && int(setting.Position) <= len(s.phases[phaseID]) {
			turnIndex = int(setting.Position)
		}
	}

	return phaseID, turnIndex
}

func (s *scheduler) AddTurn(setting *types.TurnSetting) bool {
	if !s.isValidPhaseID(setting.PhaseID) || s.existRole(setting.RoleID) {
		return false
	}

	phaseID, turnIndex := s.calculateTurnIndex(setting)

	if turnIndex == -1 {
		return false
	}

	// Increase current turn index by 1 if new turn's position
	// is less than or equal to current turn index
	if s.Turn() != nil && turnIndex <= s.turnIndex && phaseID == s.phaseID {
		s.turnIndex++
	}

	s.phases[phaseID] = slices.Insert(
		s.phases[phaseID],
		turnIndex,
		&types.Turn{
			RoleID:     setting.RoleID,
			BeginRound: setting.BeginRound,
			Priority:   setting.Priority,
		},
	)

	return true
}

// Remove turn. In the case that removed turn is the same as the current
// turn index and phase containing the turn is also  the same as current
// phase, decreases the current turn index by 1.
func (s *scheduler) RemoveTurn(roleID types.RoleID) bool {
	for phaseID, phase := range s.phases {
		for removedTurnIndex, turn := range phase {
			if turn.RoleID == roleID {
				phase = slices.Delete(phase, removedTurnIndex, removedTurnIndex+1)

				// Decrease current turn index by 1 if removed turn's position
				// is less than or equal to current turn index
				if phaseID == s.phaseID && removedTurnIndex <= s.turnIndex {
					s.turnIndex--

					// Move the current turn index to the previous phase's last turn
					// if the current phase is empty
					for s.turnIndex == -1 && !s.IsEmpty(types.PhaseID(0)) {
						s.phaseID--

						// Back to previous round
						if s.phaseID == 0 {
							s.round--
							s.phaseID = config.DuskPhaseID
						}

						s.turnIndex = len(s.Phase()) - 1
					}

					// Reset if current turn index is still -1
					if s.turnIndex == -1 {
						s.phaseID = phaseID
						s.turnIndex = 0
					}
				}

				return true
			}
		}
	}

	return false
}

func (s *scheduler) defrostCurrentTurn() bool {
	turn := s.Turn()

	if turn.FrozenLimit != config.ReachedLimit {
		if turn.FrozenLimit != config.Unlimited {
			turn.FrozenLimit--
		}

		return s.NextTurn(false)
	}

	return true
}

// Move to next turn and delete previous turn if it's times is out.
// Repeat from the beginning if the end is exceeded and return false
// if round is empty.
func (s *scheduler) NextTurn(isRemoved bool) bool {
	if s.IsEmpty(types.PhaseID(0)) {
		return false
	}

	if !isRemoved {
		s.RemoveTurn(s.Turn().RoleID)
	}

	if s.turnIndex < len(s.Phase())-1 {
		s.turnIndex++

		// Skip turn if not the time
		if s.Turn().BeginRound > s.round {
			return s.NextTurn(false)
		}
	} else {
		s.turnIndex = 0
		s.phaseID = (s.phaseID + 1) % (config.DuskPhaseID + 1)

		// Start new round
		if s.phaseID == 0 {
			s.round++
			s.phaseID = config.NightPhaseID
		}

		if s.Turn() == nil {
			return s.NextTurn(false)
		}
	}

	return s.defrostCurrentTurn()
}

func (s *scheduler) FreezeTurn(roleID types.RoleID, limit types.Limit) bool {
	for _, phase := range s.phases {
		for _, turn := range phase {
			if turn.RoleID == roleID {
				turn.FrozenLimit = limit

				return true
			}
		}
	}

	return false
}
