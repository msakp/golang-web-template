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
	ctx   context.Context
}

func NewUserRepository(db *database.Pg) *userRepo {
	return &userRepo{
		query: db.Queries(),
		ctx:   db.Context(),
	}
}

func (r *userRepo) Create(u *storage.CreateUserParams) (uuid.UUID, error) {
	err := r.query.CreateUser(r.ctx, *u)
	if err != nil{
		return uuid.UUID{}, err
	}
	user, _ := r.query.GetUserByEmail(r.ctx, u.Email)
	return user.ID, nil
}

func (r *userRepo) GetByEmail(email string) (*storage.User, error) {
	u, err := r.query.GetUserByEmail(r.ctx, email)
	return &u, err
}

func (r *userRepo) GetById(id uuid.UUID) (*storage.User, error) {
	u, err := r.query.GetUserById(r.ctx, id)
	return &u, err
}
