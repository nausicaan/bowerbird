package tasks

import (
	"os"
)

// Managed contains a sequential list of tasks run to complete the program
func Managed() {
	update()
	sift("--quiet")
	Push()
}

// Release adds the previously tested plugins to the composer-prod.json file
func Release() {
	release = prompt("Enter the current release number: ")
	checkout(relbranch)
	sift("--no-install")
	Push()
}

// Run the general composer update command to check for lock file updates
func update() {
	console("composer", "update")
}

// Run the appropriate composer require command
func require(option string) {
	if Flag == "-r" {
		console("env", "COMPOSER=composer-prod.json", "composer", "require", plugin, option)
	} else {
		console("composer", "require", plugin, option)
	}
}

// Iterate through the Args array and assign plugin and ticket values
func sift(option string) {
	for i := 2; i < len(os.Args); i++ {
		plugin = os.Args[i]
		i++
		ticket = os.Args[i]
		require(option)
		commit()
	}
}
