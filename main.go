package main

import (
	"log"
	"os"

	"github.com/TiiQu-Network/claim-verifier-prototype/actions"
	"github.com/icrowley/fake"
)

func main() {

	for i := 0; i < 10; i++ {
		println(fake.FirstName() +" "+ fake.LastName())
	}
	os.Exit(1)

	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
