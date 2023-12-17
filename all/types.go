package types

import "fmt"

func types() {
	// strings

	var nameOne string = "Mario"
	var nameTwo = "John"
	var nameThree string
	nameFour := "Bogdan" // Нельзя использовать вне функции;

	fmt.Print(nameOne, nameTwo, nameThree, nameFour)

	// ints

	var ageOne int = 20;
	var ageTwo = 30;
	var ageThree int;

	fmt.Print(ageOne, ageTwo, ageThree)

	// bytes

	var intOne int8 = 25 // -128 / 127
	var intTwo uint8 = 78

	fmt.Print(intOne, intTwo);

	// float

	var scoreOne float32 = 23123123.321;
	var scoreTwo float64 = 12312312312312321.3;

	fmt.Print(scoreOne, scoreTwo);
}