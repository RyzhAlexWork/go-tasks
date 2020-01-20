package factoryMethod

import "testing"

var (
	actions = []string{"Student", "Businessman", "Programmer"}
	names   = []string{"Alex", "Mark", "Ivan"}
	ages    = []int{23, 45, 28}
	factory = NewCreater()
	persons = []Person{
		factory.CreatePerson("Student", "Alex", 23),
		factory.CreatePerson("Businessman", "Mark", 45),
		factory.CreatePerson("Programmer", "Ivan", 28),
	}
)

func TestFactoryMethod(t *testing.T) {
	for i, person := range persons {
		if action := person.GetAction(); action != actions[i] {
			t.Errorf("Expect action to %s, but %s.\n", actions[i], action)
		}
		if name := person.GetName(); name != names[i] {
			t.Errorf("Expect name to %s, but %s.\n", names[i], name)
		}
		if age := person.GetAge(); age != ages[i] {
			t.Errorf("Expect age to %d, but %d.\n", ages[i], age)
		}
	}
}
