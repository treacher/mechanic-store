package controllers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/treacher/mechanic-store/helpers"
	"github.com/treacher/mechanic-store/models"

	"encoding/json"
	"net/http"
)

func CreateCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var company models.Company

	err := helpers.ParseJsonIntoObject(r.Body, &company)

	helpers.HandleRequestError(w, err, http.StatusBadRequest)

	err = company.PersistCompany()

	helpers.HandleRequestError(w, err, http.StatusUnprocessableEntity)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(company)
}
