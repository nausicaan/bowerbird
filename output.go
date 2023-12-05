package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	hmdr, _   = os.UserHomeDir()
	reader    = bufio.NewReader(os.Stdin)
	common    = hmdr + "/Documents/common/"
	bitbucket = hmdr + "/Documents/bitbucket/"
)

// Use JQL to query the Jira API
func apiget(query string) []byte {
	/* Temporary local fix until Jira API is accessible */
	result := read(common + query)
	// result := execute("-c", "curl", "-X", "GET", "-H", "Authorization: Bearer "+" vendor/" + jira.Token, "-H", "Content-Type: application/json", "vendor/" + jira.Base+query)
	return result
}

// Print a colourized error message
func alert(message string) {
	fmt.Printf("\n%s %s\n", message, halt)
	os.Exit(0)
}

// Provide and highlight informational messages
func tracking(message string) {
	fmt.Printf("** %s **\n", message)
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
