package storage

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Repo struct{ DB *sql.DB }

func New(dsn string) (*Repo, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Repo{DB: db}, nil
}

func (r *Repo) CreateUser(ctx context.Context, email string) (int64, error) {
	var id int64
	err := r.DB.QueryRowContext(ctx, `insert into users(email) values($1) returning id`, email).Scan(&id)
	return id, err
}

func (r *Repo) AddNote(ctx context.Context, userID int64, body string) (int64, error) {
	var id int64
	err := r.DB.QueryRowContext(ctx, `insert into notes(user_id, body) values($1,$2) returning id`, userID, body).Scan(&id)
	return id, err
}

func (r *Repo) GetNotes(ctx context.Context, userID int64) ([]string, error) {
	rows, err := r.DB.QueryContext(ctx, `select body from notes where user_id=$1 order by id`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []string
	for rows.Next() {
		var b string
		if err := rows.Scan(&b); err != nil {
			return nil, err
		}
		out = append(out, b)
	}
	return out, rows.Err()
}
