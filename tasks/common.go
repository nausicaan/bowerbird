package tasks

import (
	"os"
	"os/exec"
	"strings"
)

const (
	upbranch, relbranch string = "update/DESSO-", "release/"
)

var (
	wordpress               WordPress
	events                  Events
	number, folder          []string
	plugin, ticket, release string
	// Edict holds the type of composer command
	Edict string
	// ArgLength measures the number of total arguments
	ArgLength = len(os.Args)
	// Flag holds the type argument
	Flag = verify()
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
	} else if Flag == "-p" && folder[0] == "bcgov-plugin/events-virtual" {
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
func push(where string) {
	if Flag == "-w" {
		exec.Command("git", "push").Run()
	} else if Flag == "-r" {
		exec.Command("git", "push", "--set-upstream", "origin", where+release).Run()
	} else {
		exec.Command("git", "push", "--set-upstream", "origin", where+ticket).Run()
	}
}
