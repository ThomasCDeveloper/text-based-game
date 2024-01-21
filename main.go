package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var rooms = map[string]Room{}

var testDialog = Dialog{[]int{80, 80, 130, 90, 140}, []string{"Aaaaaaaarghh!! ", "Hey man! ", "You woke me up!! ", "You ask me who I am ?! ", "Boy oh boy, I live here!\n"}, []int{250, 450, 250, 300, 0}}

func main() {
	reader := bufio.NewReader(os.Stdin)
	testDialog.print()

	// rooms["room"] = Room{"You find yourself in a dark room... You see a \033[4mdoor\033[0m.", map[string]string{"door": "end"}}
	rooms = initRooms()

	position := "room"
	lastposition := "never"
	bag := []string{}

	for {
		if lastposition != position {
			fmt.Println()
			fmt.Println(rooms[position].getDescription())
			lastposition = position
		}

		fmt.Printf(" > ")
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1]
		words := strings.Split(input, " ")

		switch words[0] {
		case "exit", ":wq", "wq":
			goto end
		case "bag":
			if len(bag) > 0 {
				fmt.Println("In your bag, you have:\n -" + strings.Join(bag, "\n -"))
			} else {
				fmt.Println("You have nothing in your bag.")
			}
		case "help":
			if len(words) == 1 {
				fmt.Println("   help: prints this message\n   exit: exits the game\n   goto <place>: attemps to go to <place>")
			} else {
				switch words[1] {
				case "goto", "go":
					fmt.Println("  Valid destinations are often underlined.")
				case "help":
					fmt.Println("Wdym brotha?")
				default:
					fmt.Println(words[1] + " is not a valid command. Type \"help\" to see the valid commands.")
				}
			}
		case "goto", "go":
			if len(words) > 1 {
				rooms[position].Goto(words[1], &position)
			}
		case "use":
			if len(words) > 1 {
				rooms[position].Use(words[1], &bag)
			}
		case "take", "get":
			if len(words) > 1 {
				rooms[position].Take(words[1], &bag)
			}
		default:
			fmt.Println(words[0] + " is not a valid command.")
		}

	}
end:
}
