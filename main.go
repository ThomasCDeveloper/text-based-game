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

	rooms["room"] = Room{"You find yourself in a dark room... You see a \033[4mdoor\033[0m.", map[string]string{"door": "end"}}
	rooms["end"] = Room{"Nothing to see here... You can go back where you woke up in the \033[4mroom\033[0m.", map[string]string{"room": "room"}}

	position := "room"

	for {
		rooms[position].show()

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
				fmt.Println("Help for " + words[1])
				switch words[1] {
				case "goto":
					fmt.Println("Valid destinations are often underlined.")
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
	}
}
