package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var rooms = map[string]Room{}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// rooms["room"] = Room{"You find yourself in a dark room... You see a \033[4mdoor\033[0m.", map[string]string{"door": "end"}}
	rooms["room"] = Room{"You find yourself in a dark room... You see a #Bdoor#B.", map[string]string{"door": "end"}, []string{"key"}, map[string]string{"key": "There is a #Bkey#B on the ground."}}
	rooms["end"] = Room{"Nothing to see here... You can go back where you woke up in the \033[4mroom\033[0m.", map[string]string{"room": "room"}, []string{}, map[string]string{}}

	position := "room"
	bag := []string{"key"}

	for {
		rooms[position].show(bag)

		fmt.Printf("> ")
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1]
		words := strings.Split(input, " ")

		if input == "exit" {
			break
		}

		if words[0] == "help" {
			if len(words) == 1 {
				fmt.Println("  help: prints this message\n  exit: exits the game\n  goto <place>: attemps to go to <place>")
			} else {
				switch words[1] {
				case "goto":
					fmt.Println("  Valid destinations are often underlined.")
				default:
					fmt.Println(words[1] + " is not a valid command. Type \"help\" to see the valid commands.")
				}
			}
		}

		if words[0] == "goto" && len(words) > 1 {
			if exit, ok := rooms[position].exits[words[1]]; ok {
				position = exit
			}
		}
		fmt.Println()
	}
}
