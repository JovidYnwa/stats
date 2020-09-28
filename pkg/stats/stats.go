package stats

import (
	"github.com/JovidYnwa/bank/v2/pkg/types"
)

//Avg рассчитать среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	paySum := types.Money(0)
	payCount := types.Money(0)
	//pyaments := []types.Payment{}
	if len(payments) == 1 {
		return payments[0].Amount
	}

	for i := 0; i < len(payments); i++ {
		if payments[i].Status != types.StatusFail {
			paySum += payments[i].Amount
			payCount++
		}
	}

	return paySum / payCount
}

//TotalInCategory returns total paymet amount of particular category
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := 0

	if len(payments) == 1 {
		return payments[0].Amount
	}

	for _, payment := range payments {
		if category == payment.Category && payment.Status != types.StatusFail {
			sum = sum + int(payment.Amount)
		}
	}

	return types.Money(sum)
}

//FilterByCategory returns payments within marked category
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}
	return filtered
}

//CategoriesTotal return sum of paymets with each category
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {

		categories[payment.Category] += payment.Amount
	}
	return categories
}

//CategoriesAvg calculates avg payment for every category
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	countCategories := map[types.Category]types.Money{}
	sumCategories := map[types.Category]types.Money{}
	avgAmount := map[types.Category]types.Money{}

	for _, payment := range payments {
		sumCategories[payment.Category] += payment.Amount
		countCategories[payment.Category]++
	}
	//fmt.Println(sumCategories, countCategories)

	for keyValue := range sumCategories {
		avgAmount[keyValue] = sumCategories[keyValue] / countCategories[keyValue]
	}

	return avgAmount
}
