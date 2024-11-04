package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/Masterminds/squirrel"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// Сохранение данных в таблицу
func (r *Repository) SaveData(data map[string]interface{}) error {
	sqlQuery, args, err := squirrel.Insert("data").
		Columns("descr", "text").
		Values(data["descr"], data["text"]).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(context.Background(), sqlQuery, args...)
	if err != nil {
		log.Printf("Ошибка сохранения данных: %v", err)
		return err
	}
	return nil
}

// Получение данных из таблицы
func (r *Repository) GetDataByQuery(query string) (map[string]interface{}, error) {
	sqlQuery, args, err := squirrel.Select("descr", "text").
		From("data").
		Where(squirrel.Eq{"text": query}). // Убедитесь, что колонка правильно названа как "text"
		ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRowContext(context.Background(), sqlQuery, args...)
	var descr, text string
	if err := row.Scan(&descr, &text); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // если данных нет, возвращаем nil
		}
		log.Printf("Ошибка при сканировании данных: %v", err)
		return nil, err
	}

	return map[string]interface{}{
		"descr": descr,
		"text":  text,
	}, nil
}
