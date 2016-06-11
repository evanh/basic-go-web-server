package main

import (
	"log"

	"github.com/evanh/fundmyworld/server"
)

func main() {
	srvr := server.CreateAndInitializeServer()
	srvr.AddHandlers()

	log.Println("App starting...")
	log.Fatal(srvr.HTTPServer.ListenAndServe())
}
