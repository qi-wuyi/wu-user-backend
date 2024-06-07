package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"wu-user-backend/utils/pkg/constants"
)

type Claims struct {
	UserId   int64  `json:"user_id"`
	UserRole string `json:"user_role"`
	jwt.RegisteredClaims
}

func CreatToken(userId int64, userRole string) (string, error) {
	expireTime := time.Now().Add(24 * time.Hour)
	nowTime := time.Now()
	claims := Claims{
		UserId:   userId,
		UserRole: userRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			Issuer:    "wu-pao", //颁发者签名
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.JWTValue))
}

func ValidateToken(token string) (*Claims, error) {
	response, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.JWTValue), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := response.Claims.(*Claims); ok && response.Valid {
		return claims, nil
	}
	return nil, err
}
