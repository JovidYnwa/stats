package main

import (
	"fmt"

	"github.com/JovidYnwa/bank/v2/pkg/types"
	"github.com/JovidYnwa/stats/pkg/stats"
)

func main() {
	//categories := []string{"auto", "food", "beauty", "modbile", "fun"}
	//fmt.Println(categories)
	//fmt.Println(len(categories))
	//fmt.Println(cap(categories))
	//top3 := categories[0:3]
	//categories = append(categories, "Jova")
	//categories[1] = "it"
	//fmt.Println(top3)
	//fmt.Println(categories)

	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10_000,
			Category: "Category#1",
			Status:   types.StatusOk,
		},
		{
			ID:       2,
			Amount:   20_000,
			Category: "Category#3",
			Status:   types.StatusOk,
		},
		{
			ID:       3,
			Amount:   30_000,
			Category: "Category#1",
			Status:   types.StatusFail,
		},
	}

	//fmt.Println(payments, "Category#3")
	/* 	categories := map[types.Category]types.Money{
		"auto": 1_000_00,
		"food": 2_000_00,
	} */

	result := stats.CategoriesAvg(payments)
	fmt.Println(result)
}
