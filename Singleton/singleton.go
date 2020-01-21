package singleton

import "sync"

type Government struct {
	name   string
	age    int
	salary int
}

var (
	government *Government
	once       sync.Once
)

func GetGovernment() *Government {
	once.Do(func() {
		government = &Government{"Михаил Владимирович Мишустин", 53, 1500000}
	})
	return government
}
