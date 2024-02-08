package config

import (
	"akgo/env"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
)

var (
	secretKey = []byte(env.PasswordPrivateKey)

	redisClient *redis.Client
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func CreateAccessToken(user User) (string, error) {
	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(), // 15 minutes
			IssuedAt:  time.Now().Unix(),
			Subject:   "access_token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func CreateRefreshToken(user User) (string, error) {
	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days
			IssuedAt:  time.Now().Unix(),
			Subject:   "refresh_token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ParseAccessToken(tokenString string) (*Claims, error) {
	tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid access token")
}

func ParseRefreshToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if claims.Subject == "refresh_token" {
			return claims, nil
		}
	}

	return nil, fmt.Errorf("invalid refresh token")
}

func RevokeAccessToken(tokenString string) error {
	claims, err := ParseAccessToken(tokenString)
	if err != nil {
		return err
	}

	ctx := context.Background()
	key := fmt.Sprintf("token_blacklist:%s", claims.Id)
	return redisClient.SAdd(ctx, key, tokenString).Err()
}

func IsAccessTokenRevoked(tokenString string) bool {
	claims, err := ParseAccessToken(tokenString)
	if err != nil {
		return false
	}

	ctx := context.Background()
	key := fmt.Sprintf("token_blacklist:%s", claims.Id)
	exists, err := redisClient.SIsMember(ctx, key, tokenString).Result()
	return err == nil && exists
}
