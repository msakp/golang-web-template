package repository

import (
	"context"

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

func (r *userRepo) Create(u *storage.CreateUserParams) error {

	return r.query.CreateUser(r.ctx, *u)
}

func (r *userRepo) Get(email string) (*storage.User, error) {
	u, err := r.query.GetUser(r.ctx, email)
	return &u, err
}
