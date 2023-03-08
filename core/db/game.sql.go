// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: game.sql

package db

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const assignGameRoles = `-- name: AssignGameRoles :exec
INSERT INTO role_assignments (
    game_id,
    role_id,
    faction_id,
    player_id
) VALUES (
    unnest($1::text[]),
    unnest($2::text[]),
    unnest($3::text[]),
    unnest($4::text[])
)
`

type AssignGameRolesParams struct {
	GameIds    []string `json:"game_ids"`
	RoleIds    []string `json:"role_ids"`
	FactionIds []string `json:"faction_ids"`
	PlayerIds  []string `json:"player_ids"`
}

func (q *Queries) AssignGameRoles(ctx context.Context, arg AssignGameRolesParams) error {
	_, err := q.exec(ctx, q.assignGameRolesStmt, assignGameRoles,
		pq.Array(arg.GameIds),
		pq.Array(arg.RoleIds),
		pq.Array(arg.FactionIds),
		pq.Array(arg.PlayerIds),
	)
	return err
}

const createGame = `-- name: CreateGame :one
INSERT INTO games DEFAULT VALUES
RETURNING id, winning_faction_id, created_at, finished_at
`

func (q *Queries) CreateGame(ctx context.Context) (Game, error) {
	row := q.queryRow(ctx, q.createGameStmt, createGame)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.WinningFactionID,
		&i.CreatedAt,
		&i.FinishedAt,
	)
	return i, err
}

const createGameLogs = `-- name: CreateGameLogs :exec
INSERT INTO game_logs (
    game_id,
    round_id,
    actor_id,
    role_id,
    action_id,
    target_id
) VALUES (
    unnest($1::text[]),
    unnest($2::text[]),
    unnest($3::text[]),
    unnest($4::text[]),
    unnest($5::text[]),
    unnest($6::text[])
)
`

type CreateGameLogsParams struct {
	GameIds   []string `json:"game_ids"`
	RoundIds  []string `json:"round_ids"`
	ActorIds  []string `json:"actor_ids"`
	RoleIds   []string `json:"role_ids"`
	ActionIds []string `json:"action_ids"`
	TargetIds []string `json:"target_ids"`
}

func (q *Queries) CreateGameLogs(ctx context.Context, arg CreateGameLogsParams) error {
	_, err := q.exec(ctx, q.createGameLogsStmt, createGameLogs,
		pq.Array(arg.GameIds),
		pq.Array(arg.RoundIds),
		pq.Array(arg.ActorIds),
		pq.Array(arg.RoleIds),
		pq.Array(arg.ActionIds),
		pq.Array(arg.TargetIds),
	)
	return err
}

const finishGame = `-- name: FinishGame :exec
UPDATE games SET
    winning_faction_id = $2,
    finished_at = CURRENT_TIMESTAMP
WHERE id = $1
`

type FinishGameParams struct {
	ID               int64         `json:"id"`
	WinningFactionID sql.NullInt16 `json:"winning_faction_id"`
}

func (q *Queries) FinishGame(ctx context.Context, arg FinishGameParams) error {
	_, err := q.exec(ctx, q.finishGameStmt, finishGame, arg.ID, arg.WinningFactionID)
	return err
}