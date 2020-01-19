package set

import "fmt"

type Set struct {
	values map[int]bool
}

//Создание пустого множества
func ftCreate() *Set {
	return &Set{
		values: map[int]bool{},
	}
}

//Добавление элемента в множество
func ftAdd(s *Set, value int) bool {
	if s.values == nil {
		fmt.Println("Сначала используйте метод Create")
		return false
	}
	if _, ok := s.values[value]; ok {
		fmt.Println("Элемент уже включён в данное множество")
		return false
	}
	s.values[value] = true
	fmt.Println("Элемент", value, "успешно добавлен")
	return true
}

//Вывод всех элементов множества
func ftPrint(s *Set) {
	if s.values == nil || len(s.values) == 0 {
		fmt.Println("Множество пустое")
		return
	}
	i := 0
	fmt.Print("[")
	for key := range s.values {
		i++
		fmt.Print(key)
		if i < len(s.values) {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
	return
}

//Проверка на наличие в множестве искомого элемента
func ftHas(s *Set, value int) bool {
	if s.values == nil {
		fmt.Println("Сначала используйте метод Create")
		return false
	}
	if _, ok := s.values[value]; ok {
		fmt.Println("Элемент", value, "включён в данное множество")
		return true
	} else {
		fmt.Println("Элемент", value, "не включён в данное множество")
		return false
	}
}

//Очищение множества от всех элементов
func ftClear(s *Set) {
	if s.values == nil || len(s.values) == 0 {
		fmt.Println("Множество пустое")
		return
	}
	for key := range s.values {
		delete(s.values, key)
	}
	fmt.Println("Множество очищено")
}

//Удаление из множества занного элемента
func ftDelete(s *Set, value int) bool {
	if s.values == nil || len(s.values) == 0 {
		fmt.Println("Множество пустое")
		return false
	}
	for key := range s.values {
		if key == value {
			delete(s.values, key)
			fmt.Println("Элемент", value, "удалён")
			return true
		} else {
			fmt.Println("Элемент", value, "не включён в данное множество")
			return false
		}
	}
	return false
}

//Подсчёт количества элементов в множестве
func ftLem(s *Set) int {
	if s.values == nil || len(s.values) == 0 {
		return 0
	}
	return len(s.values)
}
