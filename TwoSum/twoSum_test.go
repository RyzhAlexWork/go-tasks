package twoSum

import "testing"

var (
	nums = [][]int{
		[]int{2, 7, 11, 15},
		[]int{0, 3, 11, 23, 43, 17},
	}
	targets = [][]int{
		[]int{9, 13, 17, 26, 18},
		[]int{3, 14, 40},
	}
	wantss = [][][]int{
		[][]int{
			[]int{0, 1},
			[]int{0, 2},
			[]int{0, 3},
			[]int{2, 3},
			[]int{1, 2},
		},
		[][]int{
			[]int{0, 1},
			[]int{1, 2},
			[]int{3, 5},
		},
	}
)

func TestTwoSum(t *testing.T) {
	for i, wants := range wantss {
		for j, want := range wants {
			got := twoSum(nums[i], targets[i][j])
			for k := range want {
				if want[k] != got[k] {
					t.Errorf("got: %d, want: %d", got, want)
				}
			}
		}
	}
}
