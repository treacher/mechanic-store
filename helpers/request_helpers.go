package helpers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

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
