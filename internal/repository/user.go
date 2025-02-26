package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/msakp/golang-web-template/internal/database/storage"
	"github.com/msakp/golang-web-template/internal/domain/contracts"
	"github.com/msakp/golang-web-template/internal/domain/dto"
	"github.com/msakp/golang-web-template/internal/wrapper"
	"github.com/msakp/golang-web-template/pkg/connections/postgres"
)

var _ contracts.UserRepository = (*userRepo)(nil)

type userRepo struct {
	query *storage.Queries
	pool  *pgxpool.Pool
}

func NewUserRepository(db *postgres.DB) *userRepo {
	return &userRepo{
		query: db.Queries(),
		pool:  db.Pool(),
	}
}

func (r *userRepo) Exists(ctx context.Context, email string) bool {
	ok, _ := r.query.UserExists(ctx, r.pool, email)
	return ok
}

func (r *userRepo) Create(ctx context.Context, u *dto.UserRegister) (uuid.UUID, error) {
	params := wrapper.UserRegisterWithCreateParams(u)
	err := r.query.CreateUser(ctx, r.pool, *params)
	if err != nil {
		return uuid.UUID{}, err
	}
	user, _ := r.query.GetUserByEmail(ctx, r.pool, u.Email)
	return user.ID, nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*dto.UserView, error) {
	u, err := r.query.GetUserByEmail(ctx, r.pool, email)
	return wrapper.UserWithView(&u), err
}

func (r *userRepo) GetById(ctx context.Context, id uuid.UUID) (*dto.UserView, error) {
	u, err := r.query.GetUserById(ctx, r.pool, id)
	return wrapper.UserWithView(&u), err
}
