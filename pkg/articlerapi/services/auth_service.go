package services

import (
	"errors"
	"github.com/Nealoth/articler-api/pkg/articlerapi/apierr"
	"github.com/Nealoth/articler-api/pkg/articlerapi/repositories"
	"github.com/Nealoth/articler-api/pkg/articlerapi/utils"
	"github.com/golang-jwt/jwt"
	"time"
)

type AuthService struct {
	repoLayer *repositories.RepositoryLayer
	jwtHelper *utils.JwtTokenHelper
}

func NewAuthService(
	repoLayer *repositories.RepositoryLayer,
	jwtHelper *utils.JwtTokenHelper,
) *AuthService {
	return &AuthService{
		repoLayer: repoLayer,
		jwtHelper: jwtHelper,
	}
}

func (s *AuthService) SaveRefreshTokenForUser(userID uint64, refreshToken *utils.JwtToken) error {
	return s.
		repoLayer.
		AuthTokenStorageRepository.
		StoreToken(userID, refreshToken.Token, refreshToken.Expires)
}

func (s *AuthService) GenerateAccessToken(userID uint64) (accessToken *utils.JwtToken, err error) {
	return s.jwtHelper.CreateAccessTokenForUser(userID)
}

func (s *AuthService) GenerateRefreshToken(userID uint64) (refreshToken *utils.JwtToken, err error) {
	return s.jwtHelper.CreateRefreshTokenForUser(userID)
}

func (s *AuthService) GenerateAuthTokenPair(userID uint64) (accessToken, refreshToken *utils.JwtToken, err error) {

	accessToken, err = s.GenerateAccessToken(userID)

	if err != nil {
		return
	}

	refreshToken, err = s.GenerateRefreshToken(userID)

	return
}

func (s *AuthService) ValidateAccessToken(accessToken string) (*jwt.StandardClaims, error) {

	claims, err := s.jwtHelper.ValidateToken(accessToken)

	if err != nil {
		if errors.As(err, jwt.ValidationError) {

		}
	}

	return claims, nil
}

func (s *AuthService) ValidateRefreshToken(userID uint64, refreshToken string) error {

	storedToken, err := s.repoLayer.AuthTokenStorageRepository.GetToken(userID)

	if err != nil {
		return err
	}

	if storedToken.Token != refreshToken {
		return apierr.AuthTokenInvalid
	}

	if storedToken.ExpiredAt.Before(time.Now()) {
		return apierr.AuthTokenExpired
	}

	return nil
}

func (s *AuthService) DeleteUserRefreshToken(userID uint64) error {
	return s.repoLayer.AuthTokenStorageRepository.DeleteToken(userID)
}
