package stats

import (
	"github.com/Shah-2021/bankapp.git/pkg/bank/types"
)

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	avgSum := types.Money(0)
	ind := types.Money(0)
	for _, v := range payments {
		avgSum += v.Amount
		ind ++
	}

	return types.Money(avgSum/ind)
}
