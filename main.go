package main

import (
	//"github.com/26huitailang/golang_web/app/console"
	"github.com/26huitailang/golang_web/app/http"
	"github.com/26huitailang/golang_web/framework"
	"github.com/26huitailang/golang_web/framework/provider/app"
	"github.com/26huitailang/golang_web/framework/provider/kernel"
)

func main() {
	container := framework.NewYogoContainer()
	container.Bind(&app.YogoAppProvider{})

	if engine, err := http.NewHttpEngine(container); err == nil {
		container.Bind(&kernel.YogoKernelProvider{HttpEngine: engine})
	}
	console.RunCommand(container)
}
