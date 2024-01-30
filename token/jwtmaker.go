package token

import (
	"fmt"
	"time"
)

const minSecretKeySize = 32

type JwtMaker struct {
	secretKey string
}

func NewJwtMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JwtMaker{secretKey}, nil
}

func (maker *JwtMaker) CreateToken(username string, duration time.Duration) (string, error) {
	return "", nil
}

func (maker *JwtMaker) VerifyToken(token string) (*Payload, error) {
	return &Payload{}, nil
}
