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

// Prepare switches to the desired branch, and pulls any changes
func Prepare() {
	var branch string
	if Flag == "-m" {
		branch = "development"
	} else if Flag == "-p" && folder[1] == "events-virtual" {
		branch = "main"
	} else {
		branch = "master"
	}
	execute("git", "switch", branch)
	execute("git", "pull")
}

// Take a string prompt and ask the user for input
func prompt(prompt string) string {
	fmt.Print(prompt)
	answer, _ := reader.ReadString('\n')
	// answer = strings.TrimSpace(answer)
	return strings.TrimSpace(answer)
}

// Run standard terminal commands and display the output
func execute(name string, task ...string) {
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

// Test for the minimum number of arguments
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

// Decide whether an update or release branch is needed, and make it so
func checkout(prefix string) {
	if Flag == "-r" {
		if exists(prefix) {
			execute("git", "switch", prefix+release)
		} else {
			execute("git", "checkout", "-b", prefix+release)
		}
	} else {
		execute("git", "checkout", "-b", prefix+ticket)
	}
}

// Add and commit the update
func commit() {
	execute("git", "add", ".")
	execute("git", "commit", "-m", plugin+" (DESSO-"+ticket+")")
}

// Push modified content to the git repository
func push() {
	switch Flag {
	case "-r":
		execute("git", "push", "--set-upstream", "origin", relbranch+release)
	case "-p":
		execute("git", "push", "--set-upstream", "origin", upbranch+ticket)
	default:
		execute("git", "push")
	}
}

// Alert prints a colourized error message
func Alert(message string) {
	fmt.Println(red, message, halt)
	fmt.Println(reset)
}

// Provide informational messages about the programs progress
func tracking(message string) {
	fmt.Println(yellow)
	fmt.Println("**", reset, message, yellow, "**", reset)
}
