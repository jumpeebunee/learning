package main

import "fmt"

type gifts struct {
	holiday  string
	giftList map[string]float64
	funScore float64
}

func (g gifts) formattList() string {
	var totalPrice float64 = 0
	fs := "The Zhigalove gift list: \n";

	for k, v := range g.giftList {
		fs += fmt.Sprintf("%-25v ...%0.2f \n", k + ":", v);
		totalPrice += v
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f", "Total price is:", totalPrice);

	return fs;
}

func generateGiftList(holiday string, score float64) gifts {
	return gifts{
		holiday: holiday,
		giftList: map[string]float64{
			"Minion":     2603.24,
			"Helicopter": 323.90,
			"Hamster":    232.23,
		},
		funScore: score,
	}
}