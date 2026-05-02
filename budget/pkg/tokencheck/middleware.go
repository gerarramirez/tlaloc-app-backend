package tokencheck

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// RequireAuth - Middleware simplificado con dependencia de tokencheck
func RequireAuth(jwtSecret string) echo.MiddlewareFunc {
	config := NewJWTConfig(jwtSecret)
	return ValidateToken(config)
}

// RequireAuthWithConfig - Middleware con configuración personalizada
func RequireAuthWithConfig(config *JWTConfig) echo.MiddlewareFunc {
	return ValidateToken(config)
}

// RequireRole - Middleware que requiere rol específico
func RequireRole(requiredRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role, ok := c.Get("role").(string)
			if !ok {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"error":     "forbidden",
					"message":   "User role not found",
					"timestamp": time.Now().Unix(),
					"code":      http.StatusForbidden,
				})
			}

			// Admin puede acceder a todo
			if role == "admin" || role == requiredRole {
				return next(c)
			}

			// Si no es admin, verificar rol específico
			if role != requiredRole {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"error":         "forbidden",
					"message":       "Insufficient permissions",
					"required_role": requiredRole,
					"user_role":     role,
					"timestamp":     time.Now().Unix(),
					"code":          http.StatusForbidden,
				})
			}

			return next(c)
		}
	}
}

// RequireAdmin - Middleware que requiere rol de administrador
func RequireAdmin() echo.MiddlewareFunc {
	return RequireRole("admin")
}

// GetUserID - Obtiene ID del usuario del contexto
func GetUserID(c echo.Context) (uint, error) {
	if userID, ok := c.Get("user_id").(uint); ok {
		return userID, nil
	}
	return 0, echo.NewHTTPError(http.StatusUnauthorized, "user_id not found in context")
}

// GetUserEmail - Obtiene email del usuario del contexto
func GetUserEmail(c echo.Context) (string, error) {
	if email, ok := c.Get("email").(string); ok {
		return email, nil
	}
	return "", echo.NewHTTPError(http.StatusUnauthorized, "email not found in context")
}

// GetUserRole - Obtiene rol del usuario del contexto
func GetUserRole(c echo.Context) (string, error) {
	if role, ok := c.Get("role").(string); ok {
		return role, nil
	}
	return "", echo.NewHTTPError(http.StatusUnauthorized, "role not found in context")
}

// IsAuthenticated - Verifica si el usuario está autenticado
func IsAuthenticated(c echo.Context) bool {
	_, ok := c.Get("user_id").(uint)
	return ok
}

// GetTokenInfo - Obtiene información completa del token
func GetTokenInfo(c echo.Context) (map[string]interface{}, error) {
	if metadata, ok := c.Get("token_metadata").(map[string]interface{}); ok {
		return metadata, nil
	}
	return map[string]interface{}{}, echo.NewHTTPError(http.StatusUnauthorized, "token metadata not found")
}

// CORS Headers - Helper para headers de respuesta CORS
func SetCORSHeaders(c echo.Context) {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
}

// respondSuccess - Helper para respuestas de éxito
func respondSuccess(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success":   true,
		"data":      data,
		"timestamp": time.Now().Unix(),
	})
}

// respondError - Helper para respuestas de error
func respondError(c echo.Context, status int, message string, details interface{}) error {
	errorResponse := map[string]interface{}{
		"success":   false,
		"error":     message,
		"timestamp": time.Now().Unix(),
		"code":      status,
	}

	if details != nil {
		errorResponse["details"] = details
	}

	return c.JSON(status, errorResponse)
}

// HealthCheck - Endpoint de health para verificar que la librería funciona
func HealthCheck(c echo.Context) error {
	return respondSuccess(c, map[string]interface{}{
		"service":     "tlaloc-tokencheck-lib",
		"version":     "1.0.0",
		"description": "Token validation library for Tlaloc microservices",
		"status":      "healthy",
		"timestamp":   time.Now().Unix(),
	})
}
