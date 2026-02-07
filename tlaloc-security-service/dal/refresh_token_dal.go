package dal

import (
	"tlaloc-security-service/models"

	"gorm.io/gorm"
)

type RefreshTokenDal struct {
	db *gorm.DB
}

func NewRefreshTokenDal(db *gorm.DB) *RefreshTokenDal {
	return &RefreshTokenDal{db: db}
}

func (r *RefreshTokenDal) CreateRefreshToken(token *models.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *RefreshTokenDal) GetRefreshToken(token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	err := r.db.Preload("User").Where("token = ? AND is_active = ?", token, true).First(&refreshToken).Error
	if err != nil {
		return nil, err
	}
	return &refreshToken, nil
}

func (r *RefreshTokenDal) DeactivateRefreshToken(token string) error {
	return r.db.Model(&models.RefreshToken{}).Where("token = ?", token).Update("is_active", false).Error
}

func (r *RefreshTokenDal) DeactivateUserRefreshTokens(userID uint) error {
	return r.db.Model(&models.RefreshToken{}).Where("user_id = ?", userID).Update("is_active", false).Error
}

func (r *RefreshTokenDal) DeleteExpiredTokens() error {
	return r.db.Where("expires_at < NOW()").Delete(&models.RefreshToken{}).Error
}
