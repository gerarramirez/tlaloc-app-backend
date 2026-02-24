package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"tlaloc-security-service/jwt"
)

type AuthMiddleware struct {
	jwtService *jwt.JWTService
}

func NewAuthMiddleware(jwtService *jwt.JWTService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}

func (m *AuthMiddleware) RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":     "unauthorized",
				"message":   "Authorization header is required",
				"timestamp": time.Now().Unix(),
			})
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error":     "invalid_token_format",
				"message":   "Token must be in 'Bearer <token>' format",
				"timestamp": time.Now().Unix(),
			})
		}

		tokenString := tokenParts[1]
		claims, err := m.jwtService.ValidateAccessToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":     "invalid_token",
				"message":   "Token is invalid or expired",
				"timestamp": time.Now().Unix(),
			})
		}

		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		return next(c)
	}
}

func (m *AuthMiddleware) RequireRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userRole, ok := c.Get("role").(string)
			if !ok {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"error":     "forbidden",
					"message":   "User role not found",
					"timestamp": time.Now().Unix(),
				})
			}

			if userRole != role && userRole != "admin" {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"error":         "forbidden",
					"message":       "Insufficient permissions",
					"required_role": role,
					"user_role":     userRole,
					"timestamp":     time.Now().Unix(),
				})
			}

			return next(c)
		}
	}
}
