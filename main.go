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
	zero   string = "Insufficient arguments supplied -"
)

// Launch the program and execute the appropriate code
func main() {
	switch t.Flag {
	case "-v", "--version":
		fmt.Println(yellow+"Bowerbird", green+bv)
		fmt.Println(reset)
	case "-h", "--help":
		about()
	case "-m", "--managed", "-r", "--release":
		wrtest(t.Flag)
	case "-p", "--premium":
		t.Flag = "-p"
		premtest()
	case "--zero":
		t.Alert("No flag detected -")
		about()
	default:
		t.Alert("Bad flag detected -")
		about()
	}
}

// Determine which function to call based on the passed variable
func wrtest(flag string) {
	if t.ArgLength >= 4 {
		t.Prepare()
		switch flag {
		case "-m", "--managed":
			t.Flag = "-m"
			t.Managed()
		case "-r", "--release":
			t.Flag = "-r"
			t.Release()
		}
	} else {
		t.Alert(zero)
		about()
	}
}

// Call the Premium function if the required arguments are supplied
func premtest() {
	if t.ArgLength < 4 {
		t.Alert(zero)
		about()
	} else if t.ArgLength > 4 {
		t.Alert("Too many arguments supplied -")
		about()
	} else {
		t.Premium()
	}
}

// about prints help information for using the program
func about() {
	fmt.Println(yellow, "\nUsage:", reset)
	fmt.Println("  [program] [flag] [vendor/plugin]:[version] [ticket#]")
	fmt.Println(yellow, "\nOptions:")
	fmt.Println(green, " -p, --premium", reset, "	Premium Plugin Repository Update")
	fmt.Println(green, " -r, --release", reset, "	Production Release Plugin Update")
	fmt.Println(green, " -m, --managed", reset, "	Satis & WPackagist Plugin Update")
	fmt.Println(green, " -v, --version", reset, "	Display App Version")
	fmt.Println(green, " -h, --help", reset, "		Help Information")
	fmt.Println(yellow, "\nExample:", reset)
	fmt.Println("  Against your composer.json file, run:")
	fmt.Println(green, "   bowerbird -m wpackagist-plugin/mailpoet:4.6.1 821")
	fmt.Println(yellow, "\nHelp:", reset)
	fmt.Println("  For more information go to:")
	fmt.Println(green, "   https://github.com/nausicaan/bowerbird.git")
	fmt.Println(reset)
}
