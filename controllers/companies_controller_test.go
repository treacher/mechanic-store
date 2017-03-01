package controllers_test

import (
  "github.com/treacher/mechanic-store/router"
  "github.com/treacher/mechanic-store/models"
  "github.com/treacher/mechanic-store/db"
  "gopkg.in/pg.v5"

  "net/http"
  "net/http/httptest"
  "testing"
  "strings"
  "encoding/json"
)

func TestCreateCompany(t *testing.T) {
  db.Connection = pg.Connect(&pg.Options{ User: "postgres", Database: "mechanic-store" })

  companyJson := `{"name": "dennis", "phone": "+64505050505", "email": "foo@bar.com" }`
  companyReader := strings.NewReader(companyJson)

  req, err := http.NewRequest("POST", "/companies", companyReader)

  handleFatalError(t, err)

  rr := setupRecorder(req)

  testResponseCode(t, rr.Code, http.StatusCreated)

  var company models.Company

  json.NewDecoder(rr.Body).Decode(&company)

  testCompanyPersisted(t, company.Id)
  destroyCreatedCompany(&company)
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

func testCompanyPersisted(t *testing.T, companyId uint64) {
   var company models.Company

   _, err := db.Connection.QueryOne(&company, `SELECT * FROM companies WHERE id = ?`, companyId)

  if err != nil {
    t.Errorf("expected company with id: %v to exist but it did not %v", companyId, err.Error())
  }
}

func destroyCreatedCompany(company *models.Company) {
  db.Connection.Delete(company)
}
