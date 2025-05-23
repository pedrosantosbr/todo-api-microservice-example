// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: tasks.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const DeleteTask = `-- name: DeleteTask :one
DELETE FROM
  tasks
WHERE
  id = $1
RETURNING id AS res
`

func (q *Queries) DeleteTask(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, DeleteTask, id)
	var res uuid.UUID
	err := row.Scan(&res)
	return res, err
}

const InsertTask = `-- name: InsertTask :one
INSERT INTO tasks (
  description,
  priority,
  start_date,
  due_date
)
VALUES (
  $1,
  $2,
  $3,
  $4
)
RETURNING id
`

type InsertTaskParams struct {
	Description string
	Priority    Priority
	StartDate   pgtype.Timestamp
	DueDate     pgtype.Timestamp
}

func (q *Queries) InsertTask(ctx context.Context, arg InsertTaskParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, InsertTask,
		arg.Description,
		arg.Priority,
		arg.StartDate,
		arg.DueDate,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const SelectTask = `-- name: SelectTask :one
SELECT
  id,
  description,
  priority,
  start_date,
  due_date,
  done
FROM
  tasks
WHERE
  id = $1
LIMIT 1
`

func (q *Queries) SelectTask(ctx context.Context, id uuid.UUID) (Tasks, error) {
	row := q.db.QueryRow(ctx, SelectTask, id)
	var i Tasks
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.Priority,
		&i.StartDate,
		&i.DueDate,
		&i.Done,
	)
	return i, err
}

const UpdateTask = `-- name: UpdateTask :one
UPDATE tasks SET
  description = $1,
  priority    = $2,
  start_date  = $3,
  due_date    = $4,
  done        = $5
WHERE id = $6
RETURNING id AS res
`

type UpdateTaskParams struct {
	Description string
	Priority    Priority
	StartDate   pgtype.Timestamp
	DueDate     pgtype.Timestamp
	Done        bool
	ID          uuid.UUID
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, UpdateTask,
		arg.Description,
		arg.Priority,
		arg.StartDate,
		arg.DueDate,
		arg.Done,
		arg.ID,
	)
	var res uuid.UUID
	err := row.Scan(&res)
	return res, err
}
