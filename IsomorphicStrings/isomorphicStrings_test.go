package IsomorphicStrings

import "testing"

func TestIsIsomorphic_1(t *testing.T) {
	str1 := "egg"
	str2 := "add"

	got := IsIsomorphic(str1, str2)
	want := true

	if want != got {
		t.Errorf("got %v want %v given", got, want)
	}
}

func TestIsIsomorphic_2(t *testing.T) {
	str1 := "foo"
	str2 := "bar"

	got := IsIsomorphic(str1, str2)
	want := false

	if want != got {
		t.Errorf("got %v want %v given", got, want)
	}
}

func TestIsIsomorphic_3(t *testing.T) {
	str1 := "paper"
	str2 := "title"

	got := IsIsomorphic(str1, str2)
	want := true

	if want != got {
		t.Errorf("got %v want %v given", got, want)
	}
}

func TestIsIsomorphic_4(t *testing.T) {
	str1 := "ab"
	str2 := "aa"

	got := IsIsomorphic(str1, str2)
	want := false

	if want != got {
		t.Errorf("got %v want %v given", got, want)
	}
}

func TestIsIsomorphic_5(t *testing.T) {
	str1 := "aa"
	str2 := "ab"

	got := IsIsomorphic(str1, str2)
	want := false

	if want != got {
		t.Errorf("got %v want %v given", got, want)
	}
}

func TestIsIsomorphic_6(t *testing.T) {
	str1 := "aabbccaa"
	str2 := "ooppggoo"

	got := IsIsomorphic(str1, str2)
	want := true

	if want != got {
		t.Errorf("got %v want %v given", got, want)
	}
}

func TestIsIsomorphic_7(t *testing.T) {
	str1 := "aabbcccc"
	str2 := "ooppggoo"

	got := IsIsomorphic(str1, str2)
	want := false

	if want != got {
		t.Errorf("got %v want %v given", got, want)
	}
}
