package card

import (
	"bank/pkg/bank/types"
	"fmt"
)


func ExamplePaymentSources() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active: true,
			PAN: "5058 xxxx xxxx 8888",
		},
		{
			Balance: 10_000_00,
			Active: false,
			PAN: "5058 xxxx xxxx 7777",
		},
		{
			Balance: -10_000_00,
			Active: true,
			PAN: "5058 xxxx xxxx 9999",
		},
	} 
	
	cardList := PaymentSources(cards)
	for _, v := range cardList {
		fmt.Println(v.Number)
	}
	//Output: 5058 xxxx xxxx 8888
	
}  
	