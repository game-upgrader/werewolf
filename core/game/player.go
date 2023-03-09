package game

import (
	"errors"
	"fmt"
	"uwwolf/game/contract"
	"uwwolf/game/types"
	"uwwolf/game/vars"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

// player represents the player in a game.
type player struct {
	id         types.PlayerID
	factionID  types.FactionID
	mainRoleID types.RoleID
	isDead     bool
	game       contract.Game
	roles      map[types.RoleID]contract.Role
}

func NewPlayer(game contract.Game, id types.PlayerID) contract.Player {
	return &player{
		id:        id,
		game:      game,
		factionID: vars.VillagerFactionID,
		roles:     make(map[types.RoleID]contract.Role),
	}
}

// ID returns player's ID.
func (p player) ID() types.PlayerID {
	return p.id
}

// MainRoleID returns player's main role id.
func (p player) MainRoleID() types.RoleID {
	return p.mainRoleID
}

// RoleIDs returns player's assigned role ids.
func (p player) RoleIDs() []types.RoleID {
	return maps.Keys(p.roles)
}

// Roles returns player's assigned roles.
func (p player) Roles() map[types.RoleID]contract.Role {
	return p.roles
}

// FactionID returns player's faction ID.
func (p player) FactionID() types.FactionID {
	return p.factionID
}

// IsDead checks if player is dead.
func (p player) IsDead() bool {
	return p.isDead
}

// SetFactionID assigns this player to the new faction.
func (p *player) SetFactionID(factionID types.FactionID) {
	p.factionID = factionID
}

// Die marks this player as dead and triggers roles events.
// If `isExited` is true, any trigger preventing death is ignored.
func (p *player) Die(isExited bool) bool {
	if p.isDead {
		return false
	}

	for _, role := range p.roles {
		if isDead := role.BeforeDeath(); !isDead && !isExited {
			return false
		}
	}

	p.isDead = true
	for _, role := range p.roles {
		role.AfterDeath()
	}

	return true
}

// AssignRole assigns the role to the player, and the faction can
// be updated based on this role.
func (p *player) AssignRole(roleID types.RoleID) (bool, error) {
	if slices.Contains(p.RoleIDs(), roleID) {
		return false, fmt.Errorf("This role is already assigned ¯\\_(ツ)_/¯")
	}

	if newRole, err := NewRole(roleID, p.game, p.id); err != nil {
		return false, err
	} else {
		p.roles[roleID] = newRole
		if vars.RoleWeights[newRole.ID()] > vars.RoleWeights[p.mainRoleID] {
			p.mainRoleID = newRole.ID()
			p.factionID = newRole.FactionID()
		}
		newRole.RegisterSlot()
	}

	return true, nil
}

// RevokeRole removes the role from the player, and the faction can
// be updated based on removed role.
func (p *player) RevokeRole(roleID types.RoleID) (bool, error) {
	if len(p.roles) == 1 {
		return false, errors.New("Player must player at least one role ヾ(⌐■_■)ノ♪")
	} else if p.roles[roleID] == nil {
		return false, errors.New("Non-existent role ID  ¯\\_(ツ)_/¯")
	}

	p.roles[roleID].UnregisterSlot()
	delete(p.roles, roleID)

	if roleID == p.mainRoleID {
		var newMainRole contract.Role

		for _, role := range p.roles {
			if newMainRole == nil ||
				vars.RoleWeights[role.ID()] > vars.RoleWeights[newMainRole.ID()] {
				newMainRole = role
			}
		}

		p.mainRoleID = newMainRole.ID()
		p.factionID = newMainRole.FactionID()
	}

	return true, nil
}

// ActivateAbility executes one of player's available ability.
// The executed ability is selected based on the requested
// action.
func (p *player) ActivateAbility(req *types.ActivateAbilityRequest) *types.ActionResponse {
	if p.isDead {
		return &types.ActionResponse{
			Ok:      false,
			Message: "You're died (╥﹏╥)",
		}
	} else if !p.game.Scheduler().CanPlay(p.id) {
		return &types.ActionResponse{
			Ok:      false,
			Message: "Wait for your turn, OK??",
		}
	} else {
		turn := p.game.Scheduler().Turn()
		return p.roles[turn[p.id].RoleID].ActivateAbility(req)
	}
}
