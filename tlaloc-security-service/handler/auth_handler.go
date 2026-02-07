package handler

import (
	"net/http"
	"time"

	"tlaloc-security-service/dal"
	"tlaloc-security-service/dto"
	"tlaloc-security-service/jwt"
	"tlaloc-security-service/models"
	"tlaloc-security-service/security"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	userDal         *dal.UserDal
	refreshTokenDal *dal.RefreshTokenDal
	jwtService      *jwt.JWTService
}

func NewAuthHandler(userDal *dal.UserDal, refreshTokenDal *dal.RefreshTokenDal, jwtService *jwt.JWTService) *AuthHandler {
	return &AuthHandler{
		userDal:         userDal,
		refreshTokenDal: refreshTokenDal,
		jwtService:      jwtService,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req dto.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "invalid_request",
			Message: "Invalid request body",
			Code:    http.StatusBadRequest,
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	user, err := h.userDal.GetUserByEmail(req.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "invalid_credentials",
			Message: "Invalid email or password",
			Code:    http.StatusUnauthorized,
		})
	}

	if !user.IsActive {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "account_disabled",
			Message: "Account is disabled",
			Code:    http.StatusUnauthorized,
		})
	}

	isValid, err := security.VerifyPassword(req.Password, user.Password)
	if err != nil || !isValid {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "invalid_credentials",
			Message: "Invalid email or password",
			Code:    http.StatusUnauthorized,
		})
	}

	accessToken, _, err := h.jwtService.GenerateAccessToken(user.ID, user.Email, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "token_generation_error",
			Message: "Failed to generate access token",
			Code:    http.StatusInternalServerError,
		})
	}

	refreshTokenStr, err := h.jwtService.GenerateRefreshToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "refresh_token_generation_error",
			Message: "Failed to generate refresh token",
			Code:    http.StatusInternalServerError,
		})
	}

	refreshToken := &models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshTokenStr,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
		IsActive:  true,
	}

	if err := h.refreshTokenDal.CreateRefreshToken(refreshToken); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "refresh_token_storage_error",
			Message: "Failed to store refresh token",
			Code:    http.StatusInternalServerError,
		})
	}

	response := dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenStr,
		ExpiresIn:    15 * 60,
		TokenType:    "Bearer",
		User: dto.UserInfo{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      user.Role,
			IsActive:  user.IsActive,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}

	return c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) RefreshToken(c echo.Context) error {
	var req dto.RefreshTokenRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "invalid_request",
			Message: "Invalid request body",
			Code:    http.StatusBadRequest,
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	refreshToken, err := h.refreshTokenDal.GetRefreshToken(req.RefreshToken)
	if err != nil || refreshToken == nil || !refreshToken.IsActive || time.Now().After(refreshToken.ExpiresAt) {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "invalid_refresh_token",
			Message: "Refresh token is invalid or expired",
			Code:    http.StatusUnauthorized,
		})
	}

	claims, err := h.jwtService.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "invalid_refresh_token",
			Message: "Refresh token is invalid",
			Code:    http.StatusUnauthorized,
		})
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "invalid_token_claims",
			Message: "Invalid token claims",
			Code:    http.StatusUnauthorized,
		})
	}

	userID := uint(userIDFloat)
	user, err := h.userDal.GetUserByID(userID)
	if err != nil || !user.IsActive {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "user_not_found",
			Message: "User not found or inactive",
			Code:    http.StatusUnauthorized,
		})
	}

	if err := h.refreshTokenDal.DeactivateRefreshToken(req.RefreshToken); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "refresh_token_deactivation_error",
			Message: "Failed to deactivate refresh token",
			Code:    http.StatusInternalServerError,
		})
	}

	newAccessToken, _, err := h.jwtService.GenerateAccessToken(user.ID, user.Email, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "token_generation_error",
			Message: "Failed to generate access token",
			Code:    http.StatusInternalServerError,
		})
	}

	newRefreshTokenStr, err := h.jwtService.GenerateRefreshToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "refresh_token_generation_error",
			Message: "Failed to generate refresh token",
			Code:    http.StatusInternalServerError,
		})
	}

	newRefreshToken := &models.RefreshToken{
		UserID:    user.ID,
		Token:     newRefreshTokenStr,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
		IsActive:  true,
	}

	if err := h.refreshTokenDal.CreateRefreshToken(newRefreshToken); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "refresh_token_storage_error",
			Message: "Failed to store refresh token",
			Code:    http.StatusInternalServerError,
		})
	}

	response := dto.RefreshTokenResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshTokenStr,
		ExpiresIn:    15 * 60,
		TokenType:    "Bearer",
	}

	return c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) ValidateToken(c echo.Context) error {
	var req dto.ValidateTokenRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "invalid_request",
			Message: "Invalid request body",
			Code:    http.StatusBadRequest,
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	claims, err := h.jwtService.ValidateAccessToken(req.Token)
	if err != nil {
		return c.JSON(http.StatusOK, dto.ValidateTokenResponse{
			UserID:    0,
			Email:     "",
			Role:      "",
			Valid:     false,
			ExpiresAt: 0,
		})
	}

	response := dto.ValidateTokenResponse{
		UserID:    claims.UserID,
		Email:     claims.Email,
		Role:      claims.Role,
		Valid:     true,
		ExpiresAt: claims.ExpiresAt.Unix(),
	}

	return c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) Logout(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" || len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "missing_token",
			Message: "Authorization header is missing or malformed",
			Code:    http.StatusBadRequest,
		})
	}

	token := authHeader[7:]
	claims, err := h.jwtService.ValidateAccessToken(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "invalid_token",
			Message: "Token is invalid",
			Code:    http.StatusUnauthorized,
		})
	}

	if err := h.refreshTokenDal.DeactivateUserRefreshTokens(claims.UserID); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "logout_error",
			Message: "Failed to deactivate user tokens",
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Successfully logged out",
	})
}
