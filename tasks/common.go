package tasks

import (
	"os"
	"os/exec"
	"strings"
)

const (
	upbranch, relbranch string = "update/DESSO-", "release/DESSO-"
)

var (
	wordpress      WordPress
	number, folder []string
	plugin, ticket string
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
func assign() {
	plugin, ticket = os.Args[2], os.Args[3]
	number = strings.Split(plugin, ":")
	folder = strings.Split(number[0], "/")
}

// Switch to the desired branch, pull any changes, and run a composer update
func prepare() {
	// runcmd = choose()
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
	exec.Command("git", "checkout", "-b", prefix+ticket).Run()
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
	} else {
		exec.Command("git", "push", "--set-upstream", "origin", where+ticket).Run()
	}
}
