package config

import (
	"fmt"

	"github.com/jinzhu/configor"
)

// Config Server config
type Config struct {
	Protocol              string `default:"http"`
	Port                  string `default:":8080"`
	Endpoint              string `default:"/slack/gorilla"`
	SSLCertificateFile    string
	SSLCertificateKeyFile string
	URLVerifyToken        string
	AuthorizationToken    string
	MessageFile           string
}

// New new config
func New(file string) Config {
	conf := new(Config)

	if file == "" {
		panic("invalid config file path")
	}

	if err := configor.Load(conf, file); err != nil {
		panic(fmt.Sprintf("cannot load config [%s]", file))
	}
	return *conf
}
