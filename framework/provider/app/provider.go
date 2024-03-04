package app

import "github.com/26huitailang/golang_web/framework"

type YogoAppProvider struct {
}

func (p *YogoAppProvider) Boot(c framework.Container) error {
	return nil
}

func (p *YogoAppProvider) IsDefer() bool {
	return false
}

func (p *YogoAppProvider) Params(c framework.Container) []interface{} {
	return nil
}

func (p *YogoAppProvider) Name() string {
	return "app"
}

func (p *YogoAppProvider) Register(c framework.Container) framework.NewInstance {
	return nil
}
