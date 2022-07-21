package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

type TokenManager interface {
	// NewJWT - create a new jwt token, that expires quickly
	NewJWT(uuid.UUID) (string, error)
	// NewLongLive - create a new jwt token, that will live long time
	NewLongLive(uuid.UUID) (string, time.Time, error)
	// Verify token.
	// return id of subject, true if it isn't expired yet, error
	Verify(string) (uuid.UUID, bool, error)
}

type tokenManager struct {
	key         []byte
	jwtExp      time.Duration
	longLiveExp time.Duration
	issuerName  string
}

func NewTokenManager(key string, jwtExp, longLiveExp time.Duration, issuerName string) (TokenManager, error) {
	return &tokenManager{
		key:         []byte(key),
		jwtExp:      jwtExp,
		longLiveExp: longLiveExp,
		issuerName:  issuerName,
	}, nil
}

func (manager *tokenManager) NewJWT(id uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   id.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(manager.jwtExp)),
		Issuer:    manager.issuerName,
	})

	tokenStr, err := token.SignedString(manager.key)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (manager *tokenManager) NewLongLive(_ uuid.UUID) (string, time.Time, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", time.Time{}, err
	}

	return fmt.Sprintf("%s", b), time.Now().Add(manager.longLiveExp), nil
}

func (manager *tokenManager) Verify(input string) (uuid.UUID, bool, error) {
	token, err := jwt.Parse(input, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return manager.key, nil
		}
		return nil, errors.New("unexpected signing method")
	})

	if err != nil {
		return uuid.UUID{}, false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.UUID{}, false, errors.New("failed to get claims")
	}

	id, ok := claims["sub"].(string)
	if !ok {
		return uuid.UUID{}, false, errors.New("failed to get id from claims")
	}
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, false, errors.New("failed to parse id from claims")
	}

	return parsedId, claims.VerifyExpiresAt(time.Now().UnixNano(), true), nil
}
