package tasks

import (
	"os"
	"os/exec"
)

var (
	plugin, ticket string
	flag           = os.Args[1]
)

// Switch to the desired branch, pull any changes, and run a composer update
func prepare() {
	var branch string
	if flag == "-f" {
		branch = "development"
	} else {
		branch = "master"
	}
	exec.Command("git", "switch", branch).Run()
	exec.Command("git", "pull").Run()
}

// Add and commit the update
func commit() {
	exec.Command("git", "add", ".").Run()
	exec.Command("git", "commit", "-m", plugin+" (DESSO-"+ticket+")").Run()
}

// Push to the git repository
func push() {
	if flag == "-f" {
		exec.Command("git", "push").Run()
	} else {
		exec.Command("git", "push", "--set-upstream", "origin", "update/DESSO-"+ticket).Run()
	}
}
