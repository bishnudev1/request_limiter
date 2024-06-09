package routes

import (
	"request_limiter/api"
	middlewares "request_limiter/middlewars"

	"github.com/gofiber/fiber/v2"
)

func ItemRoutes(app *fiber.App) {
	app.Get("/items", middlewares.RateLimitMiddleware(
		30, 10,
	), api.GetItems)
}
