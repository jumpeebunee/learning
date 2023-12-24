package main

import "fmt"

type gifts struct {
	holiday  string
	giftList map[string]float64
	funScore float64
}

// Получение листа с подарками
func (g gifts) formattList() string {
	var totalPrice float64 = 0
	fs := "The Zhigalove gift list: \n";

	for k, v := range g.giftList {
		fs += fmt.Sprintf("%-25v ...%0.2f \n", k + ":", v);
		totalPrice += v
	}

	fs += fmt.Sprintf("%-25v %v points \n", "Score", g.funScore);
	fs += fmt.Sprintf("%-25v ...$%0.2f", "Total price is:", totalPrice);

	return fs;
}

func (g *gifts) addGift(gift string, price float64) {
	(*g).giftList[gift] = price;
}

func (g *gifts) updateScore(score float64) {
	(*g).funScore = score;
}

func generateGiftList(holiday string, score float64) gifts {
	return gifts{
		holiday: holiday,
		giftList: map[string]float64{},
		funScore: score,
	}
}
