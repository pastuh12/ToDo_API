package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/todo_api/config"
)

const (
	TTL = 15 * time.Minute
)

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresAt    int64  `json:"tokenTTL"`
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int
}

func NewToken(userId int) (*Token, error) {
	var (
		err      error
		newToken Token = Token{
			AccessToken:  "",
			RefreshToken: "",
			ExpiresAt:    time.Now().Add(TTL).Unix(),
		}
	)

	err = newToken.CreateAccessToken(userId)
	if err != nil {
		return nil, err
	}

	err = newToken.CreateRefreshToken()
	if err != nil {
		return &newToken, err
	}

	return &newToken, nil
}

func (t *Token) CreateAccessToken(userId int) error {
	var err error
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&tokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(TTL).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			userId,
		},
	)

	t.AccessToken, err = token.SignedString([]byte(config.Get().SigningKey))
	if err != nil {
		t.AccessToken = ""
		return err
	}

	return nil
}

func (t *Token) CreateRefreshToken() error {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	_, err := r.Read(b)
	if err != nil {
		t.RefreshToken = ""
		return err
	}

	t.RefreshToken = fmt.Sprintf("%x", b)

	return nil
}
