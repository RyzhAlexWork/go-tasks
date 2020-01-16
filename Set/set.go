package set

import "fmt"

type Set struct {
	values map[int]bool
}

func Create() *Set {
	return &Set{
		values: map[int]bool{},
	}
}

func Add(s *Set, value int) bool {
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

func Print(s *Set) {
	if s.values == nil || len(s.values) == 0 {
		fmt.Println("Множество пустое")
		return
	}
	i := 0
	fmt.Print("[")
	for key, _ := range s.values {
		i++
		fmt.Print(key)
		if i < len(s.values) {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
	return
}

func Has(s *Set, value int) bool {
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

func Clear(s *Set) {
	if s.values == nil || len(s.values) == 0 {
		fmt.Println("Множество пустое")
		return
	}
	for key, _ := range s.values {
		delete(s.values, key)
	}
	fmt.Println("Множество очищено")
}

func Delete(s *Set, value int) bool {
	if s.values == nil || len(s.values) == 0 {
		fmt.Println("Множество пустое")
		return false
	}
	for key, _ := range s.values {
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

func Lem(s *Set) int {
	if s.values == nil || len(s.values) == 0 {
		return 0
	}
	i := 0
	for _, _ = range s.values {
		i++
	}
	return i
}
