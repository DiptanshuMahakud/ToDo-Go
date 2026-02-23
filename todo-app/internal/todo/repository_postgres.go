package todo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepo struct {
	db *pgxpool.Pool
}

func NewPostgresRepo(db *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) Create(ctx context.Context, title string) (Todo, error) {
	query := `
		INSERT INTO todos (title)
		VALUES ($1)
		RETURNING id , title , completed , created_at
	`

	var t Todo

	err := r.db.QueryRow(ctx, query, title).Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt)

	return t, err

}

func (r *PostgresRepo) GetAll(ctx context.Context) ([]Todo, error) {
	rows, err := r.db.Query(ctx, `SELECT id , title , completed , created_at FROM todos`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var t Todo
		if err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Completed,
			&t.CreatedAt,
		); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	return todos, nil
}
