package routes

import (
	nicepay "nicepay-api/app/domain/nicepay"

	"github.com/labstack/echo/v4"
)

// navigate app
func NewRoutes(e *echo.Echo) {
	api := e.Group("/api")

	api.POST("/registration", nicepay.Registration)
	api.POST("/payment", nicepay.Payment)
	api.POST("/status", nicepay.Status)
	api.POST("/cancel", nicepay.Cancel)
}
