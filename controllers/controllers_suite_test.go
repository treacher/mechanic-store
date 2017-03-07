package controllers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/treacher/mechanic-store/db"
	"gopkg.in/pg.v5"

	"testing"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Controllers Suite")
}

var _ = BeforeSuite(func() {
	db.Connection = pg.Connect(&pg.Options{User: "postgres", Database: "mechanic-store"})
})

var _ = AfterSuite(func() {
	db.Connection.Exec("DELETE FROM companies")
})
