package config

import (
	"flag"
)

func ExtractPortFlag() string {
	portFlag := flag.String("port", "8081", "Port for the auth service")
	flag.Parse()
	port := *portFlag
	return port
}
