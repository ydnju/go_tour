package basictypes

import "testing"

func TestStringBasic(t *testing.T) {
	str := "hahag"
	if len(str) != 5 {
		t.Error(`len(str) != 5`)
	}
}

func TestCommaNoRecur(t *testing.T) {
	res := NoRecursiveComma("12345")
	if res != "12,345" {
		t.Errorf(`NoRecurComman("12345") != "12,345", actual is %s`, res)
	}
}
