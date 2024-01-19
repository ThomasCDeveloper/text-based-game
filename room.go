package main

import "fmt"

type Room struct {
	description string
	exits       map[string]string
}

func (room Room) show() {
	fmt.Println(room.description)
}
