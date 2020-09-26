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
		if payments[i].Status != "StatusFail" {
			paySum += payments[i].Amount
			payCount++
		}
	}

	return paySum / payCount
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := 0

	if len(payments) == 1 {
		return payments[0].Amount
	}

	for _, payment := range payments {
		if category == payment.Category && payment.Status != "StatusFail" {
			sum = sum + int(payment.Amount)
		}
	}

	return types.Money(sum)
}
