package set

import "testing"

func TestCreate(t *testing.T) {
	got := ftCreate()

	if got.values == nil {
		t.Errorf("Множество не было создано")
	}
}

func TestAdd_1(t *testing.T) {
	s := ftCreate()
	got := ftAdd(s, 5)

	if got == false {
		t.Errorf("Элемент не был добавлен")
	}
}

func TestAdd_2(t *testing.T) {
	s := ftCreate()
	ftAdd(s, 5)
	got := ftAdd(s, 5)

	if got == true {
		t.Errorf("Элемент был добавлен, но он уже содержался в данном множестве")
	}
}

func TestHas_1(t *testing.T) {
	s := ftCreate()
	ftAdd(s, 5)
	got := ftHas(s, 5)

	if got == false {
		t.Errorf("Элемент не был найден, но он содержится в данном множестве")
	}
}

func TestHas_2(t *testing.T) {
	s := ftCreate()
	ftAdd(s, 5)
	got := ftHas(s, 7)

	if got == true {
		t.Errorf("Элемент был найден, но он не содержится в данном множестве")
	}
}

func TestClear(t *testing.T) {
	s := ftCreate()
	ftAdd(s, 5)
	ftAdd(s, 7)
	ftClear(s)

	if ftLem(s) != 0 {
		t.Errorf("Множество не очистилось")
	}
}

func TestDelete_1(t *testing.T) {
	s := ftCreate()
	ftAdd(s, 5)
	ftAdd(s, 7)
	ftDelete(s, 5)

	if ftHas(s, 5) == true {
		t.Errorf("Элемент не удалён")
	}
}

func TestDelete_2(t *testing.T) {
	s := ftCreate()
	ftAdd(s, 5)
	ftAdd(s, 7)
	got := ftDelete(s, 8)

	if got == true {
		t.Errorf("Элемент не включался в множество, но был удалён")
	}
}

func TestLem(t *testing.T) {
	s := ftCreate()
	ftAdd(s, 5)
	ftAdd(s, 7)
	ftAdd(s, 8)
	got := ftLem(s)

	if got != 3 {
		t.Errorf("Количество элементов посчитано неверно")
	}
}
