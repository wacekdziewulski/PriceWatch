//+build wireinject

package main

import (
	"PriceWatch/configuration"
	"PriceWatch/db"
	"PriceWatch/resource"
	"PriceWatch/service"

	"github.com/google/wire"
)

func Initialize() *Application {
	wire.Build(NewApplication, configuration.NewConfiguration, db.NewConnector, db.NewProductDao, service.NewPriceService, service.NewURLShorteningService, resource.NewPriceResource)
	return &Application{}
}
