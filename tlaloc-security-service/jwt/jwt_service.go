package jwt

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type JWTService struct {
	secretKey        []byte
	refreshSecretKey []byte
}

func NewJWTService(secretKey, refreshSecretKey string) *JWTService {
	return &JWTService{
		secretKey:        []byte(secretKey),
		refreshSecretKey: []byte(refreshSecretKey),
	}
}

func (j *JWTService) GenerateAccessToken(userID uint, email, role string) (string, *jwt.MapClaims, error) {
	claims := JWTClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "tlaloc-security-service",
			Subject:   fmt.Sprintf("%d", userID),
			ID:        generateJTI(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", nil, err
	}

	mapClaims := jwt.MapClaims{
		"user_id": claims.UserID,
		"email":   claims.Email,
		"role":    claims.Role,
		"jti":     claims.ID,
		"exp":     claims.ExpiresAt.Unix(),
		"iat":     claims.IssuedAt.Unix(),
		"iss":     claims.Issuer,
		"sub":     claims.Subject,
	}

	return tokenString, &mapClaims, nil
}

func (j *JWTService) GenerateRefreshToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"type":    "refresh",
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
		"iss":     "tlaloc-security-service",
		"jti":     generateJTI(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.refreshSecretKey)
}

func (j *JWTService) ValidateAccessToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return j.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func (j *JWTService) ValidateRefreshToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return j.refreshSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if tokenType, exists := claims["type"]; !exists || tokenType != "refresh" {
			return nil, errors.New("not a refresh token")
		}
		return claims, nil
	}

	return nil, errors.New("invalid refresh token")
}

func generateJTI() string {
	bytes := make([]byte, 24)
	rand.Read(bytes)
	return base64.URLEncoding.EncodeToString(bytes)
}
