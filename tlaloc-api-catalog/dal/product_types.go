package dal

import (
	"errors"
	"gorm.io/gorm"
	model "tlaloc-catalog/model/db" // Mantener el alias 'model' está bien
)

type ProductTypeDAO interface {
	Create(productType *model.ProductType) error // Usar singular en el parámetro
	GetAll() ([]model.ProductType, error)        // Usar GetAll o All
	Update(productType *model.ProductType) error // Usar singular en el parámetro
}

type ProductTypeDAOImpl struct {
	DB *gorm.DB
}

func NewProductTypesDAO(db *gorm.DB) ProductTypeDAO {
	return &ProductTypeDAOImpl{
		DB: db,
	}
}

func (dao *ProductTypeDAOImpl) Create(productType *model.ProductType) error {

	if productType == nil {
		return errors.New("Datos insuficientes")
	}

	db := dao.DB.Begin()

	if err := db.Table("tlaloc_api.product_types").Select("Name").Create(productType).Error; err != nil {
		db.Rollback()
		return errors.New("")
	}

	return db.Commit().Error
}

func (dao *ProductTypeDAOImpl) GetAll() ([]model.ProductType, error) {
	var productTypes []model.ProductType
	db := dao.DB.Begin()
	result := db.Table("tlaloc_api.product_types").Find(&productTypes)

	return productTypes, result.Error
}

func (dao *ProductTypeDAOImpl) Update(productType *model.ProductType) error {
	db := dao.DB.Begin()
	if err := db.Table("tlaloc_api.product_types").Save(productType).Error; err != nil {
		db.Rollback()
		return errors.New("Error en el server")
	}

	return db.Commit().Error
}
