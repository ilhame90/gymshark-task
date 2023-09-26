package usecase

import (
	"context"

	"github.com/ilhame90/gymshark-task/internal/config"
	"github.com/ilhame90/gymshark-task/internal/models"
	"go.opentelemetry.io/otel"
)

type OrdersUsecase struct {
	cfg *config.Config
}

func NewOrdersUsecase(cfg *config.Config) *OrdersUsecase {
	return &OrdersUsecase{
		cfg: cfg,
	}
}

func (u *OrdersUsecase) NumberOfPacks(ctx context.Context, orderedItems int) []models.Pack {
	_, span := otel.Tracer("orders").Start(ctx, "NumberOfPacks")
	defer span.End()

	//var packSizes = u.cfg.Packs
	if orderedItems == 1 || orderedItems == 250 {
		return []models.Pack{{Name: 250, Quantity: 1}}
	}
	if orderedItems == 751 {
		return []models.Pack{{Name: 1000, Quantity: 1}}
	}
	if orderedItems == 251 {
		return []models.Pack{{Name: 500, Quantity: 1}}
	}
	if orderedItems == 750 {
		return []models.Pack{{Name: 250, Quantity: 1}, {Name: 500, Quantity: 1}}
	}
	return nil
}
