package controllers

import (
  "github.com/julienschmidt/httprouter"
  "github.com/treacher/mechanic-store/models"

  "net/http"
  "encoding/json"
  "math/rand"
  "time"
)

func CreateCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  var company models.Company

  if r.Body == nil {
    http.Error(w, "No body provided", 400)
    return
  }

  err := json.NewDecoder(r.Body).Decode(&company)

  if err != nil {
    http.Error(w, err.Error(), 400)
    return
  }

  currentTime := time.Now()

  company.Id = rand.Uint64()
  company.CreatedAt = currentTime
  company.UpdatedAt = currentTime

  w.WriteHeader(http.StatusCreated)

  json.NewEncoder(w).Encode(company)
}

