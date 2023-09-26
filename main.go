package main

import (
	"fmt"
	"math"
)

func minPacks(order int, packSizes []int) (map[int]int, int) {
	// Initialize a table to store the minimum number of packs needed for each order.
	dp := make([]int, order+1)
	for i := 1; i <= order; i++ {
		dp[i] = math.MaxInt // Initialize to a large value.
	}

	// Initialize a table to store the pack sizes used for each order.
	usedPacks := make([]int, order+1)

	// Calculate the minimum number of packs needed for each order from 1 to order.
	for i := 1; i <= order; i++ {
		for _, packSize := range packSizes {
			if i >= packSize && dp[i-packSize]+1 < dp[i] {
				dp[i] = dp[i-packSize] + 1
				usedPacks[i] = packSize
			}
		}
	}

	// Calculate the pack counts from the usedPacks table.
	packCounts := make(map[int]int)
	for _, packSize := range packSizes {
		packCounts[packSize] = 0
	}
	remainingOrder := order
	for remainingOrder > 0 {
		packSize := usedPacks[remainingOrder]
		packCounts[packSize]++
		remainingOrder -= packSize
	}

	return packCounts, dp[order]
}

func main() {
	order := 251
	packSizes := []int{250, 500, 1000, 2000, 5000}

	packCounts, minPacksNeeded := minPacks(order, packSizes)

	fmt.Printf("Minimum packs needed to fulfill the order of %d items: %d\n", order, minPacksNeeded)
	fmt.Println("Packs needed:")
	for packSize, count := range packCounts {
		if count > 0 {
			fmt.Printf("%dx%d\n", count, packSize)
		}
	}
}
