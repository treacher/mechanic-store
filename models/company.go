package models

import (
  "github.com/treacher/mechanic-store/database"

  "time"
)

type Company struct {
  Id         uint64   `json:"id"`
  Name       string   `json:"name"`
  Phone      string   `json:"phone"`
  Email      string   `json:"email"`
  CreatedAt time.Time `json:"createdAt"`
  UpdatedAt time.Time `json:"updatedAt"`
}

func (company Company) InitDates() Company {
  var currentTime = time.Now().UTC()

  companyWithDates := &Company{
    Name: company.Name,
    Phone: company.Phone,
    Email: company.Email,
    CreatedAt: currentTime,
    UpdatedAt: currentTime,
  }

  return *companyWithDates;
}

func (company *Company) PersistCompany() error {
  err := database.Connection.Insert(company)
  return err
}
