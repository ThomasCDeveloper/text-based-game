package main

import (
	"fmt"
	"time"
)

type Dialog struct {
	millis  []int
	phrases []string
	pauses  []int
}

func (d Dialog) print() {
	for i, phrase := range d.phrases {
		for _, char := range phrase {
			fmt.Printf(string(char))
			time.Sleep(time.Millisecond * time.Duration(d.millis[i]))
		}
		time.Sleep(time.Millisecond * time.Duration(d.pauses[i]))
	}
}
