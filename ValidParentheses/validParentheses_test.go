package validParentheses

import "testing"

var (
	testStrs = []string{"()", "(()", "())", "()[]{}", "()}", "[{}", "{)", "}]", "(", "({}[])"}
	wants    = []bool{true, false, false, true, false, false, false, false, false, true}
)

func TestIsValid(t *testing.T) {
	for i, want := range wants {
		got := isValid(testStrs[i])
		if got != want {
			t.Errorf("Expect want to %v, but %v. String: %s\n", want, got, testStrs[i])
		}
	}
}
