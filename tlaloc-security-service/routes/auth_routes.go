package routes

import (
	"tlaloc-security-service/handler"
	"tlaloc-security-service/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterAuthRoutes(e *echo.Echo, authHandler *handler.AuthHandler, authMiddleware *middleware.AuthMiddleware) {
	e.POST("/auth/login", authHandler.Login)
	e.POST("/auth/refresh", authHandler.RefreshToken)
	e.POST("/auth/validate", authHandler.ValidateToken)
	e.POST("/auth/logout", authHandler.Logout)

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

	admin := e.Group("/admin")
	admin.Use(authMiddleware.RequireAuth)
	admin.Use(authMiddleware.RequireRole("admin"))
	{
		admin.GET("/test", func(c echo.Context) error {
			return c.JSON(200, map[string]string{
				"message": "Admin access granted",
			})
		})
	}
}
