package stats

import (
	"fmt"

	"github.com/JovidYnwa/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:     1,
			Amount: 10_000,
		},
		{
			ID:     2,
			Amount: 20_000,
		},
		{
			ID:     3,
			Amount: 30_000,
		},
	}

	avg := Avg(payments)
	fmt.Println(avg)
	// Output: 20000
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10_000,
			Category: "Category#1",
		},
		{
			ID:       2,
			Amount:   20_000,
			Category: "Category#2",
		},
		{
			ID:       3,
			Amount:   30_000,
			Category: "Category#3",
		},
	}

	category := types.Category("Category#1")
	sum := TotalInCategory(payments, category)
	fmt.Println(sum)
	// Output: 10000
}
