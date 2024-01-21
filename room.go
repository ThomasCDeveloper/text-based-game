package main

import (
	"fmt"
	"strings"
)

type Room interface {
	getDescription() string

	Goto(dist string, current *string)
	Use(obj string, bag *[]string)
	Take(obj string, bag *[]string)
}

type Basicroom struct {
	description string
	exits       map[string]string

	objects        map[string]bool
	objDescription map[string]string
}

func (r Basicroom) getDescription() string {
	description := r.description

	for objName, isInRoom := range r.objects {
		if isInRoom {
			description += " " + r.objDescription[objName]
		}
	}
	description = replaceInString(description, "#B", "\033[4m", "\033[0m")

	return description
}

func (r Basicroom) Goto(dist string, current *string) {
	if val, ok := r.exits[dist]; ok {
		*current = val
	} else {
		fmt.Println("\"" + dist + "\" is not a valid destination.")
	}
}

func (r Basicroom) Use(obj string, bag *[]string) {
	fmt.Println(obj + " is no use here.")
}

func (r Basicroom) Take(obj string, bag *[]string) {
	if val, ok := r.objects[obj]; val && ok {
		*bag = append(*bag, obj)
		r.objects[obj] = false
	}
}

func replaceInString(in string, del string, rep1 string, rep2 string) string {
	parts := strings.Split(in, del)
	for i := 0; i < len(parts)-1; i += 2 {
		parts[i] = parts[i] + rep1
	}
	for i := 2; i < len(parts); i += 2 {
		parts[i] = rep2 + parts[i]
	}
	return strings.Join(parts, "")
}

type StartRoom struct {
	exits map[string]string

	isDoorOpen *bool

	objects        map[string]bool
	objDescription map[string]string
}

func (r StartRoom) getDescription() string {
	description := "you find yourself in a dark room. There seems to be a #Bdoor#B."

	for objName, isInRoom := range r.objects {
		if isInRoom {
			description += " " + r.objDescription[objName]
		}
	}
	description = replaceInString(description, "#B", "\033[4m", "\033[0m")

	return description
}

func (r StartRoom) Goto(dist string, current *string) {
	if dist == "door" {
		if *r.isDoorOpen {
			*current = r.exits[dist]
		} else {
			fmt.Println("The door is locked.")
		}
	} else {
		fmt.Println("\"" + dist + "\" is not a valid destination.")
	}
}

func (r StartRoom) Use(obj string, bag *[]string) {
	if obj == "key" {
		*r.isDoorOpen = !*r.isDoorOpen
		fmt.Println("You used the key!")
	} else {
		fmt.Println(obj + " is no use here.")
	}
}

func (r StartRoom) Take(obj string, bag *[]string) {
	if val, ok := r.objects[obj]; val && ok {
		*bag = append(*bag, obj)
		r.objects[obj] = false
		fmt.Println("You took the key.")
	}
}
