package isomorphicStrings

import "testing"

var (
	inputs = [][]string{
		[]string{"egg", "add"},
		[]string{"foo", "bar"},
		[]string{"paper", "title"},
		[]string{"ab", "aa"},
		[]string{"aa", "ab"},
		[]string{"aabbccaa", "ooppggoo"},
		[]string{"aabbcccc", "ooppggoo"},
	}
	wants = []bool{true, false, true, false, false, true, false}
)

func TestIsIsomorphic(t *testing.T) {
	for i, want := range wants {
		got := isIsomorphic(inputs[i][0], inputs[i][1])
		if want != got {
			t.Errorf("got: %v, want: %v, input_1: %v, input_2: %v", got, want, inputs[i][0], inputs[i][1])
		}
	}
}
