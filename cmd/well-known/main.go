package main

import (
	"flag"
	"github.com/captnCC/well-known/internal/config"
	"github.com/captnCC/well-known/internal/server"
	"log"
)

func main() {

	var configPath string
	var port int

	flag.StringVar(&configPath, "c", "/etc/well-known/config.yaml", "Path to config file")
	flag.IntVar(&port, "p", 3000, "Port to listen on")

	flag.Parse()

	err, endpointConfig := config.ReadConfig(configPath)

	if err != nil {
		log.Fatal(err)
		return
	}

	server.Serve(endpointConfig, port, true)

	log.Printf("Bye bye")
}
