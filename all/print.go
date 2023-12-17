package formats

import "fmt"

func formats() {
	age := 20;
	score := 231.22313;
	name := "Mike";

	fmt.Print("Hello, ");
	fmt.Println("My name is", name, "and my age is", age);
	fmt.Print("world!");

	// %v - формат по умолчанию
	// %q - формат строки
	// %T - выводит тип
	// %f - плавающая точка
	
	fmt.Printf("My name is %v and my age is %v \n", name, age);
	fmt.Printf("Type of age is %T \n", age);
	fmt.Printf("You'r score is %0.1f", score);

	// Sptringf 

	var str = fmt.Sprintf("My name is %v and my age is %v \n", name, age);
	fmt.Print(str);
}