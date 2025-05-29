package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/maolchen/project_demo/config"
	"github.com/maolchen/project_demo/constants"
	"time"
)

var SecretKey = config.Cfg.Secret

type JWT struct {
	SigningKey []byte
}

type PublicClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 创建一个jwt实例
func NewJWT() *JWT {
	return &JWT{[]byte(GetSecretKey())}
}

// 获取secretkey
func GetSecretKey() []byte {
	return []byte(SecretKey)
}

// 生成token
func (j *JWT) AccessToken(username string) (string, error) {
	now := time.Now()
	expirationTime := now.Add(time.Duration(config.Cfg.JwtExpires) * time.Minute)

	claims := PublicClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "chen",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)

}

// 解析token
func (j *JWT) ParseToken(tokenString string) (*PublicClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &PublicClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		switch {
		case !token.Valid:
			return nil, errors.New(constants.TokenInvalid)
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, errors.New(constants.TokenMalformed)
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			// Invalid signature
			return nil, errors.New(constants.TokenSignatureInvalid)
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			// Token is either expired or not active yet
			return nil, errors.New(constants.TokenExpired)
		default:
			return nil, errors.New(constants.TokenParasError)

		}
	}

	if claims, ok := token.Claims.(*PublicClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New(constants.TokenInvalid)
}
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	claims, err := j.ParseToken(tokenString)
	if err != nil {
		return "", err
	}

	// 如果剩余时间小于5分钟，则刷新
	if claims.ExpiresAt.Time.Sub(time.Now()) < 5*time.Minute {
		now := time.Now()
		newExpirationTime := now.Add(time.Duration(config.Cfg.JwtExpires) * time.Minute)

		claims.IssuedAt = jwt.NewNumericDate(now)
		claims.ExpiresAt = jwt.NewNumericDate(newExpirationTime)

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString(j.SigningKey)
	}

	// 否则返回原 token
	return tokenString, nil
}
