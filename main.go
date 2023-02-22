package main

import (
	"fmt"

	t "github.com/nausicaan/bowerbird/tasks"
)

const (
	bv     string = "1.0.1"
	green  string = "\033[32m"
	yellow string = "\033[33m"
	zero   string = "Insufficient arguments supplied -"
)

// Launch the program and execute the selected program abilities
func main() {
	switch t.Flag {
	case "-v", "--version":
		fmt.Println(yellow+"Bowerbird", green+bv)
	case "-h", "--help":
		t.HelpMenu()
	case "-m", "--managed", "-r", "--release":
		wrtest(t.Flag)
	case "-p", "--premium":
		t.Flag = "-p"
		premtest()
	case "--zero":
		t.Errors("No flag detected -")
		t.HelpMenu()
	default:
		t.Errors("Bad flag detected -")
		t.HelpMenu()
	}
	t.Push()
}

// Determine which function to call based on the passed variable.
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
		t.Errors(zero)
		t.HelpMenu()
	}
}

// Call the Premium function if the required arguments are supplied
func premtest() {
	if t.ArgLength < 4 {
		t.Errors(zero)
		t.HelpMenu()
	} else if t.ArgLength > 4 {
		t.Errors("Too many arguments supplied -")
		t.HelpMenu()
	} else {
		t.Prepare()
		t.Premium()
	}
}
