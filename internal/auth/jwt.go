package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWTAuthenticator struct {
	secrete string
	iss     string
	aud     string
}

func NewJWTAuthenticator(secrete, iss, aud string) *JWTAuthenticator {
	return &JWTAuthenticator{
		secrete,
		iss,
		aud,
	}

}

func (j *JWTAuthenticator) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.secrete))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWTAuthenticator) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(j.secrete), nil
	},
		jwt.WithIssuer(j.iss),
		jwt.WithAudience(j.aud),
		jwt.WithExpirationRequired(),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
}
