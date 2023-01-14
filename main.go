package main

import (
	"fmt"
	"os"

	t "github.com/nausicaan/free/tasks"
)

// Launch the program and execute the selected program abilities
func main() {
	argLength := len(os.Args)
	if argLength >= 4 {
		choice := os.Args[1]
		switch choice {
		case "-f":
			t.Free()
		case "-p":
			t.Paid()
		default:
			fmt.Println("Incorrect flag detected - program halted")
		}
	} else {
		fmt.Println("Insufficient arguments supplied - program halted")
	}
}
