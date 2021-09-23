package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

type JwtTokenHelper struct {
	key             string
	signingMethod   jwt.SigningMethod
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

type JwtTokenType string

const (
	AccessToken  JwtTokenType = "AToken"
	RefreshToken JwtTokenType = "RToken"
)

type JwtToken struct {
	TokenType JwtTokenType
	Token     string
	TTL       time.Duration
	Expires   time.Time
}

func NewJwtTokenHelper(key string, signingMethod jwt.SigningMethod, accessTokenTTL, refreshTokenTTL time.Duration) *JwtTokenHelper {
	return &JwtTokenHelper{
		key:             key,
		signingMethod:   signingMethod,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

func (jh *JwtTokenHelper) CreateAccessTokenForUser(userID uint64) (*JwtToken, error) {
	expires := time.Now().Add(jh.accessTokenTTL)

	token, err := jwt.
		NewWithClaims(jh.signingMethod, &jwt.StandardClaims{
			Subject:   StringifyUint64(userID),
			ExpiresAt: expires.Unix(),
		}).
		SignedString([]byte(jh.key))

	if err != nil {
		return nil, err
	}

	return &JwtToken{
		TokenType: AccessToken,
		Token:     token,
		TTL:       jh.accessTokenTTL,
		Expires:   expires,
	}, nil
}

func (jh *JwtTokenHelper) CreateRefreshTokenForUser(userID uint64) (*JwtToken, error) {

	sha := sha512.New()
	sha.Write([]byte(jh.key + StringifyUint64(userID) + strconv.FormatInt(time.Now().Unix(), 36)))
	token := hex.EncodeToString(sha.Sum(nil))

	return &JwtToken{
		TokenType: RefreshToken,
		Token:     token,
		TTL:       jh.refreshTokenTTL,
		Expires:   time.Now().Add(jh.refreshTokenTTL),
	}, nil
}

func (jh *JwtTokenHelper) ValidateToken(accessToken string) (*jwt.StandardClaims, error) {
	claims := jwt.StandardClaims{}

	_, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jh.signingMethod.Alg() {
			return nil, fmt.Errorf("invalid signing algorythm: %s", token.Header["alg"])
		}
		return []byte(jh.key), nil
	})

	if err != nil {
		return &claims, err
	}

	return &claims, nil
}
