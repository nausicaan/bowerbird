package tasks

import (
	"os"
	"os/exec"
)

const (
	upbranch, relbranch string = "update/DESSO-", "release/DESSO-"
)

var (
	plugin, ticket string
	// ArgLength measures the number of total arguments
	ArgLength = len(os.Args)
	// Flag holds the type argument
	Flag = verify()
)

func verify() string {
	var f string
	if ArgLength < 2 {
		f = "-zzz"
	} else {
		f = os.Args[1]
	}
	return f
}

// Switch to the desired branch, pull any changes, and run a composer update
func prepare() {
	var branch string
	if Flag == "-f" {
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
func push() {
	if Flag == "-f" {
		exec.Command("git", "push").Run()
	} else {
		exec.Command("git", "push", "--set-upstream", "origin", upbranch+ticket).Run()
	}
}
