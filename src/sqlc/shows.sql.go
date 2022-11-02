// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: shows.sql

package sqlc

import (
	"context"
)

const getShowByID = `-- name: GetShowByID :one
SELECT id, type, legacy_id, title, description, image, default_episode FROM shows WHERE id = ?
`

func (q *Queries) GetShowByID(ctx context.Context, id int64) (Show, error) {
	row := q.db.QueryRowContext(ctx, getShowByID, id)
	var i Show
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.LegacyID,
		&i.Title,
		&i.Description,
		&i.Image,
		&i.DefaultEpisode,
	)
	return i, err
}

const listShows = `-- name: ListShows :many
SELECT id, type, legacy_id, title, description, image, default_episode FROM shows
`

func (q *Queries) ListShows(ctx context.Context) ([]Show, error) {
	rows, err := q.db.QueryContext(ctx, listShows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Show
	for rows.Next() {
		var i Show
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.LegacyID,
			&i.Title,
			&i.Description,
			&i.Image,
			&i.DefaultEpisode,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
