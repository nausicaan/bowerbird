package tasks

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// WPackagist contains a sequential list of tasks run to complete the program
func WPackagist() {
	prepare()
	update()
	sift("--quiet")
	push()
}

// Release adds the previously tested plugins to the composer-prod.json file
func Release() {
	byterel, _ := exec.Command("tail", "-n1", counter).Output()
	intrel, _ := strconv.Atoi(string(byterel))
	intrel++
	release = fmt.Sprint(intrel)
	f, _ := os.OpenFile(counter, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.Write([]byte("\n" + release))
	defer f.Close()
	prepare()
	checkout(relbranch)
	sift("--no-install")
	push()
}

// Errors prints a clolourized error message
func Errors(message string) {
	fmt.Println(red, message, halt)
	fmt.Println()
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
