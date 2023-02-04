package main

import (
	"fmt"

	t "github.com/nausicaan/upinstall/tasks"
)

var buildVersion, zero = "1.0.1", "Insufficient arguments supplied - program halted"

// Launch the program and execute the selected program abilities
func main() {
	switch t.Flag {
	case "-v":
		fmt.Println("Upcheck", buildVersion)
	case "--version":
		fmt.Println("Upcheck", buildVersion)
	case "--zero":
		fmt.Println("No flag detected - program halted")
	case "-h":
		helpMenu()
	case "--help":
		helpMenu()
	case "-w":
		testWR(t.Flag)
	case "-r":
		testWR(t.Flag)
	case "-p":
		testP()
	default:
		fmt.Println("Incorrect flag detected - program halted")
	}
}

func helpMenu() {
	fmt.Println("\nUsage: [program_name] [flag] [full_plugin_name]:[update_version] [jira_ticket_number]")
	fmt.Println("\n  -p	Premium Plugin Update")
	fmt.Println("  -r	Production Release Plugin Updates")
	fmt.Println("  -w	WPackagist Plugin Updates")
	fmt.Println("  -v	--version	Display App Version")
	fmt.Println("  -h	--help		Help Information")
	fmt.Println()
}

func testWR(flag string) {
	if t.ArgLength >= 4 {
		switch flag {
		case "-w":
			t.WPackagist()
		case "-r":
			t.Release()
		}
	} else {
		fmt.Println(zero)
	}
}

func testP() {
	if t.ArgLength < 4 {
		fmt.Println(zero)
	} else if t.ArgLength > 4 {
		fmt.Println("Too many arguments supplied - program halted")
	} else {
		t.Premium()
	}
}
