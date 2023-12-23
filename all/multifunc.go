func getName(n string) (string, string) {
	n = strings.ToUpper(n);
	names := strings.Split(n, " ");

	ini := []string{}

	for _, val := range names {
		firstName := val[:1];
		ini = append(ini, firstName);
	}
	
	
	if len(ini) >= 2 {
		return ini[0], ini[1]
	}

	return ini[0], "_";
}
