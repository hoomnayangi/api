package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/phuwn/tools/errors"
	"github.com/phuwn/tools/util"
)

// TokenInfo - data model contains user's auth info
type TokenInfo struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
	Role   int `json:"role"`
}

var (
	secretKey           = util.Getenv("SERVER_SECRET_KEY", "")
	ErrInvalidToken     = errors.New("token expired, please log out and log in again", 401)
	ErrInvalidSignature = errors.New("invalid signature", 401)
	ErrBadToken         = errors.New("bad token", 401)
)

// GenerateJWTToken - create an access_token that represents user's session
func GenerateJWTToken(info *TokenInfo, expiresAt int64) (string, error) {
	info.ExpiresAt = expiresAt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, info)
	encryptedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", errors.Customize(err, 500, "failed to sign on token")
	}
	return encryptedToken, nil
}

// VerifyUserSession - validates user's access_token and returns user's id if it's verified
func VerifyUserSession(tokenString string) (*TokenInfo, error) {
	claims := TokenInfo{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if !token.Valid {
		return nil, ErrInvalidToken
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, ErrInvalidSignature
		}
		return nil, ErrBadToken
	}
	if time.Unix(claims.ExpiresAt, 0).Before(time.Now()) {
		return nil, ErrInvalidToken
	}
	return &claims, nil
}
