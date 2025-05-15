package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type JWTData struct {
	Email string
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(data JWTData) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
	})
	return t.SignedString([]byte(j.Secret))
}

func (j *JWT) Parse(token string) (bool, *JWTData, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil, err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return false, nil, errors.New("invalid token claims")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return false, nil, errors.New("invalid email in token")
	}

	return true, &JWTData{
		Email: email,
	}, nil
}
