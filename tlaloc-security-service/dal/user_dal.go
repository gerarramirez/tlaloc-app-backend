package dal

import (
	"tlaloc-security-service/models"

	"gorm.io/gorm"
)

type UserDal struct {
	db *gorm.DB
}

func NewUserDal(db *gorm.DB) *UserDal {
	return &UserDal{db: db}
}

func (u *UserDal) CreateUser(user *models.User) error {
	return u.db.Create(user).Error
}

func (u *UserDal) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDal) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := u.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDal) UpdateUser(user *models.User) error {
	return u.db.Save(user).Error
}
