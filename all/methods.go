package methods

import (
	"fmt"
	"sort"
	"strings"
)

func methods() {
	name := "Hello world from USA";

	fmt.Println(strings.Contains(name, "USA"));
	fmt.Println(strings.ReplaceAll(name, "Hello", "Hi"));
	fmt.Println(strings.ToUpper(name));
	fmt.Println(strings.Index(name, "ll"));
	fmt.Println(strings.Split(name, " "));

	ages := []int{45, 20, 30, 12}
	names := []string{"z", "c", "a", "v"};
	index := sort.SearchInts(ages, 30);
	nameIndex := sort.SearchStrings(names, "c");

	sort.Ints(ages);
	sort.Strings(names);

	fmt.Print(ages);
	fmt.Print(index);
	fmt.Print(names);
	fmt.Print(nameIndex);
}