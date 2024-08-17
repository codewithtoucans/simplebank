package token

import (
	"time"

	"github.com/o1egl/paseto/v2"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (*PasetoMaker, error) {
	if len(symmetricKey) < chacha20poly1305.KeySize {
		return nil, ErrInvalidKeySize
	}

	return &PasetoMaker{
		paseto:       &paseto.V2{},
		symmetricKey: []byte(symmetricKey),
	}, nil
}

func (p *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayLoad(username, duration)
	if err != nil {
		return "", nil
	}
	return p.paseto.Encrypt(p.symmetricKey, payload, nil)
}

func (p *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := p.paseto.Decrypt(token, p.symmetricKey, payload, nil)
	if err != nil {
		return nil, err
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
