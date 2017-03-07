package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/treacher/mechanic-store/controllers"

	"net/http"
)

func Router() http.Handler {
	router := httprouter.New()

	router.POST("/companies", controllers.CreateCompany)
	router.POST("/customers", controllers.CreateCustomer)

	return router
}
