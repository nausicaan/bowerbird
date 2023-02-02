package tasks

import (
	"os"
	"os/exec"
)

// Free contains a sequential list of tasks run to complete the program
func Free() {
	prepare()
	update()
	sift()
	push()
}

// Run the general composer update command to check for lock file updates
func update() {
	exec.Command("composer", "update").Run()
}

// Run the composer require command
func require() {
	if Flag == "-r" {
		exec.Command("COMPOSER=composer-prod.json composer", "require", plugin).Run()
	} else {
		exec.Command("composer", "require", plugin).Run()
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
