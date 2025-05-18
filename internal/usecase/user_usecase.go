package usecase

import (
	"context"
	"pos-kasir/internal/domain"
	"pos-kasir/utils"
)

type UserUsecase struct {
	repo domain.UserRepository
	hash utils.Hasher
	jwt  utils.JWTService
}

func NewUserUsecase(r domain.UserRepository, h utils.Hasher, j utils.JWTService) *UserUsecase {
	return &UserUsecase{repo: r, hash: h, jwt: j}
}

func (u *UserUsecase) Register(ctx context.Context, user *domain.User) error {
	hash, err := u.hash.Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	return u.repo.Create(ctx, user)
}

func (u *UserUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if !u.hash.Compare(user.Password, password) {
		return "", utils.ErrInvalidCredentials
	}
	return u.jwt.Generate(user.ID, string(user.Role))
}
