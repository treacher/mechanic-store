package controllers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/treacher/mechanic-store/db"
	"github.com/treacher/mechanic-store/models"
	"github.com/treacher/mechanic-store/router"

	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = Describe("Customers Controller", func() {
	var (
		companyJson      string
		responseRecorder *httptest.ResponseRecorder
		customer         models.Customer
	)

	Describe("Creating a customer", func() {
		JustBeforeEach(func() {
			companyReader := strings.NewReader(companyJson)
			req, err := http.NewRequest("POST", "/customers", companyReader)

			Expect(err).NotTo(HaveOccurred())

			responseRecorder = httptest.NewRecorder()

			router := router.Router()
			router.ServeHTTP(responseRecorder, req)
		})

		Context("Valid JSON", func() {
			BeforeEach(func() {
				company := models.Company{Name: "Company", Phone: "+6454545454", Email: "bar@foo.com"}

				err := company.Persist()
				Expect(err).NotTo(HaveOccurred())

				companyJson = `{"name": "dennis", "phone": "+64505050505", "email": "foo@bar.com", "company_id": %v }`
				companyJson = fmt.Sprintf(companyJson, company.Id)
			})

			It("Persists the customer", func() {
				_, err := db.Connection.QueryOne(&customer, `SELECT * FROM customers WHERE name = ?`, "dennis")

				Expect(err).NotTo(HaveOccurred())
				Expect(customer.Phone).To(Equal("+64505050505"))
				Expect(customer.Email).To(Equal("foo@bar.com"))
				Expect(customer.CompanyId).ToNot(BeNil())
				Expect(customer.CreatedAt).ToNot(BeNil())
				Expect(customer.UpdatedAt).ToNot(BeNil())
			})

			It("responds with a StatusCreated response code", func() {
				Expect(responseRecorder.Code).To(Equal(http.StatusCreated))
			})
		})

		Context("Invalid JSON", func() {
			BeforeEach(func() {
				companyJson = "foo"
			})

			It("Responds with a BadRequest status", func() {
				Expect(responseRecorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})
