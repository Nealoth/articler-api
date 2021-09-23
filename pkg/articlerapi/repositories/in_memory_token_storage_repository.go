package repositories

import (
	"fmt"
	"time"
)

type InMemoryTokenStorageRepository struct {
	storage map[uint64]StoredTokenValue
}

type StoredTokenValue struct {
	Token     string
	ExpiredAt time.Time
}

func NewInMemoryTokenStorageRepository() *InMemoryTokenStorageRepository {
	return &InMemoryTokenStorageRepository{
		storage: make(map[uint64]StoredTokenValue),
	}
}

func (r *InMemoryTokenStorageRepository) StoreToken(userID uint64, token string, expiresAt time.Time) error {
	r.storage[userID] = StoredTokenValue{
		Token:     token,
		ExpiredAt: expiresAt,
	}

	return nil
}

func (r *InMemoryTokenStorageRepository) GetToken(userID uint64) (StoredTokenValue, error) {

	token, found := r.storage[userID]

	if !found {
		return token, fmt.Errorf("cannot found token by key: %d", userID)
	}

	return token, nil
}

func (r *InMemoryTokenStorageRepository) DeleteToken(userID uint64) error {
	delete(r.storage, userID)
	return nil
}
