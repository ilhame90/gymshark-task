package orders

//go:generate mockgen -source ./usecase.go -package mocks -destination ../domain/mocks/usecase.mock.gen.go
import (
	"context"

	"github.com/ilhame90/gymshark-task/internal/models"
)

type Usecase interface {
	NumberOfPacks(ctx context.Context, orderedItems int) []models.Pack
}
