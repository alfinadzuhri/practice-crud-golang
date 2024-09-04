package repository

import (
	"github.com/alfinadzuhri/practice-crud-golang/model/entity"

	"gorm.io/gorm"
)

type CustomerReposository interface {
	Save(customerReq entity.Customer) (entity.Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) CustomerReposository {
	return &customerRepository{
		db: db,
	}
}

func (c *customerRepository) Save(customerReq entity.Customer) (entity.Customer, error) {

	result := c.db.Create(&customerReq)

	if result.Error != nil {
		return entity.Customer{}, result.Error
	}

	return customerReq, nil

}
