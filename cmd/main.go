package main

import (
	"coffee-server/internal/router"
	"log"
)

func main() {
	//setup router
	r := router.Router()
	//start listen
	log.Fatal(r.Listen(":3000"))
}
