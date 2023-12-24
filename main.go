package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var list gifts;


func getInput() string {
	reader := bufio.NewReader(os.Stdin);
	input, _ := reader.ReadString('\n');
	input = strings.TrimSpace(input);

	return input;
}

func validateOption(option string) {
	switch option {
	case "s": 
		saveCheck();
	case "a": 
		addNewItem();
		chooseOption();
	case "p" :
		addPoints();
		chooseOption();
	}
}

func chooseOption() {
	fmt.Print("Choose option (a - add item, s - save gift, p - add points): ");
	option := getInput();
	validateOption(option);
}

func addNewItem() {
	fmt.Print("Item name: ");
	name := getInput();
	fmt.Print("Item price: ");
	price := getInput();

	val, _ := strconv.ParseFloat(price, 64);

	list.addGift(name, val);
	fmt.Printf("Item added: %v, price: %v rub \n", name, price);
}

func saveCheck() {
	fmt.Println("Check has been saved");
	fs := list.formattList();
	fmt.Print(fs);
}

func addPoints() {
	fmt.Print("Add points: ");
	points := getInput();
	val, _ := strconv.ParseFloat(points, 64);

	list.updateScore(val);
	fmt.Printf("Points has been updated to %v \n", points);
}

func createNewGift() gifts {
	fmt.Print("Create a new gift name for friend: ");

	name := getInput();

	fmt.Printf("Create new gift for - %v \n", name);

	gl := generateGiftList(name, 0);
	return gl;
}

func main() {
	list = createNewGift();
	chooseOption();
}