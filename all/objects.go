func objects() {
	menu := map[string]float64{
		"apple": 30.1,
		"orange": 20,
	}

	menu["apple"] = 42;

	for k, v := range menu {
		fmt.Println("Key is:", k, "value is", v);
	}
}