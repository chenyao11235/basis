package facade

import (
	"fmt"
)

//API 对对外提供的接口(门面接口)
type API interface {
	Request() string
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

//NewAPI 新建
func NewAPI() API {
	return &apiImpl{
		a: &aModuleImpl{},
		b: &bModuleImpl{},
	}
}

func (api *apiImpl) Request() string {
	resultA := api.a.RequestA()
	resultB := api.b.RequestB()

	return fmt.Sprintf("%s,%s", resultA, resultB)
}

//AModuleAPI 模块A接口
type AModuleAPI interface {
	RequestA() string
}

type aModuleImpl struct {
}

func (a *aModuleImpl) RequestA() string {
	return "a module"
}

//BModuleAPI 模块B接口
type BModuleAPI interface {
	RequestB() string
}

type bModuleImpl struct {
}

func (b *bModuleImpl) RequestB() string {
	return "b module"
}
