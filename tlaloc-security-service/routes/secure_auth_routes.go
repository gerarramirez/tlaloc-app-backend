package routes

import (
	"tlaloc-security-service/handler"
	"tlaloc-security-service/middleware"

	"github.com/labstack/echo/v4"
)

// RegisterSecureAuthRoutes - Registra endpoints de autenticación segura
func RegisterSecureAuthRoutes(e *echo.Echo, secureHandler *handler.SecureAuthHandler, authMiddleware *middleware.AuthMiddleware) {
	// Health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status":    "healthy",
			"service":   "tlaloc-security-service",
			"version":   "1.0.0",
			"auth_type": "challenge-response",
		})
	})

	// Endpoint público para obtener challenge
	e.POST("/auth/challenge", secureHandler.GetChallenge)

	// Endpoint de login seguro con challenge
	e.POST("/auth/login", secureHandler.SecureLogin)

	// Endpoint de refresh token
	e.POST("/auth/refresh", secureHandler.RefreshToken)

	// Protected routes
	protected := e.Group("/auth")
	protected.Use(authMiddleware.RequireAuth)
	{
		protected.GET("/me", func(c echo.Context) error {
			userID := c.Get("user_id").(uint)
			email := c.Get("email").(string)
			role := c.Get("role").(string)

			return c.JSON(200, map[string]interface{}{
				"user_id": userID,
				"email":   email,
				"role":    role,
			})
		})
	}

}
