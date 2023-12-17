package arrays

import "fmt"

func arrays() {
	// var array [3]int = [3]int{24, 26, 13};
	var arrayTwo = [3]int{12, 2, 3};

	fmt.Print(arrayTwo, len(arrayTwo));

	// slices

	var score = []int{2, 50, 92, 23, 1};
	score[1] = 660;

	score = append(score, 20);

	// slice

	var rangeOne = score[1:3];
	var rangeTwo = score[1:];
	var rangeThree = score[:2];
	
	fmt.Print(rangeOne);
	fmt.Print(rangeTwo);
	fmt.Print(rangeThree);
}