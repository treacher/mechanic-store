package models

import (
	"github.com/treacher/mechanic-store/db"

	"time"
)

type Customer struct {
	Id        uint64    `json:"id"`
	CompanyId uint64    `json:"company_id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (customer *Customer) Persist() error {
	var currentTime = time.Now().UTC()

	customer.CreatedAt = currentTime
	customer.UpdatedAt = currentTime

	err := db.Connection.Insert(&customer)
	return err
}
