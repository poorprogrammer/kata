package kata

import "testing"

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

func errorCheck(t *testing.T, result, expected string) {
	if result != expected {
		t.Errorf("expect %s but got %s", expected, result)
	}
}
