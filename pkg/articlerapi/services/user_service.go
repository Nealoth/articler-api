package services

import (
	"github.com/Nealoth/articler-api/pkg/articlerapi/domain"
	"github.com/Nealoth/articler-api/pkg/articlerapi/repositories"
	"github.com/Nealoth/articler-api/pkg/articlerapi/utils"
)

type UserService struct {
	repoLayer *repositories.RepositoryLayer
	hasher    *utils.Hasher
}

func NewUserService(
	repoLayer *repositories.RepositoryLayer,
	hasher *utils.Hasher,
) *UserService {
	return &UserService{
		repoLayer: repoLayer,
		hasher:    hasher,
	}
}

//TODO change back for entity
func (s *UserService) CreateUser(credentials domain.UserCredentials) (uint64, error) {

	userID, err := s.
		repoLayer.
		UserRepository.
		Create(credentials.Email, s.hasher.Hash(credentials.Password))

	if err != nil {
		return 0, err
	}

	return userID, nil
}

//TODO do we need dto passed by reference?
func (s *UserService) GetUserProfileByID(userID uint64) (*domain.UserProfile, error) {
	userEntity, err := s.repoLayer.UserRepository.GetByID(userID)

	//TODO if sqlnorows -> custom error "client not found" with code
	if err != nil {
		return nil, nil
	}

	return &domain.UserProfile{
		ID:    userEntity.ID,
		Email: userEntity.Email,
	}, nil
}

func (s *UserService) GetUserProfileByCredentials(credentials domain.UserCredentials) (*domain.UserProfile, error) {
	userEntity, err := s.
		repoLayer.
		UserRepository.
		GetByEmailAndPassword(credentials.Email, s.hasher.Hash(credentials.Password))

	//TODO if sqlnorows -> custom error "client not found" with code
	if err != nil {
		return nil, nil
	}

	return &domain.UserProfile{
		ID:    userEntity.ID,
		Email: userEntity.Email,
	}, nil
}
