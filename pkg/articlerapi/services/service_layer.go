package services

import (
	"github.com/Nealoth/articler-api/pkg/articlerapi/configuration"
	"github.com/Nealoth/articler-api/pkg/articlerapi/domain"
	"github.com/Nealoth/articler-api/pkg/articlerapi/repositories"
	"github.com/Nealoth/articler-api/pkg/articlerapi/utils"
	"github.com/golang-jwt/jwt"
)

type IUserService interface {
	CreateUser(credentials domain.UserCredentials) (uint64, error)
	GetUserProfileByID(userID uint64) (*domain.UserProfile, error)
	GetUserProfileByCredentials(credentials domain.UserCredentials) (*domain.UserProfile, error)
}

type IAuthService interface {
	GenerateAccessToken(userID uint64) (accessToken *utils.JwtToken, err error)
	GenerateRefreshToken(userID uint64) (refreshToken *utils.JwtToken, err error)
	GenerateAuthTokenPair(userID uint64) (accessToken, refreshToken *utils.JwtToken, err error)
	ValidateAccessToken(accessToken string) (*jwt.StandardClaims, error)
	ValidateRefreshToken(userID uint64, refreshToken string) error
	SaveRefreshTokenForUser(userID uint64, refreshToken *utils.JwtToken) error
	DeleteUserRefreshToken(userID uint64) error
}

type ServiceLayer struct {
	UserService IUserService
	AuthService IAuthService
}

func InitServiceLayer(
	conf configuration.ServerConfiguration,
	authConf configuration.JwtConfiguration,
	repoLayer *repositories.RepositoryLayer,
) *ServiceLayer {
	return &ServiceLayer{

		UserService: NewUserService(
			repoLayer,
			utils.NewHasher(conf.UserPasswordSalt),
		),

		AuthService: NewAuthService(
			repoLayer,
			utils.NewJwtTokenHelper(
				authConf.SigningKey,
				jwt.SigningMethodHS256,
				utils.SafeParseDuration(authConf.AccessTokenTTL),
				utils.SafeParseDuration(authConf.RefreshTokenTTL),
			),
		),
	}
}
