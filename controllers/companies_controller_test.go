package controllers_test

import (
  "github.com/treacher/mechanic-store/router"
  "gopkg.in/DATA-DOG/go-sqlmock.v1"

  "net/http"
  "net/http/httptest"
  "testing"
  "strings"
)

func TestCreateCompany(t *testing.T) {
  db, mock, err := sqlmock.New()
  if err != nil {
      t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
  }
  defer db.Close()

  mock.ExpectBegin()
  mock.ExpectExec("INSERT INTO companies").WithArgs("dennis", "+64505050505", "foo@bar.com")
  mock.ExpectCommit()

  companyJson := `{"name": "dennis", "phone": "+64505050505", "email": "foo@bar.com" }`
  companyReader := strings.NewReader(companyJson)

  req, err := http.NewRequest("POST", "/companies", companyReader)

  handleFatalError(t, err)

  rr := setupRecorder(req)

  testResponseCode(t, rr.Code, http.StatusCreated)

  if err := mock.ExpectationsWereMet(); err != nil {
    t.Errorf("there were unfulfilled expections: %s", err)
  }
}

func TestCreateCompanyWithEmptyBody(t *testing.T) {
  req, err := http.NewRequest("POST", "/companies", nil)

  handleFatalError(t, err)

  rr := setupRecorder(req)

  testResponseCode(t, rr.Code, http.StatusBadRequest)
}

func TestCreateCompanyWithInvalidJSON(t *testing.T) {
  stringReader := strings.NewReader("foo")
  req, err := http.NewRequest("POST", "/companies", stringReader)

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
