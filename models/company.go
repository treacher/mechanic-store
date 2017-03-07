package models

import (
	"github.com/treacher/mechanic-store/db"

	"time"
)

type Company struct {
	Id        uint64    `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (company *Company) PersistCompany() error {
	var currentTime = time.Now().UTC()

	company.CreatedAt = currentTime
	company.UpdatedAt = currentTime

	err := db.Connection.Insert(&company)
	return err
}
