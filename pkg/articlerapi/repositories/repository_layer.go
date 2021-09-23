package repositories

import (
	"context"
	"github.com/Nealoth/articler-api/pkg/articlerapi/configuration"
	"github.com/Nealoth/articler-api/pkg/articlerapi/domain"
	"github.com/jmoiron/sqlx"
	"time"
)

type IUserRepository interface {
	Create(email, password string) (uint64, error)
	GetByEmailAndPassword(email string, password string) (*domain.UserEntity, error)
	GetByID(userID uint64) (*domain.UserEntity, error)
}

type ITokenStorageRepository interface {
	StoreToken(userID uint64, token string, expiresAt time.Time) error
	GetToken(userID uint64) (StoredTokenValue, error)
	DeleteToken(userID uint64) error
}

type RepositoryLayer struct {
	UserRepository             IUserRepository
	AuthTokenStorageRepository ITokenStorageRepository
}

func InitDbRepositoryLayer(ctx context.Context, dbConf configuration.DatabaseConfiguration, db *sqlx.DB) *RepositoryLayer {
	return &RepositoryLayer{
		UserRepository:             NewUserRepository(ctx, db),
		AuthTokenStorageRepository: NewInMemoryTokenStorageRepository(),
	}
}
