package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput() string {
	reader := bufio.NewReader(os.Stdin);
	name, _ := reader.ReadString('\n');
	name = strings.TrimSpace(name);

	return name;
}

func createNewGift() gifts {
	fmt.Print("Create a new gift name for friend: ");

	name := getInput();

	fmt.Printf("Create new gift for - %v \n", name);

	gl := generateGiftList(name, 0);
	return gl;
}

func main() {
	giftList := createNewGift();
	fmt.Print(giftList);
}