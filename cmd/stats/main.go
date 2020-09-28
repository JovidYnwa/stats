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

	firstPaymet := map[types.Category]types.Money{
		"Category1": 10,
		"Category2": 20,
	}

	secondPayment := map[types.Category]types.Money{
		"Category1": 10,
		"Category2": 15,
		"Category3": 5,
	}
	//fmt.Println(payments, "Category#3")
	/* 	categories := map[types.Category]types.Money{
		"auto": 1_000_00,
		"food": 2_000_00,
	} */

	//result := stats.CategoriesAvg(payments)
	result := stats.PeriodsDynamic(firstPaymet, secondPayment)

	fmt.Println(result)
}
