package tokencheck

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// Claims - Estructura de claims para tokens JWT
type Claims struct {
	UserID    uint   `json:"user_id"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	ExpiresAt int64  `json:"expires_at"`
	IssuedAt  int64  `json:"issued_at"`
	Issuer    string `json:"issuer"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

// TokenMetadata - Metadatos adicionales del token
type TokenMetadata struct {
	UserID    uint   `json:"user_id"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	ExpiresIn int64  `json:"expires_in"`
	IssuedAt  int64  `json:"issued_at"`
	TokenType string `json:"token_type"`
	Issuer    string `json:"issuer"`
}

// ErrorResponse - Estructura unificada para respuestas de error
type ErrorResponse struct {
	Error     string `json:"error"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
	Path      string `json:"path"`
	Code      int    `json:"code"`
}

// JWTConfig - Configuración para validación JWT
type JWTConfig struct {
	SecretKey       string `json:"-"`
	RequiredForTest bool   `json:"-"`
	SkipValidation  bool   `json:"-"`
	TestUserID      uint   `json:"-"`
	TestEmail       string `json:"-"`
	TestRole        string `json:"-"`
}

// NewJWTConfig - Crea configuración JWT por defecto
func NewJWTConfig(secretKey string) *JWTConfig {
	return &JWTConfig{
		SecretKey:       secretKey,
		RequiredForTest: false,
		SkipValidation:  false,
	}
}

// DecodeJWTWithoutValidation - Decodifica token sin validar firma/expiración
func DecodeJWTWithoutValidation(tokenString string) (*Claims, error) {
	parser := jwt.NewParser()
	token, _, err := parser.ParseUnverified(tokenString, &Claims{})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid claims")
}

// ValidateToken - Middleware para validación de tokens JWT
func ValidateToken(config *JWTConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Para testing: permitir tokens sin firma si está configurado
			if config.RequiredForTest {
				token := c.Get("test_token")
				if token != nil {
					testClaims := &Claims{
						UserID:    config.TestUserID,
						Email:     config.TestEmail,
						Role:      config.TestRole,
						ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
						IssuedAt:  time.Now().Unix(),
						Issuer:    "tlaloc-security-test",
						TokenType: "access",
					}
					c.Set("user_id", testClaims.UserID)
					c.Set("email", testClaims.Email)
					c.Set("role", testClaims.Role)
					c.Set("test_mode", true)
					return next(c)
				}
			}

			// Extraer token del header Authorization
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return respondError(c, http.StatusUnauthorized, "Authorization header is required", map[string]interface{}{
					"error_type": "missing_token",
				})
			}

			// Validar formato: "Bearer <token>"
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				return respondError(c, http.StatusBadRequest, "Token must be in 'Bearer <token>' format", map[string]interface{}{
					"error_type": "invalid_token_format",
				})
			}

			tokenString := tokenParts[1]
			if tokenString == "" {
				return respondError(c, http.StatusBadRequest, "Token is required in Authorization header", map[string]interface{}{
					"error_type": "missing_token",
				})
			}

			// Omitir validación si está configurado para testing
			var claims *Claims
			var err error

			if !config.SkipValidation {
				// Parsear y validar token
				token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}
					return []byte(config.SecretKey), nil
				})

				if err != nil {
					if strings.Contains(err.Error(), "token is malformed") {
						return respondError(c, http.StatusBadRequest, "Token is malformed", map[string]interface{}{
							"error_type": "malformed_token",
						})
					}
					if strings.Contains(err.Error(), "token is expired") {
						return respondError(c, http.StatusUnauthorized, "Token is invalid or expired", map[string]interface{}{
							"error_type": "invalid_token",
						})
					}
					return respondError(c, http.StatusUnauthorized, "Token is invalid", map[string]interface{}{
						"error_type": "invalid_token",
					})
				}

				if !token.Valid {
					return respondError(c, http.StatusUnauthorized, "Token is invalid", map[string]interface{}{
						"error_type": "invalid_token",
					})
				}

				claims = token.Claims.(*Claims)
			} else {
				// Solo decodificar sin validación para testing
				claims, err = DecodeJWTWithoutValidation(tokenString)
				if err != nil || claims == nil {
					return respondError(c, http.StatusUnauthorized, "Token is malformed", map[string]interface{}{
						"error_type": "malformed_token",
					})
				}
			}

			// Validar expiración
			if !config.SkipValidation && claims.ExpiresAt < time.Now().Unix() {
				return respondError(c, http.StatusUnauthorized, "Token has expired", map[string]interface{}{
					"error_type": "token_expired",
				})
			}

			// Validar issuer (opcional pero recomendado)
			if !config.SkipValidation && claims.Issuer != "tlaloc-security" {
				return respondError(c, http.StatusUnauthorized, "Token is not from a trusted issuer", map[string]interface{}{
					"error_type": "invalid_token",
				})
			}

			// Setear variables de contexto para el usuario
			c.Set("user_id", claims.UserID)
			c.Set("email", claims.Email)
			c.Set("role", claims.Role)
			c.Set("token_issued_at", claims.IssuedAt)
			c.Set("token_expires_at", claims.ExpiresAt)
			c.Set("token_issuer", claims.Issuer)

			// Crear metadata del token para respuestas
			tokenMetadata := TokenMetadata{
				UserID:    claims.UserID,
				Email:     claims.Email,
				Role:      claims.Role,
				ExpiresIn: claims.ExpiresAt - time.Now().Unix(),
				IssuedAt:  claims.IssuedAt,
				TokenType: claims.TokenType,
				Issuer:    claims.Issuer,
			}

			// Setear metadata como JSON para fácil acceso
			c.Set("token_metadata", tokenMetadata)
			c.Set("expires_in", tokenMetadata.ExpiresIn)

			// Log de depuración (solo en desarrollo)
			if config.RequiredForTest {
				c.Set("debug_mode", true)
			}

			return next(c)
		}
	}
}

// GetTokenMetadata - Obtiene metadatos del token actual
func GetTokenMetadata(c echo.Context) (*TokenMetadata, error) {
	if metadata, ok := c.Get("token_metadata").(*TokenMetadata); ok {
		return metadata, nil
	}
	return nil, echo.NewHTTPError(http.StatusUnauthorized, "no_token_metadata", "Token metadata not found")
}

// IsTokenExpired - Verifica si el token está expirado
func IsTokenExpired(c echo.Context) bool {
	if expiresAt, ok := c.Get("token_expires_at").(int64); ok {
		return time.Now().Unix() >= expiresAt
	}
	return true
}

// IsTestMode - Verifica si está en modo de testing
func IsTestMode(c echo.Context) bool {
	if testMode, ok := c.Get("test_mode").(bool); ok {
		return testMode
	}
	return false
}

// GenerateTestToken - Genera token de prueba para testing
func GenerateTestToken(config *JWTConfig) (string, error) {
	claims := &Claims{
		UserID:    config.TestUserID,
		Email:     config.TestEmail,
		Role:      config.TestRole,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "tlaloc-security-test",
		TokenType: "access",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SecretKey))
}
