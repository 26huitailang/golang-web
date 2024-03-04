package mygin

import (
	"github.com/26huitailang/golang_web/framework"
	"github.com/gin-gonic/gin"
)

type YogoEngin struct {
	*gin.Engine
	container framework.Container
}

func (engine *YogoEngin) SetContainer(container framework.Container) {
	engine.container = container
}

func (engine *YogoEngin) GetContainer() framework.Container {
	return engine.container
}

func (engine *YogoEngin) Bind(provider framework.ServiceProvider) error {
	return engine.container.Bind(provider)
}

func (engine *YogoEngin) IsBind(key string) bool {
	return engine.container.IsBind(key)
}
