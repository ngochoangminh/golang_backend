package util

import (
	"fmt"
	"time"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)
type Payload struct {
    UserID    uuid.UUID `json:"user_id"`
    IssuedAt  time.Time `json:"issued_at"`
    ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(userId uuid.UUID, duration time.Duration) (*Payload, error) {
    payload := &Payload{
        UserID:        userId,
        IssuedAt:  time.Now(),
        ExpiredAt: time.Now().Add(duration),
    }
    return payload, nil
}


func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func JwtEncode(userId uuid.UUID, key string) (string, error) {
	payload, err := NewPayload(userId, time.Hour)
	if err != nil {
		fmt.Println(err)
        return "", err
    }

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, payload)

	token, err := jwtToken.SignedString([]byte(key))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return token, nil
}

func JwtDecode(token string, key string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(key), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}