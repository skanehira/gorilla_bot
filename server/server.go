package server

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/skanehira/gorilla_bot/common"
	"github.com/skanehira/gorilla_bot/config"
)

// Server https server
type Server struct {
	http.Server
	config.Config
	net.Listener
}

// New new https server
func New() *Server {
	s := new(Server)

	// initialize server
	s.SetConfig()
	s.SetHandler()

	return s
}

// SetHandler add handler to server
func (s *Server) SetHandler() {
	http.HandleFunc(s.Endpoint, s.Handler)
}

// SetConfig load config
func (s *Server) SetConfig() {
	s.Config = config.New("config/config.yaml")
}

// SetTLSListener new ssl listener
func (s *Server) SetTLSListener() {
	tlsConfig := new(tls.Config)
	tlsConfig.Certificates = make([]tls.Certificate, 1)

	var err error
	tlsConfig.Certificates[0], err = tls.LoadX509KeyPair(s.SSLCertificateFile, s.SSLCertificateKeyFile)
	if err != nil {
		panic(err)
	}

	tlsConfig.BuildNameToCertificate()

	listener, err := tls.Listen("tcp", s.Port, tlsConfig)
	if err != nil {
		panic(err)
	}

	s.Listener = listener
}

// StartTLS start https server
func (s *Server) StartTLS() {
	s.SetTLSListener()

	// server start
	fmt.Printf("[%s] start https server in %s\n", common.TimeNow(), s.Port)
	log.Fatal(s.Serve(s.Listener))
}

// Start start http server
func (s *Server) Start() {
	fmt.Printf("[%s] start http server in %s\n", common.TimeNow(), s.Port)
	log.Fatal(http.ListenAndServe(s.Port, nil))
}
