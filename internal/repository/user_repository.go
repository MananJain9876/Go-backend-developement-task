package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// UserDBModel represents the user as stored in the database.
type UserDBModel struct {
	ID   int64
	Name string
	DOB  time.Time
}

// UserRepository defines methods to interact with users in the database.
type UserRepository interface {
	CreateUser(ctx context.Context, name string, dob time.Time) (UserDBModel, error)
	GetUser(ctx context.Context, id int64) (UserDBModel, error)
	ListUsers(ctx context.Context) ([]UserDBModel, error)
	UpdateUser(ctx context.Context, id int64, name string, dob time.Time) (UserDBModel, error)
	DeleteUser(ctx context.Context, id int64) error
}

type userRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) UserRepository {
	return &userRepository{pool: pool}
}

func (r *userRepository) CreateUser(ctx context.Context, name string, dob time.Time) (UserDBModel, error) {
	var u UserDBModel
	err := r.pool.QueryRow(ctx, `INSERT INTO users (name, dob) VALUES ($1, $2) RETURNING id, name, dob`, name, dob).
		Scan(&u.ID, &u.Name, &u.DOB)
	return u, err
}

func (r *userRepository) GetUser(ctx context.Context, id int64) (UserDBModel, error) {
	var u UserDBModel
	err := r.pool.QueryRow(ctx, `SELECT id, name, dob FROM users WHERE id = $1`, id).
		Scan(&u.ID, &u.Name, &u.DOB)
	return u, err
}

func (r *userRepository) ListUsers(ctx context.Context) ([]UserDBModel, error) {
	rows, err := r.pool.Query(ctx, `SELECT id, name, dob FROM users ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []UserDBModel
	for rows.Next() {
		var u UserDBModel
		if err := rows.Scan(&u.ID, &u.Name, &u.DOB); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return users, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, id int64, name string, dob time.Time) (UserDBModel, error) {
	var u UserDBModel
	err := r.pool.QueryRow(ctx, `UPDATE users SET name = $2, dob = $3 WHERE id = $1 RETURNING id, name, dob`, id, name, dob).
		Scan(&u.ID, &u.Name, &u.DOB)
	return u, err
}

func (r *userRepository) DeleteUser(ctx context.Context, id int64) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM users WHERE id = $1`, id)
	return err
}


