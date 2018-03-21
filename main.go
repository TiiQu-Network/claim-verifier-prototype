package main

import (
	"github.com/TiiQu-Network/claim-verifier-prototype/http/routes"
	"log"
	"net/http"
)

func main() {
	routes.Routes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
