package dal

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"tlaloc-security-service/models"
)

type RefreshTokenDal struct {
	DB *gorm.DB
}

func NewRefreshTokenDal(db *gorm.DB) *RefreshTokenDal {
	return &RefreshTokenDal{DB: db}
}

func (d *RefreshTokenDal) Create(token *models.RefreshToken) error {
	if err := d.DB.Create(token).Error; err != nil {
		return errors.New("failed to create refresh token")
	}
	return nil
}

func (d *RefreshTokenDal) CreateRefreshToken(token *models.RefreshToken) error {
	return d.Create(token)
}

func (d *RefreshTokenDal) FindByToken(tokenString string) (*models.RefreshToken, error) {
	var token models.RefreshToken
	if err := d.DB.Where("token = ? AND revoked = ?", tokenString, false).First(&token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("refresh token not found or revoked")
		}
		return nil, err
	}
	return &token, nil
}

func (d *RefreshTokenDal) GetRefreshToken(tokenString string) (*models.RefreshToken, error) {
	return d.FindByToken(tokenString)
}

func (d *RefreshTokenDal) FindByUserID(userID uint) ([]models.RefreshToken, error) {
	var tokens []models.RefreshToken
	if err := d.DB.Where("user_id = ? AND revoked = ?", userID, false).Find(&tokens).Error; err != nil {
		return nil, err
	}
	return tokens, nil
}

func (d *RefreshTokenDal) Revoke(tokenString string) error {
	result := d.DB.Model(&models.RefreshToken{}).Where("token = ?", tokenString).Update("revoked", true)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("token not found")
	}
	return nil
}

func (d *RefreshTokenDal) RevokeAllUserTokens(userID uint) error {
	result := d.DB.Model(&models.RefreshToken{}).Where("user_id = ?", userID).Update("revoked", true)
	return result.Error
}

func (d *RefreshTokenDal) DeactivateRefreshToken(tokenString string) error {
	return d.Revoke(tokenString)
}

func (d *RefreshTokenDal) DeleteExpired() error {
	result := d.DB.Where("expires_at < ?", time.Now()).Delete(&models.RefreshToken{})
	return result.Error
}
