package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

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

// Confirm the current working directory is correct
func changedir() {
	os.Chdir(bitbucket + "blog_gov_bc_ca")
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

// Add and Commit the queued update
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
