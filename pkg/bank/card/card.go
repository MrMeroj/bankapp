package card

import "bank/pkg/bank/types"

func PaymentSources(cards []types.Card) []types.PaymentSource {
	a := []types.PaymentSource{}
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			a = append(a, types.PaymentSource{
				Type: "card",
				Number: card.PAN,
				Balance: card.Balance,
			})
		}
	}

	return a
}