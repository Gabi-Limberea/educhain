package handler

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AuthToken struct {
	TTL            time.Duration `json:"ttl"`
	PID            int64         `json:"pid"`
	IssuedAt       time.Time     `json:"issuedAt"`
	ExpirationTime time.Time     `json:"expirationTime"`
	Issuer         string        `json:"issuer"`
}

func (t *AuthToken) ToJWT() *jwt.Token {
	claims := jwt.MapClaims{}
	claims["iss"] = t.Issuer
	claims["iat"] = t.IssuedAt.Unix()
	claims["exp"] = t.ExpirationTime.Unix()
	claims["pid"] = t.PID
	claims["ttl"] = t.TTL

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

func FromJWT(t *jwt.Token) (*AuthToken, error) {
	auth := &AuthToken{}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	if iss, ok := claims["iss"]; !ok {
		return nil, fmt.Errorf("missing token issuer")
	} else {
		auth.Issuer = iss.(string)
	}

	if pid, ok := claims["pid"]; !ok {
		return nil, fmt.Errorf("missing token pid")
	} else {
		auth.PID = int64(pid.(float64))
	}

	if ttl, ok := claims["ttl"]; !ok {
		return nil, fmt.Errorf("missing token ttl")
	} else {
		auth.TTL = time.Duration(ttl.(float64))
	}

	if iat, ok := claims["iat"]; !ok {
		return nil, fmt.Errorf("missing token issuedAt")
	} else {
		auth.IssuedAt = time.Unix(int64(iat.(float64)), 0)
	}

	if exp, ok := claims["exp"]; !ok {
		return nil, fmt.Errorf("missing token expirationTime")
	} else {
		auth.ExpirationTime = time.Unix(int64(exp.(float64)), 0)
	}

	return auth, nil
}
