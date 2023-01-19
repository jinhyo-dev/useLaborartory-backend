package dbrepo

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"useLaborartory-backend/models"
)

type MariaDBRepo struct {
	DB *gorm.DB
}

func (m *MariaDBRepo) Connection() *gorm.DB {
	m.DB.Table("registration").AutoMigrate(&models.RegisterLabType{})
	return m.DB
}

func (m *MariaDBRepo) Registration(requestPayload models.RegisterLabType) (bool, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(requestPayload.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return false, err
	}
	values := models.RegisterLabType{
		LabNumber: requestPayload.LabNumber,
		Applicant: requestPayload.Applicant,
		StudentId: requestPayload.StudentId,
		Password:  string(password),
		Reason:    requestPayload.Reason,
	}
	success := m.DB.Table("registration").Create(&values)

	if success != nil {
		return true, nil
	}

	return false, errors.New("registration is not complete")
}
