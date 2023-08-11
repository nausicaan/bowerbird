package tasks

import (
	"bufio"
	"fmt"
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
	tracking("Preparing Branch")
	var branch string
	if Flag == "-p" && folder[1] == "events-virtual" {
		branch = "main"
	} else if Flag == "-p" {
		branch = "master"
	} else {
		branch = "development"
	}
	execute("git", "switch", branch)
	execute("git", "pull")
}

// Get user input via screen prompt
func solicit(prompt string) string {
	fmt.Print(prompt)
	response, _ := reader.ReadString('\n')
	return strings.TrimSpace(response)
}

// Run standard terminal commands and display the output
func execute(task string, args ...string) {
	osCmd := exec.Command(task, args...)
	osCmd.Stdout = os.Stdout
	osCmd.Stderr = os.Stderr
	err := osCmd.Run()
	inspect(err)
}

// Check for errors, print the result if found
func inspect(err error) {
	if err != nil {
		fmt.Println(err)
		return
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

// Check for edge cases which require the -W flag
func edge() bool {
	found := false
	if strings.Contains(plugin, "roots/wordpress") {
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

// Provide and highlight informational messages
func tracking(message string) {
	fmt.Println(yellow)
	fmt.Println("**", reset, message, yellow, "**", reset)
}
