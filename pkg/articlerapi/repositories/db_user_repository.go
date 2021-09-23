package repositories

import (
	"context"
	"github.com/Nealoth/articler-api/pkg/articlerapi/domain"
	"github.com/jmoiron/sqlx"
)

type DbUserRepository struct {
	ctx context.Context
	db  *sqlx.DB
}

func NewUserRepository(ctx context.Context, db *sqlx.DB) *DbUserRepository {
	return &DbUserRepository{
		ctx: ctx,
		db:  db,
	}
}

func (r DbUserRepository) Create(email, password string) (uint64, error) {

	tx, err := r.db.BeginTxx(r.ctx, nil)

	if err != nil {
		return 0, err
	}

	row := tx.
		QueryRowxContext(
			r.ctx,
			"INSERT INTO USERS (email, hashed_password) VALUES ($1, $2) RETURNING id",
			email, password,
		)

	var userID uint64

	if err := row.Scan(&userID); err != nil {
		if err := tx.Rollback(); err != nil {
			return 0, err
		}
		return 0, err
	}

	return userID, tx.Commit()
}

func (r DbUserRepository) GetByEmailAndPassword(email string, password string) (*domain.UserEntity, error) {

	var userEntity domain.UserEntity

	err := r.db.
		QueryRowxContext(
			r.ctx,
			"SELECT * FROM users WHERE email=$1 AND hashed_password=$2", email, password).
		StructScan(&userEntity)

	if err != nil {
		return nil, err
	}

	return &userEntity, nil
}

func (r DbUserRepository) GetByID(userID uint64) (*domain.UserEntity, error) {

	var userEntity domain.UserEntity

	err := r.db.
		QueryRowxContext(
			r.ctx,
			"SELECT * FROM users WHERE id=$1", userID).
		StructScan(&userEntity)

	if err != nil {
		return nil, err
	}

	return &userEntity, nil
}
