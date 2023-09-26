package http

import (
	"net/http"

	"github.com/ilhame90/gymshark-task/internal/config"

	"github.com/ilhame90/gymshark-task/internal/orders"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
)

type OrdersHandler struct {
	cfg      *config.Config
	ordersUC orders.Usecase
}

func NewOrdersHandler(cfg *config.Config, ordersUC orders.Usecase) *OrdersHandler {
	return &OrdersHandler{
		cfg:      cfg,
		ordersUC: ordersUC,
	}
}

func (h *OrdersHandler) GetHealthcheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{"status": "ok"})
}

func (h *OrdersHandler) PostOrder(ctx echo.Context) error {
	c, span := otel.Tracer("orders").Start(ctx.Request().Context(), "PostOrder")
	defer span.End()

	body := &PostOrderJSONRequestBody{}
	if err := ctx.Bind(body); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		)
	}
	if body.OrderedItems < 1 {
		return ctx.JSON(http.StatusBadRequest,
			Error{
				Code:    http.StatusBadRequest,
				Message: "order amount should be greater that 0",
			},
		)
	}

	packs := h.ordersUC.NumberOfPacks(c, body.OrderedItems)

	res := make([]Pack, len(packs))

	for i, pack := range packs {
		res[i] = Pack{
			Name:     pack.Name,
			Quantity: pack.Quantity,
		}
	}

	return ctx.JSON(http.StatusOK, res)
}
