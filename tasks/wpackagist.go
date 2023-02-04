package tasks

import (
	"os"
	"os/exec"
)

// WPackagist contains a sequential list of tasks run to complete the program
func WPackagist() {
	prepare()
	update()
	sift()
	push("")
}

// Release adds the previously tested plugins to the composer-prod.json file
func Release() {
	assign()
	prepare()
	checkout(relbranch)
	sift()
	push(relbranch)
}

// Run the general composer update command to check for lock file updates
func update() {
	exec.Command("composer", "update").Run()
}

// Run the composer require command
func require() {
	runcmd.Run()
	// if Flag == "-r" {
	// 	exec.Command(prodprefix, "composer", "require", plugin).Run()
	// } else {
	// 	exec.Command("composer", "require", plugin).Run()
	// }
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
