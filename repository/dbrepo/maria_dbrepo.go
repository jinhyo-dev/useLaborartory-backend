package dbrepo

import (
	"errors"
	"gorm.io/gorm"
	"useLaborartory-backend/models"
)

type MariaDBRepo struct {
	DB *gorm.DB
}

func (m *MariaDBRepo) Connection() *gorm.DB {
	return m.DB
}

func (m *MariaDBRepo) Registration(requestPayload models.RegisterLabType) (bool, error) {
	m.DB.Table("registration").AutoMigrate(&models.RegisterLabType{})
	success := m.DB.Table("registration").Create(&requestPayload)

	if success != nil {
		return true, nil
	}

	return false, errors.New("registration is not complete")
}
