package converter

func roman(n int) string {

	romanMap := map[int]string{
		1: "I",
		4: "IV",
		5: "V",
		6: "VI",
		7: "VII",
	}
	if _, ok := romanMap[n]; ok {
		return romanMap[n]
	}

	result := ""
	for i := 0; i < n; i++ {
		result = result + romanMap[1]
	}

	return result
}
