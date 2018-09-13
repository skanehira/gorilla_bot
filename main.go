package main

import (
	"log"

	"github.com/skanehira/gorilla_bot/server"
)

func main() {
	// log setting
	log.SetFlags(log.Llongfile)

	// new server
	s := server.New()

	// start server
	switch s.Protocol {
	case "http":
		s.Start()
	case "https":
		s.StartTLS()
	}

}
