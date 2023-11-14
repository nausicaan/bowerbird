package main

// A sequential list of tasks run to complete the program
func managed() {
	tracking("Composer Update")
	execute("composer", "update")
	tracking("Plugin Update")
	sift()
	tracking("Git Push")
	push()
}

// Add the previously tested plugins to the composer-prod.json file
func released() {
	release = solicit("Enter the current release number: ")
	checkout(relbranch)
	managed()
}

// Iterate through the updates array and assign plugin and ticket values
func sift() {
	for i := 0; i < inputs; i++ {
		plugin = updates[i]
		i++
		ticket = updates[i]
		require()
		commit()
	}
}

// Run the appropriate composer require command based on the flag value
func require() {
	if flag == "-r" {
		if edge() {
			execute("env", "COMPOSER=composer-prod.json", "composer", "require", plugin, "-W", "--no-install")
		} else {
			execute("env", "COMPOSER=composer-prod.json", "composer", "require", plugin, "--no-install")
		}
	} else {
		if edge() {
			execute("composer", "require", plugin, "-W")
		} else {
			execute("composer", "require", plugin)
		}
	}
}
