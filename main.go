package main

import (
	"gorilla_bot/server"
	"log"
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
