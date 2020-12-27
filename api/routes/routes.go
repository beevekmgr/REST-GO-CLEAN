package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewProductRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

func NewRoutes(
	product ProductRoutes,
) Routes {
	return Routes{
		product,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
