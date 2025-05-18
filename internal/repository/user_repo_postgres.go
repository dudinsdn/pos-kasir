package repository

import (
	"context"
	"database/sql"
	"pos-kasir/internal/domain"

	"github.com/google/uuid"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) domain.UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user *domain.User) error {
	user.ID = uuid.New().String()
	_, err := r.db.ExecContext(ctx, `INSERT INTO users (id, name, email, password, role) VALUES ($1, $2, $3, $4, $5)`,
		user.ID, user.Name, user.Email, user.Password, user.Role)
	return err
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	row := r.db.QueryRowContext(ctx, `SELECT id, name, email, password, role FROM users WHERE email=$1`, email)
	var user domain.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role); err != nil {
		return nil, err
	}
	return &user, nil
}
