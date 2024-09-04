package entity

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name        string    `gorm:"type:varchar(100)" json:"name"`
	PhoneNumber string    `gorm:"type:varchar(20);unique" json:"phoneNumber"`
	Address     string    `gorm:"type:text" json:"address"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
