package TwoSum

import "testing"

func TestTwoSum_1(t *testing.T) {

	nums := []int{2, 7, 11, 15}
	target := 9

	got := TwoSum(nums, target)
	want := []int{0, 1}

	for i := 0; i < 2; i++ {
		if want[i] != got[i] {
			t.Errorf("got %d want %d given", got, want)
		}
	}
}

func TestTwoSum_2(t *testing.T) {

	nums := []int{2, 7, 11, 15}
	target := 13

	got := TwoSum(nums, target)
	want := []int{0, 2}

	for i := 0; i < 2; i++ {
		if want[i] != got[i] {
			t.Errorf("got %d want %d given", got, want)
		}
	}
}

func TestTwoSum_3(t *testing.T) {

	nums := []int{2, 7, 11, 15}
	target := 17

	got := TwoSum(nums, target)
	want := []int{0, 3}

	for i := 0; i < 2; i++ {
		if want[i] != got[i] {
			t.Errorf("got %d want %d given", got, want)
		}
	}
}

func TestTwoSum_4(t *testing.T) {

	nums := []int{2, 7, 11, 15}
	target := 26

	got := TwoSum(nums, target)
	want := []int{2, 3}

	for i := 0; i < 2; i++ {
		if want[i] != got[i] {
			t.Errorf("got %d want %d given", got, want)
		}
	}
}

func TestTwoSum_5(t *testing.T) {

	nums := []int{2, 7, 11, 15}
	target := 18

	got := TwoSum(nums, target)
	want := []int{1, 2}

	for i := 0; i < 2; i++ {
		if want[i] != got[i] {
			t.Errorf("got %d want %d given", got, want)
		}
	}
}

func TestTwoSum_6(t *testing.T) {

	nums := []int{0, 3, 11, 23, 43, 17}
	target := 3

	got := TwoSum(nums, target)
	want := []int{0, 1}

	for i := 0; i < 2; i++ {
		if want[i] != got[i] {
			t.Errorf("got %d want %d given", got, want)
		}
	}
}

func TestTwoSum_7(t *testing.T) {

	nums := []int{0, 3, 11, 23, 43, 17}
	target := 14

	got := TwoSum(nums, target)
	want := []int{1, 2}

	for i := 0; i < 2; i++ {
		if want[i] != got[i] {
			t.Errorf("got %d want %d given", got, want)
		}
	}
}

func TestTwoSum_8(t *testing.T) {

	nums := []int{0, 3, 11, 23, 43, 17}
	target := 40

	got := TwoSum(nums, target)
	want := []int{3, 5}

	for i := 0; i < 2; i++ {
		if want[i] != got[i] {
			t.Errorf("got %d want %d given", got, want)
		}
	}
}
