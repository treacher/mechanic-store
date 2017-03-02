package controllers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/treacher/mechanic-store/models"

	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func CreateCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var company models.Company

	err := ParseJsonIntoObject(r.Body, &company)

	HandleRequestError(w, err, http.StatusBadRequest)

	err = company.PersistCompany()

	HandleRequestError(w, err, http.StatusUnprocessableEntity)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(company)
}

func ParseJsonIntoObject(body io.Reader, obj interface{}) error {
	if body == nil {
		return errors.New("No body provided")
	}

	err := json.NewDecoder(body).Decode(obj)

	return err
}

func HandleRequestError(w http.ResponseWriter, err error, code int) {
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}
}
