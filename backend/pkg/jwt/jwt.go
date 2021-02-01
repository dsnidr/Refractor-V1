package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sniddunc/refractor/refractor"
	"time"
)

// Claims represents the auth claims made in our JWTs
type Claims struct {
	UserID      int64 `json:"id"`
	AccessLevel int   `json:"accessLevel"`
	jwt.StandardClaims
}

// RefreshClaims represents the claims made in our refresh JWTs
type RefreshClaims struct {
	UserID int64 `json:"id"`
	jwt.StandardClaims
}

type TokenPair struct {
	AuthToken    string
	RefreshToken string
}

// GenerateAuthTokens generates an authorization and a refresh token
func GenerateAuthTokens(user *refractor.User, secret []byte) (string, string, error) {
	authClaims := Claims{
		UserID:      user.UserID,
		AccessLevel: user.AccessLevel,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
	}

	// Create auth token
	authToken := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims)

	// Create auth token string
	authTokenString, err := authToken.SignedString(secret)
	if err != nil {
		return "", "", err
	}

	// Create refresh token
	refreshClaims := RefreshClaims{
		UserID: user.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(14 * (24 * time.Hour)).Unix(),
		},
	}

	// Create refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	// Create refresh token string
	refreshTokenString, err := refreshToken.SignedString(secret)
	if err != nil {
		return "", "", err
	}

	return authTokenString, refreshTokenString, nil
}

func ExtractRefreshClaims(refreshToken string, jwtSecret string) (*RefreshClaims, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*RefreshClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func ExtractAuthClaims(authToken string, jwtSecret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(authToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
