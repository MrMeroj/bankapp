package card

import "bank/pkg/bank/types"

func Total(cards []types.Card) types.Money  {
	total := types.Money(0) 
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			total += card.Balance
		}
	}

	return total
}