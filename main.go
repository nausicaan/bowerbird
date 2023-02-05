package main

import (
	"fmt"

	t "github.com/nausicaan/upinstall/tasks"
)

// Colour Palette
const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
)

var buildVersion, zero = "1.0.1", "Insufficient arguments supplied - program halted"

// Launch the program and execute the selected program abilities
func main() {
	switch t.Flag {
	case "-v", "--version":
		fmt.Println(yellow+"upinstall", green+buildVersion)
	case "-h", "--help":
		helpMenu()
	case "-w", "--wpackagist", "-r", "--release":
		testWR(t.Flag)
	case "-p", "--premium":
		testP()
	case "--zero":
		fmt.Println("No flag detected -", red+"program halted")
		fmt.Println()
	default:
		fmt.Println("Incorrect flag detected -", red+"program halted")
		fmt.Println()
	}
}

// Print the help information
func helpMenu() {
	fmt.Println(yellow, "\nUsage:", reset)
	fmt.Println("  [program] [flag] [vendor/plugin]:[version] [ticket#]")
	fmt.Println(yellow, "\nOptions:")
	fmt.Println(green, " -p, --premium", reset, "	Premium Plugin Update")
	fmt.Println(green, " -r, --release", reset, "	Production Release Plugin Updates")
	fmt.Println(green, " -w, --wpackagist", reset, "	WPackagist Plugin Updates")
	fmt.Println(green, " -v, --version", reset, "	Display App Version")
	fmt.Println(green, " -h, --help", reset, "		Help Information")
	fmt.Println(yellow, "\nExample:", reset)
	fmt.Println("  Navigate to the folder containing your composer.json file and run:")
	fmt.Println(green, "\n    ~/Documents/programs/upinstall -w wpackagist-plugin/mailpoet:5.5.2 762")
	fmt.Println()
}

// Determine which function to call based on the passed variable.
func testWR(flag string) {
	if t.ArgLength >= 4 {
		switch flag {
		case "-w", "--wpackagist":
			t.WPackagist()
		case "-r", "--release":
			t.Release()
		}
	} else {
		fmt.Println(zero)
	}
}

// Call the Premium function if the required arguments are supplied
func testP() {
	if t.ArgLength < 4 {
		fmt.Println(zero)
	} else if t.ArgLength > 4 {
		fmt.Println("Too many arguments supplied - program halted")
	} else {
		t.Premium()
	}
}
