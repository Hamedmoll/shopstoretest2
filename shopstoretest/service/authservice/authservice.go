package authservice

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"shopstoretest/cfg"
	"shopstoretest/entity"
	"strings"
	"time"
)

type Service struct {
	signKey               string
	accessExpirationTime  time.Duration
	refreshExpirationTime time.Duration
	accessSubject         string
	refreshSubject        string
}

func New(cfg cfg.Cfg) Service {
	authCfg := cfg.AuthCfg
	newService := Service{
		signKey:               authCfg.SignKey,
		accessExpirationTime:  authCfg.AccessExpirationTime,
		refreshExpirationTime: authCfg.RefreshExpirationTime,
		accessSubject:         authCfg.AccessSubject,
		refreshSubject:        authCfg.RefreshSubject,
	}

	return newService
}

func (s Service) CreateToken(id uint, expireDuration time.Duration, subject string, role entity.Role) (string, error) {
	claims := &Claim{
		ID:           id,
		TokenSubject: subject,
		Role:         role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.signKey))
	if err != nil {

		return "", fmt.Errorf("cant signed token %v", err)
	}

	return tokenString, nil
}

func (s Service) CreateAccessToken(id uint, role entity.Role) (string, error) {
	tokenString, err := s.CreateToken(id, s.accessExpirationTime, s.accessSubject, role)
	if err != nil {

		return "", err
	}

	return tokenString, nil
}

func (s Service) CreateRefreshToken(id uint, role entity.Role) (string, error) {
	tokenString, err := s.CreateToken(id, s.refreshExpirationTime, s.refreshSubject, role)
	if err != nil {

		return "", err
	}

	return tokenString, nil
}

func (s Service) ParseToken(BearerTokenString string) (Claim, error) {
	tokenString := strings.Replace(BearerTokenString, "Bearer ", "", 1)
	claims := Claim{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.signKey), nil
	})

	if err != nil {

		return Claim{}, fmt.Errorf("cant parse token %v", err)
	}

	return claims, nil
}
