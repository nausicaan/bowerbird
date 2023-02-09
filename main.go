package main

import (
	"fmt"

	t "github.com/nausicaan/bowerbird/tasks"
)

const (
	bv     string = "1.0.1"
	reset  string = "\033[0m"
	green  string = "\033[32m"
	yellow string = "\033[33m"
	red    string = "\033[41m"
	halt   string = "program halted "
	zero   string = "Insufficient arguments supplied -"
)

// Launch the program and execute the selected program abilities
func main() {
	switch t.Flag {
	case "-v", "--version":
		fmt.Println(yellow+"Bowerbird", green+bv)
	case "-h", "--help":
		helpMenu()
	case "-w", "--wpackagist", "-r", "--release":
		testWR(t.Flag)
	case "-p", "--premium":
		t.Flag = "-p"
		testP()
	case "--zero":
		errors(" No flag detected -")
	default:
		errors(" Incorrect flag detected -")
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
	fmt.Println("  In the folder containing your composer.json file, run:")
	fmt.Println(green, "\n    bowerbird -w wpackagist-plugin/mailpoet:4.6.1 821")
	fmt.Println(yellow, "\nHelp:", reset)
	fmt.Println("  For more information go to:")
	fmt.Println(green, "\n    https://github.com/nausicaan/bowerbird.git")
	fmt.Println()
}

// Determine which function to call based on the passed variable.
func testWR(flag string) {
	if t.ArgLength >= 4 {
		switch flag {
		case "-w", "--wpackagist":
			t.Flag = "-w"
			t.WPackagist()
		case "-r", "--release":
			t.Flag = "-r"
			t.Release()
		}
	} else {
		errors(zero)
	}
}

// Call the Premium function if the required arguments are supplied
func testP() {
	if t.ArgLength < 4 {
		errors(zero)
	} else if t.ArgLength > 4 {
		errors("Too many arguments supplied -")
	} else {
		t.Premium()
	}
}

func errors(message string) {
	fmt.Println(red, message, halt)
	fmt.Println()
}
