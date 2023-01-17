package repository

import (
	"gorm.io/gorm"
	"useLaborartory-backend/models"
)

type DatabaseRepo interface {
	Connection() *gorm.DB
	Registration(requestPayload models.RegisterLabType) (bool, error)
}
