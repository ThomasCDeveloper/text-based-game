package main

import (
	"fmt"
	"slices"
	"strings"
)

type Room struct {
	description string
	exits       map[string]string

	objects        []string
	objDescription map[string]string
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

func (room Room) show(bag []string) {
	description := room.description

	for _, objName := range room.objects {
		if !slices.Contains(bag, objName) {
			description += " " + room.objDescription[objName]
		}
	}

	description = replaceInString(description, "#B", "\033[4m", "\033[0m")
	fmt.Println(description)
}
