package converter

import (
	"testing"
)

func TestInput1ShouldGetI(t *testing.T) {
	result := roman(1)
	errorCheck(t, result, "I")
}

func TestInput3ShoudGetIII(t *testing.T) {
	result := roman(3)
	errorCheck(t, result, "III")
}

func Test4ShouldGetIX(t *testing.T) {
	result := roman(4)
	errorCheck(t, result, "IV")
}

func Test5ShouldGetX(t *testing.T) {
	result := roman(5)
	errorCheck(t, result, "V")
}
func Test6ShouldGetX(t *testing.T) {
	result := roman(6)
	errorCheck(t, result, "VI")
}

func Test7ShouldGetVII(t *testing.T) {
	result := roman(7)
	errorCheck(t, result, "VII")
}
func TestSum2Number(t *testing.T) {
	result := sumRomanNumber("I", "IV")
	errorCheck(t, result, "V")
}

func errorCheck(t *testing.T, result, expected string) {
	if result != expected {
		t.Errorf("expect %s but got %s", expected, result)
	}
}

func sumRomanNumber(n1 string, n2 string) string {
	sumRomanNumber := romanStringToInt(n1) + romanStringToInt(n2)
	return roman(sumRomanNumber)
}

func romanStringToInt(n string) int {
	romanMap := map[string]int{
		"I":   1,
		"IV":  4,
		"V":   5,
		"VI":  6,
		"VII": 7,
		"IX":  9,
	}

	result := 0
	for i := 0; i < len(romanMap); i++ {
		if romanMap[n] == i {
			result = romanMap[n]
		}
	}

	return result
}

func roman(n int) string {

	romanMap := map[int]string{
		1: "I",
		4: "IV",
		5: "V",
		6: "VI",
		7: "VII",
		9: "IX",
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
