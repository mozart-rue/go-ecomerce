package main

import (
	"log"

	"github.com/mozart-rue/go-ecomerce/cmd/api"
)

func main() {
  server := api.NewApiServer(":8080", nil)
  if err := server.Run(); err != nil {
    log.Fatal(err)
  }
}
