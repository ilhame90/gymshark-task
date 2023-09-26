package server

import (
	"net/http"

	ordersHTTP "github.com/ilhame90/gymshark-task/internal/orders/delivery/http"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, handler *ordersHTTP.OrdersHandler) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
	})
	grp := e.Group("/v1")

	ordersHTTP.RegisterHandlers(grp, handler)
}
