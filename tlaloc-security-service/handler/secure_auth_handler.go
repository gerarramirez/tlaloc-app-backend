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

type SecureAuthHandler struct {
	userDal          *dal.UserDal
	refreshTokenDal  *dal.RefreshTokenDal
	authChallengeDal *dal.AuthChallengeDal
	jwtService       *jwt.JWTService
}

func NewSecureAuthHandler(userDal *dal.UserDal, refreshTokenDal *dal.RefreshTokenDal, authChallengeDal *dal.AuthChallengeDal, jwtService *jwt.JWTService) *SecureAuthHandler {
	return &SecureAuthHandler{
		userDal:          userDal,
		refreshTokenDal:  refreshTokenDal,
		authChallengeDal: authChallengeDal,
		jwtService:       jwtService,
	}
}

// ChallengeRequest - Request para obtener challenge
type ChallengeRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// SecureLoginRequest - Request para login seguro
type SecureLoginRequest struct {
	Email        string `json:"email" validate:"required,email"`
	Challenge    string `json:"challenge" validate:"required"`
	Nonce        string `json:"nonce" validate:"required"`
	Timestamp    int64  `json:"timestamp" validate:"required"`
	ResponseHash string `json:"response_hash" validate:"required"`
}

// GetChallenge - Endpoint para obtener un challenge de login
func (h *SecureAuthHandler) GetChallenge(c echo.Context) error {
	var req ChallengeRequest
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

	// Verificar si el usuario existe y está activo
	user, err := h.userDal.GetUserByEmail(req.Email)
	if err != nil {
		// Error genérico por seguridad - no revelar si el usuario existe
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "invalid_request",
			Message: "Invalid email address",
			Code:    http.StatusBadRequest,
		})
	}

	if !user.IsActive {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "account_disabled",
			Message: "Account is disabled",
			Code:    http.StatusUnauthorized,
		})
	}

	// Generar challenge seguro
	challenge, err := security.GenerateChallenge()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "challenge_generation_error",
			Message: "Failed to generate challenge",
			Code:    http.StatusInternalServerError,
		})
	}

	// Almacenar challenge en base de datos
	authChallenge := &dal.AuthChallenge{
		Email:     req.Email,
		Challenge: challenge.Challenge,
		Nonce:     challenge.Nonce,
		ExpiresAt: challenge.ExpiresAt,
		IsUsed:    false,
	}

	if err := h.authChallengeDal.CreateChallenge(authChallenge); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "challenge_storage_error",
			Message: "Failed to store challenge",
			Code:    http.StatusInternalServerError,
		})
	}

	// Limpiar challenges expirados
	h.authChallengeDal.CleanupExpiredChallenges()

	// Devolver challenge response sin el ID interno
	response := security.ChallengeResponse{
		Challenge: challenge.Challenge,
		ExpiresAt: challenge.ExpiresAt,
		Timestamp: challenge.Timestamp,
		Nonce:     challenge.Nonce,
	}

	return c.JSON(http.StatusOK, response)
}

// SecureLogin - Login seguro con challenge-response
func (h *SecureAuthHandler) SecureLogin(c echo.Context) error {
	var req SecureLoginRequest
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

	// Obtener el challenge almacenado de la base de datos
	authChallenge, err := h.authChallengeDal.GetActiveChallengeByEmail(req.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "invalid_challenge",
			Message: "Challenge not found or expired",
			Code:    http.StatusUnauthorized,
		})
	}

	// Convertir a ChallengeResponse para validación
	storedChallenge := &security.ChallengeResponse{
		Challenge: authChallenge.Challenge,
		Nonce:     authChallenge.Nonce,
		ExpiresAt: authChallenge.ExpiresAt,
		Timestamp: authChallenge.CreatedAt,
	}

	// Verificar si el challenge es válido
	if !security.IsChallengeValid(storedChallenge) {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "challenge_expired",
			Message: "Challenge has expired",
			Code:    http.StatusUnauthorized,
		})
	}

	// Verificar que el challenge coincide
	if req.Challenge != storedChallenge.Challenge ||
		req.Nonce != storedChallenge.Nonce ||
		req.Timestamp != storedChallenge.Timestamp.Unix() {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "challenge_mismatch",
			Message: "Challenge parameters do not match",
			Code:    http.StatusUnauthorized,
		})
	}

	// Obtener el usuario
	user, err := h.userDal.GetUserByEmail(req.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "invalid_credentials",
			Message: "Invalid email or challenge response",
			Code:    http.StatusUnauthorized,
		})
	}

	// Verificar el hash de respuesta
	isValid, err := security.VerifyChallengeHash(user.Password, storedChallenge.Challenge, storedChallenge.Nonce, req.ResponseHash, storedChallenge.Timestamp.Unix())
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "invalid_credentials",
			Message: "Invalid credentials",
			Code:    http.StatusUnauthorized,
		})
	}

	if !isValid {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "invalid_credentials",
			Message: "Invalid credentials",
			Code:    http.StatusUnauthorized,
		})
	}

	// Marcar challenge como usado en la base de datos
	if err := h.authChallengeDal.MarkChallengeAsUsed(authChallenge.ID); err != nil {
		// Log error pero continuar con authentication
		// El challenge expirará automáticamente por tiempo
	}

	// Generar tokens
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

// RefreshToken - Refresca el access token usando un refresh token válido
func (h *SecureAuthHandler) RefreshToken(c echo.Context) error {
	// Obtener refresh token del request
	refreshToken := c.FormValue("refresh_token")
	if refreshToken == "" {
		// También intentar obtener del body JSON
		var req struct {
			RefreshToken string `json:"refresh_token" form:"refresh_token"`
		}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Error:   "invalid_request",
				Message: "Refresh token is required",
				Code:    http.StatusBadRequest,
			})
		}
		refreshToken = req.RefreshToken
	}

	if refreshToken == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "missing_refresh_token",
			Message: "Refresh token is required",
			Code:    http.StatusBadRequest,
		})
	}

	// Validar refresh token en la base de datos
	tokenRecord, err := h.refreshTokenDal.GetRefreshToken(refreshToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "invalid_refresh_token",
			Message: "Refresh token is invalid or expired",
			Code:    http.StatusUnauthorized,
		})
	}

	// Verificar que el refresh token esté activo y no expirado
	if !tokenRecord.IsActive || tokenRecord.ExpiresAt.Before(time.Now()) {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "expired_refresh_token",
			Message: "Refresh token has expired",
			Code:    http.StatusUnauthorized,
		})
	}

	// Obtener información del usuario
	user, err := h.userDal.GetUserByID(tokenRecord.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "user_not_found",
			Message: "User not found",
			Code:    http.StatusInternalServerError,
		})
	}

	if !user.IsActive {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:   "account_disabled",
			Message: "Account is disabled",
			Code:    http.StatusUnauthorized,
		})
	}

	// Generar nuevo access token
	newAccessToken, _, err := h.jwtService.GenerateAccessToken(user.ID, user.Email, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "token_generation_error",
			Message: "Failed to generate new access token",
			Code:    http.StatusInternalServerError,
		})
	}

	// Generar nuevo refresh token
	newRefreshToken, err := h.jwtService.GenerateRefreshToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "refresh_token_generation_error",
			Message: "Failed to generate new refresh token",
			Code:    http.StatusInternalServerError,
		})
	}

	// Crear nuevo registro de refresh token
	newTokenRecord := &models.RefreshToken{
		UserID:    user.ID,
		Token:     newRefreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour), // 7 días
		IsActive:  true,
	}

	if err := h.refreshTokenDal.CreateRefreshToken(newTokenRecord); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "refresh_token_storage_error",
			Message: "Failed to store new refresh token",
			Code:    http.StatusInternalServerError,
		})
	}

	// Invalidar refresh token antiguo (revocarlo)
	if err := h.refreshTokenDal.DeactivateRefreshToken(refreshToken); err != nil {
		// Log error pero continuar (no es crítico si falla la revocación)
		// En producción, podríamos loggear esto para auditoría
	}

	// Retornar nuevos tokens
	response := map[string]interface{}{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
		"expires_in":    15 * 60, // 15 minutos en segundos
		"token_type":    "Bearer",
	}

	return c.JSON(http.StatusOK, response)
}
