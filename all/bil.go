type bill struct {
	name string
	items map[string]int
	tip float64
}

func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]int{},
		tip: 0,
	}
}