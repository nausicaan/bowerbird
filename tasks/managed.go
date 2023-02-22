package tasks

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// Managed contains a sequential list of tasks run to complete the program
func Managed() {
	update()
	sift("--quiet")
}

// Release adds the previously tested plugins to the composer-prod.json file
func Release() {
	getRel()
	writeRel()
	checkout(relbranch)
	sift("--no-install")
}

// Get the current release number from the count.txt file
func getRel() {
	byterel, _ := exec.Command("tail", "-n1", counter).Output()
	intrel, _ := strconv.Atoi(string(byterel))
	intrel++
	release = fmt.Sprint(intrel)
}

func writeRel() {
	f, _ := os.OpenFile(counter, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.Write([]byte("\n" + release))
	defer f.Close()
}

// Run the general composer update command to check for lock file updates
func update() {
	exec.Command("composer", "update").Run()
}

// Run the appropriate composer require command
func require(option string) {
	exec.Command("composer", "require", plugin, option).Run()
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
