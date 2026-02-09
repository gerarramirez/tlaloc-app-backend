package main

import (
	"fmt"
	"log"

	"tlaloc-security-service/config"
	"tlaloc-security-service/dal"
	"tlaloc-security-service/handler"
	"tlaloc-security-service/jwt"
	mw "tlaloc-security-service/middleware"
	"tlaloc-security-service/models"
	"tlaloc-security-service/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	if err := db.AutoMigrate(&models.User{}, &models.RefreshToken{}, &models.UserSession{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	userDal := dal.NewUserDal(db)
	refreshTokenDal := dal.NewRefreshTokenDal(db)

	jwtService := jwt.NewJWTService(cfg.JWTSecret, cfg.JWTRefresh)

	// Usar solo el handler seguro
	authChallengeDal := dal.NewAuthChallengeDal(db)
	secureAuthHandler := handler.NewSecureAuthHandler(userDal, refreshTokenDal, authChallengeDal, jwtService)
	authMw := mw.NewAuthMiddleware(jwtService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"*"},
	}))

	routes.RegisterSecureAuthRoutes(e, secureAuthHandler, authMw)

	port := ":" + cfg.ServerPort
	if port == ":" {
		port = ":8081"
	}

	e.Logger.Fatal(e.Start(port))
}
