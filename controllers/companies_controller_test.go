package controllers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/treacher/mechanic-store/db"
	"github.com/treacher/mechanic-store/models"
	"github.com/treacher/mechanic-store/router"

	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = Describe("Companies Controller", func() {
	var (
		companyJson      string
		responseRecorder *httptest.ResponseRecorder
		company          models.Company
	)

	Describe("Creating a company", func() {
		JustBeforeEach(func() {
			companyReader := strings.NewReader(companyJson)
			req, err := http.NewRequest("POST", "/companies", companyReader)

			Expect(err).NotTo(HaveOccurred())

			responseRecorder = httptest.NewRecorder()

			router := router.Router()

			router.ServeHTTP(responseRecorder, req)
		})

		Context("Valid JSON", func() {
			BeforeEach(func() {
				companyJson = `{"name": "dennis", "phone": "+64505050505", "email": "foo@bar.com" }`
			})

			It("Persists the company", func() {
				_, err := db.Connection.QueryOne(&company, `SELECT * FROM companies WHERE name = ?`, "dennis")

				Expect(err).NotTo(HaveOccurred())
				Expect(company.Phone).To(Equal("+64505050505"))
				Expect(company.Email).To(Equal("foo@bar.com"))
				Expect(company.CreatedAt).ToNot(BeNil())
				Expect(company.UpdatedAt).ToNot(BeNil())
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
