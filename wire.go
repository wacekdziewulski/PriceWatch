//+build wireinject

package main

import (
	"github.com/google/wire"
)

func Initialize() *Application {
	wire.Build(NewApplication, NewConfiguration, NewDbConnector, NewProductDao, NewPriceResource)
	return &Application{}
}
