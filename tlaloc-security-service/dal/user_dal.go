package dal

import (
	"errors"

	"gorm.io/gorm"
	"tlaloc-security-service/models"
)

type UserDal struct {
	DB *gorm.DB
}

func NewUserDal(db *gorm.DB) *UserDal {
	return &UserDal{DB: db}
}

func (d *UserDal) Create(user *models.User) error {
	if err := d.DB.Create(user).Error; err != nil {
		return errors.New("failed to create user")
	}
	return nil
}

func (d *UserDal) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := d.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (d *UserDal) GetUserByEmail(email string) (*models.User, error) {
	return d.FindByEmail(email)
}

func (d *UserDal) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := d.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (d *UserDal) GetUserByID(id uint) (*models.User, error) {
	return d.FindByID(id)
}

func (d *UserDal) Update(user *models.User) error {
	if err := d.DB.Save(user).Error; err != nil {
		return errors.New("failed to update user")
	}
	return nil
}

func (d *UserDal) Delete(id uint) error {
	if err := d.DB.Delete(&models.User{}, id).Error; err != nil {
		return errors.New("failed to delete user")
	}
	return nil
}
