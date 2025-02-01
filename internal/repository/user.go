package repository

import (
	"context"

	contracts "github.com/msakp/golang-web-template/domain/cotracts"
	"github.com/msakp/golang-web-template/domain/dto"
	"github.com/msakp/golang-web-template/internal/infrastructure/bind"
	"github.com/msakp/golang-web-template/internal/infrastructure/database"
	"github.com/msakp/golang-web-template/internal/infrastructure/database/sqlc/storage"
)

var _ contracts.UserRepository = (*userRepo)(nil)

type userRepo struct{
	query *storage.Queries
	ctx context.Context
}


func NewUserRepository(db *database.Pg) *userRepo{
	return &userRepo{
		query: db.Queries(),
		ctx: db.Context(),
	}
}


func (r *userRepo) Create(u *dto.UserRegister) error{
	p := bind.WithUserRegister(u)	
	return r.query.CreateUser(r.ctx, p)
}
