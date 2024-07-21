// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: mangas.sql

package database

import (
	"context"
	"database/sql"
)

const getMangaById = `-- name: GetMangaById :one
SELECT id, title, issue_number, publication_date, storyline, cover_art_url, read, user_id, updated_at, images, authors, serializations, genres, explicit_genres, themes, demographics, score, scored_by, rank, popularity, members, favorites, synopsis, background, relations, external_links FROM mangas WHERE id = $1
`

func (q *Queries) GetMangaById(ctx context.Context, id string) (Manga, error) {
	row := q.db.QueryRowContext(ctx, getMangaById, id)
	var i Manga
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.IssueNumber,
		&i.PublicationDate,
		&i.Storyline,
		&i.CoverArtUrl,
		&i.Read,
		&i.UserID,
		&i.UpdatedAt,
		&i.Images,
		&i.Authors,
		&i.Serializations,
		&i.Genres,
		&i.ExplicitGenres,
		&i.Themes,
		&i.Demographics,
		&i.Score,
		&i.ScoredBy,
		&i.Rank,
		&i.Popularity,
		&i.Members,
		&i.Favorites,
		&i.Synopsis,
		&i.Background,
		&i.Relations,
		&i.ExternalLinks,
	)
	return i, err
}

const getUserFavorites = `-- name: GetUserFavorites :many

SELECT mangas.id, mangas.title, mangas.issue_number, mangas.publication_date, mangas.storyline, mangas.cover_art_url, mangas.read, mangas.user_id, mangas.updated_at, mangas.images, mangas.authors, mangas.serializations, mangas.genres, mangas.explicit_genres, mangas.themes, mangas.demographics, mangas.score, mangas.scored_by, mangas.rank, mangas.popularity, mangas.members, mangas.favorites, mangas.synopsis, mangas.background, mangas.relations, mangas.external_links
FROM mangas
JOIN user_mangas ON mangas.id = user_mangas.manga_id
WHERE user_mangas.user_id = $1
`

func (q *Queries) GetUserFavorites(ctx context.Context, userID sql.NullString) ([]Manga, error) {
	rows, err := q.db.QueryContext(ctx, getUserFavorites, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Manga
	for rows.Next() {
		var i Manga
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.IssueNumber,
			&i.PublicationDate,
			&i.Storyline,
			&i.CoverArtUrl,
			&i.Read,
			&i.UserID,
			&i.UpdatedAt,
			&i.Images,
			&i.Authors,
			&i.Serializations,
			&i.Genres,
			&i.ExplicitGenres,
			&i.Themes,
			&i.Demographics,
			&i.Score,
			&i.ScoredBy,
			&i.Rank,
			&i.Popularity,
			&i.Members,
			&i.Favorites,
			&i.Synopsis,
			&i.Background,
			&i.Relations,
			&i.ExternalLinks,
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

const updateStatusRead = `-- name: UpdateStatusRead :exec

UPDATE mangas SET read = $1, updated_at = $2
WHERE id = $3
`

type UpdateStatusReadParams struct {
	Read      sql.NullBool
	UpdatedAt sql.NullTime
	ID        string
}

func (q *Queries) UpdateStatusRead(ctx context.Context, arg UpdateStatusReadParams) error {
	_, err := q.db.ExecContext(ctx, updateStatusRead, arg.Read, arg.UpdatedAt, arg.ID)
	return err
}
