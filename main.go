package main

import (
	"fmt"

	t "github.com/nausicaan/wp-updater/tasks"
)

var BuildVersion = "1.0.0"

// Launch the program and execute the selected program abilities
func main() {
	switch t.Flag {
	case "-v":
		fmt.Println("Updater", BuildVersion)
	case "-zzz":
		fmt.Println("No flag detected - program halted")
	case "-h":
		fmt.Println()
		fmt.Println("Usage: [program_name] [flag] [full_plugin_name]:[update_version] [jira_ticket_number]")
		fmt.Println("\n  -f	Free Plugin Update")
		fmt.Println("  -h	Help Information")
		fmt.Println("  -p	Premium Plugin Update")
		fmt.Println("  -v	Display App Version")
		fmt.Println()
	case "-f":
		if t.ArgLength >= 4 {
			t.Free()
		} else {
			fmt.Println("Insufficient arguments supplied - program halted")
		}
	case "-p":
		if t.ArgLength < 4 {
			fmt.Println("Insufficient arguments supplied - program halted")
		} else if t.ArgLength > 4 {
			fmt.Println("Too many arguments supplied - program halted")
		} else {
			t.Paid()
		}
	default:
		fmt.Println("Incorrect flag detected - program halted")
	}
}
