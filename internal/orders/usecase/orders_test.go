package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/ilhame90/gymshark-task/internal/config"
	"github.com/ilhame90/gymshark-task/internal/models"
)

func TestNumberOfPacks(t *testing.T) {

	testCases := []struct {
		name         string
		packs        []int
		orderedItems int
		expected     []models.Pack
		result       []models.Pack
	}{

		{
			name:         "ORDER: 1 --- PACKS: 250, 500, 1000, 2000, 5000",
			orderedItems: 1,
			packs:        []int{250, 500, 1000, 2000, 5000},
			expected:     []models.Pack{{Name: 250, Quantity: 1}},
		},
		{
			name:         "ORDER: 250 --- PACKS: 250, 500, 1000, 2000, 5000",
			orderedItems: 250,
			packs:        []int{250, 500, 1000, 2000, 5000},
			expected:     []models.Pack{{Name: 250, Quantity: 1}},
		},

		{
			name:         "ORDER: 251 --- PACKS: 250, 500, 1000, 2000, 5000",
			orderedItems: 251,
			packs:        []int{250, 500, 1000, 2000, 5000},
			expected:     []models.Pack{{Name: 500, Quantity: 1}},
		},
		{
			name:         "ORDER: 750 --- PACKS: 250, 500, 1000, 2000, 5000",
			orderedItems: 750,
			packs:        []int{250, 500, 1000, 2000, 5000},
			expected:     []models.Pack{{Name: 500, Quantity: 1}, {Name: 250, Quantity: 1}},
		},
		{
			name:         "ORDER: 751 --- PACKS: 250, 500, 1000, 2000, 5000",
			orderedItems: 751,
			packs:        []int{250, 500, 1000, 2000, 5000},
			expected:     []models.Pack{{Name: 1000, Quantity: 1}},
		},
		{
			//*! This case fails =)
			name:         "ORDER: 751 --- PACKS: 250, 500, 750, 1000, 2000, 5000",
			orderedItems: 751,
			packs:        []int{250, 500, 750, 1000, 2000, 5000},
			expected:     []models.Pack{{Name: 1000, Quantity: 1}},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			cfg := &config.Config{Packs: tc.packs}
			usecase := NewOrdersUsecase(cfg)
			result := usecase.NumberOfPacks(context.Background(), tc.orderedItems)
			if !testEq(tc.expected, result) {
				t.Errorf("For ordered_items=%v, packs=%v  , expected %v, but got %v", tc.orderedItems, tc.packs, tc.expected, result)
			}

		})
	}

}

func testEq(expected, actual []models.Pack) bool {
	equal := cmp.Equal(expected, actual, cmpopts.SortSlices(func(a, b models.Pack) bool { return a.Name < b.Name }))
	return equal
}
