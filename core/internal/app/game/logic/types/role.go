package types

// FactionID is ID of faction.
type FactionId = uint8

// RoleID is ID type of role.
type RoleId = uint8

// ActivateAbilityRequest contains information for ability activating.
type ActivateAbilityRequest struct {
	// AbilityIndex is activated ability index.
	AbilityIndex uint8

	// TargetID  is player ID of target player.
	TargetID PlayerId

	// IsSkipped marks that the request is ignored.
	IsSkipped bool
}

// ActionResponse contains action execution's result.
type RoleResponse struct {
	// RoundID is round ID which the action is executed.
	RoundID

	// RoleID is ID of role executing the action.
	RoleId

	// ActionResponse is result of action execution.
	ActionResponse
}
