package TwoSum

var (
	solution []int = nil
)

func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				solution := append(solution, i, j)
				return solution
			}
		}
	}
	return solution
}
