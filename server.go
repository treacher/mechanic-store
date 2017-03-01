package main

import (
  "github.com/treacher/mechanic-store/router"
  "github.com/treacher/mechanic-store/db"
  "gopkg.in/pg.v5"

  "net/http"
  "log"
)

func main() {
  db.Connection = pg.Connect(&pg.Options{ User: "postgres", Database: "mechanic-store" })

  log.Fatal(http.ListenAndServe(":8080", router.Router()))
}
