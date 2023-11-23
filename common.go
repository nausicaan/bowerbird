package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	bv        string = "2.1"
	upbranch  string = "update/"
	relbranch string = "release/"
	halt      string = "program halted "
	zero      string = "Insufficient arguments supplied -"
)

var (
	inputs    int
	event     Event
	satis     Satis
	plugin    string
	ticket    string
	release   string
	number    []string
	folder    []string
	updates   []string
	flag      = os.Args[1]
	hmdr, _   = os.UserHomeDir()
	reader    = bufio.NewReader(os.Stdin)
	common    = hmdr + "/Documents/common/"
	bitbucket = hmdr + "/Documents/bitbucket/"
)

func discovery(filepath string) {
	goals := read(filepath)
	updates = strings.Split(string(goals), " ")
	updates = updates[:len(updates)-1]
	inputs = len(updates)
}

// Read any file and return the contents as a byte variable
func read(file string) []byte {
	outcome, problem := os.ReadFile(file)
	inspect(problem)
	return outcome
}

// Open a file and append a string
func atf(name, content string) {
	// Open a file for appending, create it if it doesn't exist
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	inspect(err)
	defer file.Close()

	// Write the content to the file
	_, err = file.WriteString(content)
	inspect(err)
}

// Record a list of files in a folder
func ls(folder string) []string {
	var content []string
	dir := expose(folder)

	files, err := dir.ReadDir(0)
	inspect(err)

	for _, f := range files {
		content = append(content, f.Name())
	}
	return content
}

// Open a file for reading and return an os.File variable
func expose(file string) *os.File {
	outcome, err := os.Open(file)
	inspect(err)
	return outcome
}

// Confirm the current working directory is correct
func doublecheck() {
	var filePath string = "composer-prod.json"

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		alert("This is not the correct folder,")
	}
}

// Switch to the desired branch, and pull any changes
func prepare() {
	tracking("Preparing Branch")
	var branch string

	if flag == "-p" && folder[1] == "events-virtual" {
		branch = "main"
	} else if flag == "-p" {
		branch = "master"
	} else {
		branch = "development"
	}
	execute("git", "switch", branch)
	execute("git", "pull")
}

// Write a passed variable to a named file
func document(name string, d []byte) {
	inspect(os.WriteFile(name, d, 0644))
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

// Remove files or directories
func cleanup(cut ...string) {
	inspect(os.Remove(cut[0.]))
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

// Checkout an update or release branch
func checkout(prefix string) {
	if flag == "-r" {
		if exists(prefix) {
			execute("git", "switch", prefix+release)
		} else {
			execute("git", "checkout", "-b", prefix+release)
		}
	} else {
		execute("git", "checkout", "-b", prefix+ticket)
	}
}

// Add and Commit the update
func commit() {
	execute("git", "add", ".")
	execute("git", "commit", "-m", plugin+" ("+ticket+")")
}

// Push modified content to a git repository
func push() {
	switch flag {
	case "-r":
		execute("git", "push", "--set-upstream", "origin", relbranch+release)
	case "-p":
		execute("git", "push", "--set-upstream", "origin", upbranch+ticket)
	default:
		execute("git", "push")
	}
}
