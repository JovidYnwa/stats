package stats

import (
	"fmt"

	"github.com/JovidYnwa/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:     1,
			Amount: 10_000,
			Status: "Clean",
		},
		{
			ID:     2,
			Amount: 20_000,
			Status: "Clean",
		},
		{
			ID:     3,
			Amount: 30_000,
			Status: "StatusFail",
		},
	}

	avg := Avg(payments)
	fmt.Println(avg)
	// Output: 15000
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10_000,
			Category: "Category#1",
			Status:   "Clean",
		},
		{
			ID:       2,
			Amount:   20_000,
			Category: "Category#2",
			Status:   "Clean",
		},
		{
			ID:       3,
			Amount:   30_000,
			Category: "Category#3",
			Status:   "StatusFail",
		},
	}

	category := types.Category("Category#1")
	sum := TotalInCategory(payments, category)
	fmt.Println(sum)
	// Output: 10000
}
