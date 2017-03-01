package controllers

import (
  "github.com/julienschmidt/httprouter"
  "github.com/treacher/mechanic-store/models"

  "net/http"
  "encoding/json"
)

func CreateCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  var company models.Company

  if r.Body == nil {
    http.Error(w, "No body provided", http.StatusBadRequest)
    return
  }

  err := json.NewDecoder(r.Body).Decode(&company)

  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  company = company.InitDates()

  err = company.PersistCompany()

  if err != nil {
    http.Error(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }

  w.WriteHeader(http.StatusCreated)

  json.NewEncoder(w).Encode(company)
}

