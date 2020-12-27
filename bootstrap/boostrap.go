package bootstrap

import (
	"REST-GO-CLEAN/api/controllers"
	"REST-GO-CLEAN/api/repository"
	"REST-GO-CLEAN/api/routes"
	"REST-GO-CLEAN/api/services"
	"REST-GO-CLEAN/lib"
	"REST-GO-CLEAN/models"
	"context"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	controllers.Module,
	repository.Module,
	routes.Module,
	services.Module,
	lib.Module,
	models.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler lib.RequestHandler,
	routes routes.Routes,
	env lib.Env,
	logger lib.Logger,
	database lib.Database,
	migrations models.Migrations,
) {

	_, cancel := context.WithCancel(context.Background())
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application")
			logger.Zap.Info("-----------------------------")
			logger.Zap.Info("------- SHOP-API 🚀 -------")
			logger.Zap.Info("-----------------------------")

			logger.Zap.Info("Migrating database schemas")
			migrations.Migrate()

			go func() {
				routes.Setup()
				if env.ServerPort == "" {
					handler.Gin.Run()
				} else {
					handler.Gin.Run(env.ServerPort)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {

			logger.Zap.Info("Stopping Application")
			database.DB.Close()

			logger.Zap.Info("Closing context")
			cancel()
			return nil
		},
	})

}
