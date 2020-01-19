package twoSum

func twoSum(nums []int, target int) []int {
	//Создание слайса под результат
	solution := make([]int, 0, 2)
	//Создание первого цикла для итерации по значениям
	for i := range nums {
		//Создание второго цикла для итерации по значениям
		for j := i + 1; j < len(nums); j++ {
			//Сравнение суммы двух чисел с целью
			if nums[i]+nums[j] == target {
				//Добавление в слайс результата (позиций искомых чисел)
				solution = append(solution, i, j)
				return solution
			}
		}
	}
	return solution
}
