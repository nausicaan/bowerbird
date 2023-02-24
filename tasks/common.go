package tasks

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	reset     string = "\033[0m"
	green     string = "\033[32m"
	yellow    string = "\033[33m"
	red       string = "\033[41m"
	relbranch string = "release/"
	upbranch  string = "update/DESSO-"
	halt      string = "program halted "
	counter   string = "/Users/byron/Documents/programs/count.txt"
)

var (
	norm Satis
	odd  Event
	// Flag holds the type argument
	Flag   = verify()
	reader = bufio.NewReader(os.Stdin)
	// ArgLength measures the number of total arguments
	ArgLength               = len(os.Args)
	number, folder          []string
	plugin, ticket, release string
)

// HelpMenu prints the help information
func HelpMenu() {
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

// Prepare switches to the desired branch, pull any changes, and run a composer update
func Prepare() {
	var branch string
	if Flag == "-m" {
		branch = "development"
	} else if Flag == "-p" && folder[1] == "events-virtual" {
		branch = "main"
	} else {
		branch = "master"
	}
	console("git", "switch", branch)
	console("git", "pull")
	// exec.Command("git", "switch", branch).Run()
	// exec.Command("git", "pull").Run()
}

// Takes a string prompt and asks the user for input.
func prompt(prompt string) string {
	fmt.Print(prompt)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)
	return answer
}

// Runs standard terminal commands and displays the output
func console(name string, task ...string) {
	path, err := exec.LookPath(name)
	osCmd := exec.Command(path, task...)
	osCmd.Stdout = os.Stdout
	osCmd.Stderr = os.Stderr
	err = osCmd.Run()
	problem(err)
}

// Check for errors, halt the program if found, and log the result
func problem(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Test for the minimum amount of arguments
func verify() string {
	var f string
	if ArgLength < 2 {
		f = "--zero"
	} else {
		f = os.Args[1]
	}
	return f
}

// Check to see if the current release branch already exists locally
func exists(prefix string) bool {
	found := false
	b, _ := exec.Command("git", "branch").Output()
	if strings.Contains(string(b), prefix+release) {
		found = true
	}
	return found
}

// Deceide whether an update or release branch is needed, and make it so
func checkout(prefix string) {
	if Flag == "-r" {
		if exists(prefix) {
			console("git", "switch", prefix+release)
			// exec.Command("git", "switch", prefix+release).Run()
		} else {
			console("git", "checkout", "-b", prefix+release)
			// exec.Command("git", "checkout", "-b", prefix+release).Run()
		}
	} else {
		console("git", "checkout", "-b", prefix+ticket)
		// exec.Command("git", "checkout", "-b", prefix+ticket).Run()
	}
}

// Add and commit the update
func commit() {
	console("git", "add", ".")
	console("git", "commit", "-m", plugin+" (DESSO-"+ticket+")")
	// exec.Command("git", "add", ".").Run()
	// exec.Command("git", "commit", "-m", plugin+" (DESSO-"+ticket+")").Run()
}

// Errors prints a clolourized error message
func Errors(message string) {
	fmt.Println(red, message, halt)
	fmt.Println(reset)
}

// Push to the git repository
func Push() {
	switch Flag {
	case "-r":
		console("git", "push", "--set-upstream", "origin", relbranch+release)
		// exec.Command("git", "push", "--set-upstream", "origin", relbranch+release).Run()
	case "-p":
		console("git", "push", "--set-upstream", "origin", upbranch+ticket)
		// exec.Command("git", "push", "--set-upstream", "origin", upbranch+ticket).Run()
	default:
		console("git", "push")
		// exec.Command("git", "push").Run()
	}
}

// Tracking provides informational messages about the programs progress
func Tracking(message string) {
	fmt.Println(yellow)
	fmt.Println("**", reset, message, yellow, "**", reset)
}
