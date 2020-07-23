package main

import (
	"golang_web/config"
	"golang_web/server"
)

func main() {
	e := server.NewServer()
	e.Logger.Fatal(e.Start(config.Config.Port))
}
