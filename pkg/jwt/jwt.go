package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"rest_go_kv/pkg/logger"
)

type JWTData struct {
	Email string
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	logger.Debug("JWT instance created")
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(data JWTData) (string, error) {
	if j == nil {
		logger.Error("JWT instance is nil")
		return "", errors.New("jwt instance is nil")
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
	})

	token, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		logger.Error("Failed to create token: %v\n", err)
		return "", err
	}

	logger.Info("Token created successfully for email: %s\n", data.Email)
	return token, nil

}

func (j *JWT) Parse(token string) (bool, *JWTData, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.Error("Unexpected signing method")
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.Secret), nil
	})
	if err != nil {
		logger.Error("Failed to parse token: %v\n", err)
		return false, nil, err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		logger.Error("Invalid token claims")
		return false, nil, errors.New("invalid token claims")
	}

	email, ok := claims["email"].(string)
	if !ok {
		logger.Error("Invalid email in token")
		return false, nil, errors.New("invalid email in token")
	}

	logger.Info("Token parsed successfully for email: %s\n", email)
	return true, &JWTData{
		Email: email,
	}, nil
}
