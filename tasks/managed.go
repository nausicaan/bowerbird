package tasks

import (
	"os"
)

// Managed contains a sequential list of tasks run to complete the program
func Managed() {
	Tracking("Running composer update")
	update()
	Tracking("Creating commits")
	sift()
	push()
}

// Release adds the previously tested plugins to the composer-prod.json file
func Release() {
	release = prompt("Enter the current release number: ")
	checkout(relbranch)
	sift()
	push()
}

// Run the general composer update command to check for lock file updates
func update() {
	console("composer", "update")
}

// Run the appropriate composer require command
func require() {
	if Flag == "-r" {
		console("env", "COMPOSER=composer-prod.json", "composer", "require", plugin, "--no-install")
	} else {
		console("composer", "require", plugin)
	}
}

// Iterate through the Args array and assign plugin and ticket values
func sift() {
	for i := 2; i < len(os.Args); i++ {
		plugin = os.Args[i]
		i++
		ticket = os.Args[i]
		require()
		commit()
	}
}
