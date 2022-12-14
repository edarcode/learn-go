package structs

import "fmt"

type users struct {
	humans
	role string
}

func CreateUser(name string, age int, role string) users {
	propsHuman := humans{name, age}
	return users{humans: propsHuman, role: role}
}

func (user users) GetRole() string {
	return user.role
}

func (user users) Greet() string {
	msg := fmt.Sprintf("Hola soy %s me encanta comer pastas y carne", user.name)
	return msg
}
