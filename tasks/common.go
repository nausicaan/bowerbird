package tasks

import (
	"os"
	"os/exec"
	"strings"
)

const (
	upbranch, relbranch, counter string = "update/DESSO-", "release/", "/Users/byron/Documents/programs/count.txt"
)

var (
	// Flag holds the type argument
	Flag = verify()
	// Edict holds the type of composer command
	Edict string
	// ArgLength measures the number of total arguments
	ArgLength = len(os.Args)

	plugin, ticket, release string
	number, folder          []string
)

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

// Split the supplied arguments and assign them to variables
func assign(p, t string) {
	plugin, ticket = p, t
	number = strings.Split(plugin, ":")
	folder = strings.Split(number[0], "/")
}

// Switch to the desired branch, pull any changes, and run a composer update
func prepare() {
	var branch string
	if Flag == "-w" {
		branch = "development"
	} else if Flag == "-p" && folder[1] == "events-virtual" {
		branch = "main"
	} else {
		branch = "master"
	}
	exec.Command("git", "switch", branch).Run()
	exec.Command("git", "pull").Run()
}

// Create an update branch to work from
func checkout(prefix string) {
	if Flag == "-r" {
		exec.Command("git", "checkout", "-b", prefix+release).Run()
	} else {
		exec.Command("git", "checkout", "-b", prefix+ticket).Run()
	}
}

// Add and commit the update
func commit() {
	exec.Command("git", "add", ".").Run()
	exec.Command("git", "commit", "-m", plugin+" (DESSO-"+ticket+")").Run()
}

// Push to the git repository
func push() {
	switch Flag {
	case "-r":
		exec.Command("git", "push", "--set-upstream", "origin", relbranch+release).Run()
	case "-p":
		exec.Command("git", "push", "--set-upstream", "origin", upbranch+ticket).Run()
	default:
		exec.Command("git", "push").Run()
	}
}
