package domain

import "context"

type Role string

const (
	RoleAdmin Role = "admin"
	RoleKasir Role = "kasir"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
	Role     Role
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}
