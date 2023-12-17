package loops

import "fmt"

func loops() {
	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}

	names := []string{"vova", "alex", "petr"}

	for value, index := range names {
		result := fmt.Sprintf("The value of index %v, is %v \n", index, value)
		fmt.Print(result)
	}
}