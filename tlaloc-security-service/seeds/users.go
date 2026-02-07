package seeds

import (
	"tlaloc-security-service/models"
	"tlaloc-security-service/security"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	hashedPassword, err := security.HashPassword("admin123456")
	if err != nil {
		return err
	}

	adminUser := models.User{
		Email:     "admin@tlaloc.com",
		Password:  hashedPassword,
		FirstName: "Admin",
		LastName:  "User",
		Role:      "admin",
		IsActive:  true,
	}

	if err := db.Where("email = ?", adminUser.Email).FirstOrCreate(&adminUser).Error; err != nil {
		return err
	}

	userHashedPassword, err := security.HashPassword("user123456")
	if err != nil {
		return err
	}

	normalUser := models.User{
		Email:     "user@tlaloc.com",
		Password:  userHashedPassword,
		FirstName: "Normal",
		LastName:  "User",
		Role:      "user",
		IsActive:  true,
	}

	if err := db.Where("email = ?", normalUser.Email).FirstOrCreate(&normalUser).Error; err != nil {
		return err
	}

	return nil
}
