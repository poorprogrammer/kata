package converter

import "testing"

const (
	ONE = "I"
)

func TestInput1ShouldGetI(t *testing.T) {
	result := roman(1)
	errorCheck(t, result, ONE)
}

func TestInput2ShoudGetII(t *testing.T) {
	result := roman(2)
	errorCheck(t, result, "II")
}

func TestInput3ShoudGetIII(t *testing.T) {
	result := roman(3)
	errorCheck(t, result, "III")
}

func errorCheck(t *testing.T, result, expected string) {
	if result != expected {
		t.Errorf("expect %s but got %s", expected, result)
	}
}

func roman(n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result = result + ONE
	}
	return result
}
