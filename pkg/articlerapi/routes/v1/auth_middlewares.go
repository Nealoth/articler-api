package v1apiroutes

import (
	"github.com/Nealoth/articler-api/pkg/articlerapi/apierr"
	"github.com/Nealoth/articler-api/pkg/articlerapi/logger"
	"github.com/Nealoth/articler-api/pkg/articlerapi/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *RoutesStore) jwtAuthMiddleware(c *gin.Context) {
	sl := r.serviceLayer

	accessToken, err := c.Request.Cookie(string(utils.AccessToken))

	if err != nil {
		switch err {
		case http.ErrNoCookie:
			respondError(c, http.StatusUnauthorized, err)
		default:
			respondError(c, http.StatusInternalServerError, err)
		}
		return
	}

	claims, err := sl.AuthService.ValidateAccessToken(accessToken.Value)

	if err != nil {
		switch err {
		case apierr.AuthTokenExpired:
			//TODO nil claims
			userID, _ := utils.ParseUint64(claims.Subject)
			refreshToken, _ := c.Request.Cookie(string(utils.RefreshToken))

			if rTokValErr := sl.AuthService.ValidateRefreshToken(userID, refreshToken.Value); rTokValErr != nil {
				respondError(c, http.StatusForbidden, err)
				logger.Debug("Refresh token is invalid")
				return
			}

			accessToken, _ := sl.AuthService.GenerateAccessToken(userID)
			logger.Debug("Access token has been successfully renewed")
			SetJwtCookie(c, accessToken, r.authPath)
		default:
			respondError(c, http.StatusForbidden, err)
			return
		}

	}

	c.Set(utils.UserIDSessionAttribute, claims.Subject)
}
