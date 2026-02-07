package middleware

import (
	"net/http"
	"strings"

	"tlaloc-security-service/dto"
	"tlaloc-security-service/jwt"

	"github.com/labstack/echo/v4"
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
			return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Error:   "missing_token",
				Message: "Authorization header is required",
				Code:    http.StatusUnauthorized,
			})
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Error:   "invalid_token_format",
				Message: "Token must be in 'Bearer <token>' format",
				Code:    http.StatusUnauthorized,
			})
		}

		token := tokenParts[1]
		claims, err := m.jwtService.ValidateAccessToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Error:   "invalid_token",
				Message: "Token is invalid or expired",
				Code:    http.StatusUnauthorized,
			})
		}

		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		return next(c)
	}
}

func (m *AuthMiddleware) RequireRole(requiredRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userRole, ok := c.Get("role").(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
					Error:   "unauthorized",
					Message: "User role not found",
					Code:    http.StatusUnauthorized,
				})
			}

			if userRole != requiredRole && userRole != "admin" {
				return c.JSON(http.StatusForbidden, dto.ErrorResponse{
					Error:   "forbidden",
					Message: "Insufficient permissions",
					Code:    http.StatusForbidden,
				})
			}

			return next(c)
		}
	}
}
