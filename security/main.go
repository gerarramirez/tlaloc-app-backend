package main

import (
	"fmt"
	"log"
	"os"

	"tlaloc-security-service/dal"
	"tlaloc-security-service/handler"
	"tlaloc-security-service/jwt"
	mw "tlaloc-security-service/middleware"
	"tlaloc-security-service/routes"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Println("No se encontró config.env, usando variables de entorno del sistema")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	db.Exec("SET search_path TO 'tlaloc_security_user'")

	userDal := dal.NewUserDal(db)
	refreshTokenDal := dal.NewRefreshTokenDal(db)

	jwtService := jwt.NewJWTService(os.Getenv("JWT_SECRET"), os.Getenv("JWT_REFRESH_SECRET"))

	// Usar solo el handler seguro
	authChallengeDal := dal.NewAuthChallengeDal(db)
	secureAuthHandler := handler.NewSecureAuthHandler(userDal, refreshTokenDal, authChallengeDal, jwtService)
	authMw := mw.NewAuthMiddleware(jwtService)

	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"*"},
	}))

	routes.RegisterSecureAuthRoutes(e, secureAuthHandler, authMw)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8081"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
