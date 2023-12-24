package main

import "fmt"

func main() {
	gift := generateGiftList("birthday", 10)
	fl := gift.formattList();

	fmt.Println(gift);
	fmt.Print(fl);
}