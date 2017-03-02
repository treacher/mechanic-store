package controllers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/treacher/mechanic-store/helpers"
	"github.com/treacher/mechanic-store/models"

	"encoding/json"
	"net/http"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var customer models.Customer

	err := helpers.ParseJsonIntoObject(r.Body, &customer)

	helpers.HandleRequestError(w, err, http.StatusBadRequest)

	err = customer.PersistCompany()

	helpers.HandleRequestError(w, err, http.StatusUnprocessableEntity)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(customer)
}
