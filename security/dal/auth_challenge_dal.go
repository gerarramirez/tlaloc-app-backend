package dal

import (
	"time"

	"gorm.io/gorm"
)

type AuthChallengeDal struct {
	db *gorm.DB
}

func NewAuthChallengeDal(db *gorm.DB) *AuthChallengeDal {
	return &AuthChallengeDal{db: db}
}

// TableName especifica el schema y tabla para AuthChallenge
func (AuthChallenge) TableName() string {
	return "tlaloc_security_user.auth_challenges"
}

// AuthChallenge model for database operations
type AuthChallenge struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"not null;index" json:"email"`
	Challenge string    `gorm:"not null;size:64" json:"challenge"`
	Nonce     string    `gorm:"not null;size:32" json:"nonce"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	IsUsed    bool      `gorm:"not null;default:false" json:"is_used"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateChallenge - Crea un nuevo challenge en la base de datos
func (d *AuthChallengeDal) CreateChallenge(challenge *AuthChallenge) error {
	err := d.db.Where("email = ?", challenge.Email).Delete(&AuthChallenge{}).Error
	if err != nil {
		return err
	}
	return d.db.Create(challenge).Error
}

// GetActiveChallengeByEmail - Obtiene el challenge activo para un email
func (d *AuthChallengeDal) GetActiveChallengeByEmail(email string) (*AuthChallenge, error) {
	var challenge AuthChallenge
	err := d.db.Where("email = ? AND is_used = ? AND expires_at > ?", email, false, time.Now()).First(&challenge).Error
	if err != nil {
		return nil, err
	}
	return &challenge, nil
}

// MarkChallengeAsUsed - Marca un challenge como usado
func (d *AuthChallengeDal) MarkChallengeAsUsed(challengeID uint) error {
	return d.db.Model(&AuthChallenge{}).Where("id = ?", challengeID).Update("is_used", true).Error
}

// CleanupExpiredChallenges - Limpia challenges expirados
func (d *AuthChallengeDal) CleanupExpiredChallenges() error {
	return d.db.Where("expires_at < ? OR is_used = ?", time.Now(), true).Delete(&AuthChallenge{}).Error
}
