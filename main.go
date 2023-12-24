package main

import "fmt"

func main() {
	gift := generateGiftList("birthday", 10)
	gift.updateScore(100);

	gift.addGift("Minion",  2603.24);
	gift.addGift("Helicopter",  323.90);
	gift.addGift("Hamster",  232.23);


	fl := gift.formattList();

	fmt.Println(gift);
	fmt.Print(fl);
}