package models

import "time"

type Company struct {
  Id         uint64   `json:"id"`
  Name       string   `json:"name"`
  Phone      string   `json:"phone"`
  Email      string   `json:"email"`
  CreatedAt time.Time `json:"createdAt"`
  UpdatedAt time.Time `json:"updatedAt"`
}
