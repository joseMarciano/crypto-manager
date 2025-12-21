package repository

import (
	"time"

	"github.com/joseMarciano/crypto-manager/internal/app/user/domain"
)

type userEntity struct {
	ID             string    `gorm:"column:id;primaryKey"`
	Name           string    `gorm:"column:name"`
	Birthday       time.Time `gorm:"column:birthday"`
	DocumentNumber string    `gorm:"column:document_number"`
}

func (u userEntity) toDomain() domain.User {
	return domain.User(u)
}

func (u userEntity) TableName() string {
	return "users"
}

func toEntity(u domain.User) userEntity {
	return userEntity(u)
}
