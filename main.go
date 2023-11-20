package main

import (
	"fmt"
	"os"
	"strings"
)

// Launch the program and execute the appropriate code
func main() {
	switch flag {
	case "-v", "--version":
		version()
	case "-h", "--help":
		about()
	case "-m", "--managed", "-r", "--release":
		os.Chdir(hmdr + "/Documents/bitbucket/blog_gov_bc_ca")
		doublecheck()
		delegate(flag)
	case "-p", "--premium":
		flag = "-p"
		files := ls(hmdr + "/Documents/interactions/premium/")
		for _, file := range files {
			if file == ".DS_Store" {
				cleanup(file)
			} else {
				os.Chdir(hmdr + "/Documents/bitbucket/" + strings.TrimSuffix(file, ".txt"))
				discovery(hmdr + "/Documents/interactions/premium/" + file)
				premium()
			}
		}
	case "--zero":
		about()
		alert("No flag detected -")
	default:
		about()
		alert("Bad flag detected -")
	}
}

// Determine which function to call based on the passed variable
func delegate(flag string) {
	prepare()
	switch flag {
	case "-m", "--managed":
		flag = "-m"
		discovery(hmdr + "/Documents/interactions/managed.txt")
		managed()
	case "-r", "--release":
		flag = "-r"
		discovery(hmdr + "/Documents/interactions/release.txt")
		released()
	}
}

// Print a colourized error message
func alert(message string) {
	fmt.Println(red, message, halt, reset)
	os.Exit(0)
}

// Provide and highlight informational messages
func tracking(message string) {
	fmt.Println(yellow)
	fmt.Println("**", reset, message, yellow, "**", reset)
}

// Print the build version of the program
func version() {
	fmt.Println(yellow+"Bowerbird", green+bv, reset)
}

// Print help information for using the program
func about() {
	fmt.Println(yellow, "\nUsage:", reset)
	fmt.Println("  [program] [flag]")
	fmt.Println(yellow, "\nOptions:")
	fmt.Println(green, " -p, --premium", reset, "	Premium Plugin Repository Update")
	fmt.Println(green, " -r, --release", reset, "	Production Release Plugin Update")
	fmt.Println(green, " -m, --managed", reset, "	Satis & WPackagist Plugin Update")
	fmt.Println(green, " -v, --version", reset, "	Display App Version")
	fmt.Println(green, " -h, --help", reset, "		Help Information")
	fmt.Println(yellow, "\nExample:", reset)
	fmt.Println("  Against your composer.json file, run:")
	fmt.Println(green, "   bowerbird -m")
	fmt.Println(yellow, "\nHelp:", reset)
	fmt.Println("  For more information go to:")
	fmt.Println(green, "   https://github.com/nausicaan/bowerbird.git")
	fmt.Println(reset)
}
