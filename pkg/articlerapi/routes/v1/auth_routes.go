package v1apiroutes

import (
	"github.com/Nealoth/articler-api/pkg/articlerapi/domain"
	"github.com/Nealoth/articler-api/pkg/articlerapi/logger"
	"github.com/Nealoth/articler-api/pkg/articlerapi/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *RoutesStore) signUp(c *gin.Context) {
	var userCredentials domain.UserCredentials
	sl := r.serviceLayer

	if err := c.BindJSON(&userCredentials); err != nil {
		logger.Error(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userID, err := sl.UserService.CreateUser(userCredentials)

	if err != nil {
		logger.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = r.createTokenSessionForUser(c, userID)
	if err != nil {
		logger.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(200)
}

func (r *RoutesStore) signIn(c *gin.Context) {
	var userCredentials domain.UserCredentials
	sl := r.serviceLayer
	err := c.BindJSON(&userCredentials)

	if err != nil {
		logger.Error(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userProfile, err := sl.UserService.GetUserProfileByCredentials(userCredentials)

	if err != nil {
		//TODO declare global api errors
		logger.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = r.createTokenSessionForUser(c, userProfile.ID)
	if err != nil {
		logger.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(200)
}

func (r *RoutesStore) logout(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)

	if err != nil {
		respondError(c, http.StatusUnauthorized, err)
		return
	}

	DeleteJwtCookie(c, utils.AccessToken, r.authPath)
	DeleteJwtCookie(c, utils.RefreshToken, r.authPath)

	err = r.serviceLayer.AuthService.DeleteUserRefreshToken(userID)

	if err != nil {
		respondError(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(200)
}

func DeleteJwtCookie(c *gin.Context, tokenType utils.JwtTokenType, path string) {
	c.SetCookie(
		string(tokenType),
		"",
		-1,
		path,
		"",
		false,
		true,
	)
}

//TODO replace
func SetJwtCookie(c *gin.Context, token *utils.JwtToken, path string) {
	c.SetCookie(
		string(token.TokenType),
		token.Token,
		int(token.TTL.Seconds()),
		path,
		"",
		false,
		true,
	)
}

func (r *RoutesStore) createTokenSessionForUser(c *gin.Context, userID uint64) error {
	sl := r.serviceLayer
	accessToken, refreshToken, err := sl.AuthService.GenerateAuthTokenPair(userID)

	if err != nil {
		return err
	}

	err = sl.AuthService.SaveRefreshTokenForUser(userID, refreshToken)

	if err != nil {
		return err
	}

	SetJwtCookie(c, accessToken, r.authPath)
	SetJwtCookie(c, refreshToken, r.authPath)

	return nil
}
