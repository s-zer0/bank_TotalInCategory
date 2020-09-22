package stats

import "github.com/s-zer0/bank/pkg/bank/types"

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	avg := types.Money(0)
	quantity := types.Money(0)
	for _, v := range payments {
		avg += v.Amount
		quantity ++
	}
	return avg/quantity
}