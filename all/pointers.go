import "fmt"

func updateName(name *string) {
	*name = "alex"
}

func pointers() {
	name := "vova"
	m := &name

	updateName(m);
	fmt.Print(name);
}