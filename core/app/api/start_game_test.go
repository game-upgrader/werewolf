package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"uwwolf/app/api"
	"uwwolf/app/data"
	"uwwolf/app/enum"
	mock_service "uwwolf/mock/app/service"
	mock_game "uwwolf/mock/game"
	"uwwolf/util"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func (ass ApiServiceSuite) TestStartGame() {
	tests := []struct {
		name  string
		setup func(ctx *gin.Context, gameSvc *mock_service.MockGameService, mod *mock_game.MockModerator)
		check func(res *httptest.ResponseRecorder)
	}{
		{
			name:  "Failure (Forget to use WaitingRoomOwner middleware)",
			setup: func(ctx *gin.Context, gameSvc *mock_service.MockGameService, mod *mock_game.MockModerator) {},
			check: func(res *httptest.ResponseRecorder) {
				ass.Equal(http.StatusInternalServerError, res.Code)
				ass.Equal(
					"Unable to update game config!",
					util.JsonToMap(res.Body.String())["message"],
				)
			},
		},
		{
			name: "Failure (Check falied)",
			setup: func(ctx *gin.Context, gameSvc *mock_service.MockGameService, mod *mock_game.MockModerator) {
				room := &data.WaitingRoom{
					ID: "room_id",
				}
				ctx.Set(enum.WaitingRoomCtxKey, room)

				cfg := data.GameConfig{
					NumberWerewolves: 1,
				}
				gameSvc.EXPECT().GameConfig(room.ID).Return(cfg)
				gameSvc.EXPECT().CheckBeforeRegistration(*room, cfg).Return(fmt.Errorf("Check failed"))
			},
			check: func(res *httptest.ResponseRecorder) {
				ass.Equal(http.StatusBadRequest, res.Code)
				ass.Equal(
					"Check failed",
					util.JsonToMap(res.Body.String())["message"],
				)
			},
		},
		{
			name: "Failure (Register falied)",
			setup: func(ctx *gin.Context, gameSvc *mock_service.MockGameService, mod *mock_game.MockModerator) {
				room := &data.WaitingRoom{
					ID: "room_id",
				}
				ctx.Set(enum.WaitingRoomCtxKey, room)

				cfg := data.GameConfig{
					NumberWerewolves: 1,
				}
				gameSvc.EXPECT().GameConfig(room.ID).Return(cfg)
				gameSvc.EXPECT().CheckBeforeRegistration(*room, cfg)
				gameSvc.EXPECT().RegisterGame(cfg, room.PlayerIDs).
					Return(nil, fmt.Errorf("Register failed"))
			},
			check: func(res *httptest.ResponseRecorder) {
				ass.Equal(http.StatusInternalServerError, res.Code)
				ass.Equal(
					"Register failed",
					util.JsonToMap(res.Body.String())["message"],
				)
			},
		},
		{
			name: "Ok",
			setup: func(ctx *gin.Context, gameSvc *mock_service.MockGameService, mod *mock_game.MockModerator) {
				room := &data.WaitingRoom{
					ID: "room_id",
				}
				ctx.Set(enum.WaitingRoomCtxKey, room)

				cfg := data.GameConfig{
					NumberWerewolves: 1,
				}
				gameSvc.EXPECT().GameConfig(room.ID).Return(cfg)
				gameSvc.EXPECT().CheckBeforeRegistration(*room, cfg)
				gameSvc.EXPECT().RegisterGame(cfg, room.PlayerIDs).Return(mod, nil)
				mod.EXPECT().StartGame()
			},
			check: func(res *httptest.ResponseRecorder) {
				ass.Equal(http.StatusOK, res.Code)
			},
		},
	}

	for _, test := range tests {
		ass.Run(test.name, func() {
			ctrl := gomock.NewController(ass.T())
			defer ctrl.Finish()
			gameSvc := mock_service.NewMockGameService(ctrl)
			mod := mock_game.NewMockModerator(ctrl)

			res := httptest.NewRecorder()
			ctx, r := gin.CreateTestContext(res)

			test.setup(ctx, gameSvc, mod)

			svr := api.NewAPIServer(nil, gameSvc)
			r.POST("/test", func(_ *gin.Context) {
				svr.StartGame(ctx)
			})

			ctx.Request, _ = http.NewRequest(http.MethodPost, "/test", nil)
			r.ServeHTTP(res, ctx.Request)

			test.check(res)
		})
	}
}
