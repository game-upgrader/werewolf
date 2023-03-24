package contract

import "uwwolf/game/types"

// Moderator controlls a game.
type Moderator interface {
	// StartGame starts the game.
	StartGame() int64

	// FinishGame ends the game.
	FinishGame() bool

	// RequestPlay receives the play request from the player.
	RequestPlay(playerID types.PlayerID, req *types.ActivateAbilityRequest) *types.ActionResponse
}
