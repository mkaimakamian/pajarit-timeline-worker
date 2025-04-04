package main

import (
	"log"
	"pajarit-timeline-worker/config"
	"pajarit-timeline-worker/server"
)

func main() {
	println("loading configuration")
	cfg := config.LoadConfiguration()

	println("loading dependencies")
	deps, err := config.BuildDependencies(cfg)
	if err != nil {
		log.Fatalln("can't load dependencies")
	}

	err = server.StartServer(cfg, deps)
	if err != nil {
		log.Fatalln("can't start server %v", err)
	}

	println("server started in port %d", cfg.ServerPort)
}
