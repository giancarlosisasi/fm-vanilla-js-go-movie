// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: actors.sql

package database

import (
	"context"
)

const getActorByID = `-- name: GetActorByID :one
SELECT id, first_name, last_name, image_url FROM actors WHERE id = $1
`

// Get an actor by their ID
func (q *Queries) GetActorByID(ctx context.Context, id int32) (Actor, error) {
	row := q.db.QueryRowContext(ctx, getActorByID, id)
	var i Actor
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.ImageUrl,
	)
	return i, err
}
