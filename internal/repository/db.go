package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) SaveData(data map[string]interface{}) error {
	sql, args, err := squirrel.Insert("ai_data").
		Columns("description", "text").
		Values(data["description"], data["text"]).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.pool.Exec(context.Background(), sql, args...)
	return err
}
