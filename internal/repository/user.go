package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/msakp/golang-web-template/internal/domain/contracts"
	"github.com/msakp/golang-web-template/internal/infrastructure/database"
	"github.com/msakp/golang-web-template/internal/infrastructure/database/sqlc/storage"
)

var _ contracts.UserRepository = (*userRepo)(nil)

type userRepo struct {
	query *storage.Queries
}

func NewUserRepository(db *database.Pg) *userRepo {
	return &userRepo{
		query: db.Queries(),
	}
}

func (r *userRepo) Create(ctx context.Context, u *storage.CreateUserParams) (uuid.UUID, error) {
	err := r.query.CreateUser(ctx, *u)
	if err != nil{
		return uuid.UUID{}, err
	}
	user, _ := r.query.GetUserByEmail(ctx, u.Email)
	return user.ID, nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*storage.User, error) {
	u, err := r.query.GetUserByEmail(ctx, email)
	return &u, err
}

func (r *userRepo) GetById(ctx context.Context, id uuid.UUID) (*storage.User, error) {
	u, err := r.query.GetUserById(ctx, id)
	return &u, err
}
