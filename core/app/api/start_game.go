package api

import (
	"net/http"
	"time"
	"uwwolf/game/types"

	"github.com/gin-gonic/gin"
)

// startGame creates a game moderator and then starts the game.
func (s ApiServer) startGame(ctx *gin.Context) {
	playerID := types.PlayerID(ctx.GetString("playerID"))

	room := s.roomService.PlayerWaitingRoom(playerID)
	if room == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "You're not in any room!",
		})
		return
	}

	if playerID != room.OwnerID {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Only the room owner can start the game!",
		})
		return
	}

	config := s.gameService.GameConfig(room.ID)
	mod, err := s.gameService.RegisterGame(&types.GameConfig{
		TurnDuration:       time.Duration(config.TurnDuration) * time.Second,
		DiscussionDuration: time.Duration(config.DiscussionDuration) * time.Second,
		GameInitialization: types.GameInitialization{
			RoleIDs:          config.RoleIDs,
			RequiredRoleIDs:  config.RequiredRoleIDs,
			NumberWerewolves: config.NumberWerewolves,
			PlayerIDs:        room.PlayerIDs,
		},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	mod.StartGame()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Ok",
	})
}
