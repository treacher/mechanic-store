package main

import (
  "github.com/treacher/mechanic-store/router"
  "net/http"
  "log"
)

func main() {
  log.Fatal(http.ListenAndServe(":8080", router.Router()))
}
