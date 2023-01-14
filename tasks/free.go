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

func update() {
	exec.Command("composer", "update").Run()
}

// Run the composer require command
func require() {
	exec.Command("composer", "require", plugin).Run()
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
