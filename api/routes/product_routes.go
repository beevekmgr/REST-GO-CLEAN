package routes

import (
	"REST-GO-CLEAN/lib"

	"REST-GO-CLEAN/api/controllers"
)

type ProductRoutes struct {
	Logger     lib.Logger
	Handler    lib.RequestHandler
	Controller controllers.
	Env        lib.Env
}

// Creates a new ProductRoutes
func (
	logger lib.Logger,
	handler lib.RequestHandler,
	controller controllers.DashboardController,
	env lib.Env,
) ProductRoutes {
	return ProductRoutes{
		Logger:         logger,
		Handler:        handler,
		Controller:     controller,
		Env:            env,
	}
}


func (s ProductRoutes) Setup() {
	s.Logger.Zap.Info("Setting up product routes")
	product := s.Handler.Gin.Group("/products").Use()
	{
	product.POST("",s.Controller.CreateProduct())
	}
}