package main

import (
	"github.com/26huitailang/golang_web/config"
	"github.com/26huitailang/golang_web/server"
)

func main() {
	e := server.NewServer()
	e.Logger.Fatal(e.Start(config.Config.Port))
}
