package structs

import "fmt"

type singers struct {
	humans
	nickname string
}

func CreateSinger(name string, age int, nickname string) singers {
	return singers{humans: humans{name: name, age: age}, nickname: nickname}
}

func (singer singers) GetNickname() string {
	return singer.nickname
}

func (singer singers) Greet() string {
	msg := fmt.Sprintf("Hola soy %s y me encanta cantar", singer.nickname)
	return msg
}
