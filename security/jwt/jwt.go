package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secretKey     string
	refreshKey    string
	accessExpiry  time.Duration
	refreshExpiry time.Duration
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func NewJWTService(secretKey, refreshKey string) *JWTService {
	return &JWTService{
		secretKey:     secretKey,
		refreshKey:    refreshKey,
		accessExpiry:  15 * time.Minute,
		refreshExpiry: 7 * 24 * time.Hour,
	}
}

func (s *JWTService) GenerateAccessToken(userID uint, email, role string) (string, string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.accessExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "tlaloc-security",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", "", err
	}

	// Also return the expiry time as a string
	expiry := time.Now().Add(s.accessExpiry).Format(time.RFC3339)
	return accessToken, expiry, nil
}

func (s *JWTService) GenerateRefreshToken(userID uint) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.refreshExpiry)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "tlaloc-security",
		Subject:   string(rune(userID)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.refreshKey))
}

func (s *JWTService) ValidateAccessToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func (s *JWTService) ValidateRefreshToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.refreshKey), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.RegisteredClaims); ok && token.Valid {
		userID := claims.Subject
		var userIDUint uint
		_, err := parseUserID(userID, &userIDUint)
		if err != nil {
			return 0, err
		}
		return userIDUint, nil
	}

	return 0, errors.New("invalid token")
}

func parseUserID(s string, v *uint) (bool, error) {
	var u uint
	for _, c := range s {
		if c < '0' || c > '9' {
			return false, errors.New("invalid user id")
		}
		u = u*10 + uint(c-'0')
	}
	*v = u
	return true, nil
}
