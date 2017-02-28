package controllers_test

import (
  "github.com/treacher/mechanic-store/router"

  "net/http"
  "net/http/httptest"
  "testing"
  "strings"
)

func TestCreateCompany(t *testing.T) {
  companyJson := `{"name": "dennis", "phone": "+64505050505", "email": "foo@bar.com" }`
  companyReader := strings.NewReader(companyJson)

  req, err := http.NewRequest("POST", "/companies", companyReader)

  handleFatalError(t, err)

  rr := setupRecorder(req)

  testResponseCode(t, rr.Code, http.StatusCreated)
}

func TestCreateCompanyWithEmptyBody(t *testing.T) {
  req, err := http.NewRequest("POST", "/companies", nil)

  handleFatalError(t, err)

  rr := setupRecorder(req)

  testResponseCode(t, rr.Code, http.StatusBadRequest)
}

func setupRecorder(request *http.Request) *httptest.ResponseRecorder {
  rr := httptest.NewRecorder()

  router := router.Router()

  router.ServeHTTP(rr, request)

  return rr
}

func handleFatalError(t *testing.T, err error) {
  if err != nil {
      t.Fatal(err)
  }
}

func testResponseCode(t *testing.T, actual int, expected int) {
  if actual != expected {
    t.Errorf("handler returned wrong status code: got %v want %v", actual, expected)
  }
}
