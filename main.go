package main

import "fmt"

func main() {
	users := make(chan []string, 1)
	singers := make(chan []string, 1)

	go getUsers(users)
	go getSingers(singers)

	channelCount := 2
	for i := 0; i < channelCount; i++ {
		select {
		case users := <-users:
			fmt.Println("usuarios:", users)
		case singers := <-singers:
			fmt.Println("cantantes:", singers)
		}
	}
}

func getUsers(users chan<- []string) {
	users <- []string{"edarcode", "lore"}
}

func getSingers(singers chan<- []string) {
	singers <- []string{"shakira", "bad bunny"}
}
