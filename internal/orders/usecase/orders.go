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

	packs := calculate(u.cfg.Packs, orderedItems)
	res := []models.Pack{}
	for k, v := range packs {
		if v == 0 {
			continue
		}
		res = append(res, models.Pack{
			Name:     k,
			Quantity: v,
		})
	}
	return res
}

func calculate(packs []int, order int) map[int]int {
	orderQty := order
	res := make(map[int]int)

	for i := len(packs) - 1; i >= 0; i-- {
		pack := packs[i]

		// skip packs greater than order
		if orderQty < pack {
			if i == 0 && orderQty > 0 {
				res[pack]++
			}
			continue
		}

		packQty := orderQty / pack
		orderQty = orderQty % pack

		res[pack] = packQty

		if i == 0 && orderQty > 0 {
			res[pack]++
		}
	}

	return optimizeOrder(packs, res)
}

func optimizeOrder(packs []int, order map[int]int) map[int]int {
	packMap := make(map[int]struct{})
	for _, pack := range packs {
		packMap[pack] = struct{}{}
	}

	for pack, qty := range order {
		if qty <= 1 {
			continue
		}
		for {
			checkPack := pack * qty
			_, exists := packMap[checkPack]
			if !exists || qty <= 1 {
				order[pack] = qty
				break
			}

			order[pack] -= qty
			order[checkPack]++

			pack = checkPack
			qty = order[checkPack]

		}

	}

	return order
}
