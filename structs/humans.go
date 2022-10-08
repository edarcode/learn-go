package structs

type humans struct {
	name string
	age  int
}

func (human humans) GetName() string {
	return human.name
}

func (human humans) GetAge() int {
	return human.age
}
