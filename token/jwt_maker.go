package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
	jwt.RegisteredClaims
}

var ErrInvalidKeySize error = fmt.Errorf("invalid key size, must be at least %d characters", minSecretKeySize)

func NewJWTMaker(secretKey string, duration time.Duration) (*JWTMaker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, ErrInvalidKeySize
	}
	return &JWTMaker{secretKey, jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
	}}, nil
}

func (j *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, j.RegisteredClaims)
	return jwtToken.SignedString(j.secretKey)
}

func (j *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyfunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(j.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, keyfunc)
	if err != nil {
		return nil, err
	}

	payload, ok := jwtToken.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	fmt.Printf("payload %v", payload)

	return nil, nil
}
