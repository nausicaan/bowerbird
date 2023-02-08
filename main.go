package main

import (
	"fmt"

	t "github.com/nausicaan/bowerbird/tasks"
)

const (
	bv     string = "1.0.1"
	reset  string = "\033[0m"
	red    string = "\033[31m"
	green  string = "\033[32m"
	yellow string = "\033[33m"
	halt   string = red + "program halted"
	zero   string = "Insufficient arguments supplied - " + halt
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
		fmt.Println("No flag detected -", halt)
		fmt.Println()
	default:
		fmt.Println("Incorrect flag detected -", halt)
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
		fmt.Println(zero)
		fmt.Println()
	}
}

// Call the Premium function if the required arguments are supplied
func testP() {
	if t.ArgLength < 4 {
		fmt.Println(zero)
	} else if t.ArgLength > 4 {
		fmt.Println("Too many arguments supplied -", halt)
		fmt.Println()
	} else {
		t.Premium()
	}
}
