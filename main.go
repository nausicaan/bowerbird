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
	case "-w", "--wpackagist", "-r", "--release", "-ap", "approved":
		os.Chdir(bitbucket + "blog_gov_bc_ca")
		doublecheck()
		delegate(flag)
	case "-p", "--premium":
		flag = "-p"
		files := ls(common + "premium/")
		for _, file := range files {
			if file == ".DS_Store" {
				cleanup(file)
			} else {
				os.Chdir(bitbucket + strings.TrimSuffix(file, ".txt"))
				discovery(common + "premium/" + file)
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
	case "-w", "--wpackagist":
		discovery(common + "operational/wpackagist.txt")
		wpackagist()
	case "-r", "--release":
		flag = "-r"
		discovery(common + "operational/release.txt")
		released()
	case "-ap", "approved":
		discovery(common + "operational/approved.txt")
		wpackagist()
	}
}

// Print a colourized error message
func alert(message string) {
	fmt.Println(message, halt)
	os.Exit(0)
}

// Provide and highlight informational messages
func tracking(message string) {
	fmt.Println("**", message, "**")
}

// Print the build version of the program
func version() {
	fmt.Println("Bowerbird", bv)
}

// Print help information for using the program
func about() {
	fmt.Println("\nUsage:")
	fmt.Println("  [program] [flag]")
	fmt.Println("\nOptions:")
	fmt.Println("  -p, --premium", "	Premium Plugin Repository Update")
	fmt.Println("  -r, --release", "	Production Release Plugin Update")
	fmt.Println("  -w, --wpackagist", "	Satis & WPackagist Plugin Update")
	fmt.Println("  -v, --version", "	Display App Version")
	fmt.Println("  -h, --help", "		Help Information")
	fmt.Println("\nExample:")
	fmt.Println("  bowerbird -w")
	fmt.Println("\nHelp:")
	fmt.Println("  For more information go to:")
	fmt.Println("    https://github.com/nausicaan/bowerbird.git")
	fmt.Println()
}
