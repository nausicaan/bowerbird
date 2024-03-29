package main

// A sequential list of tasks run to complete the program
func wpackagist() {
	tracking("Composer Update")
	execute("composer", "update", "--no-install")
	tracking("Plugin Update")
	sift()
	tracking("Git Push")
	push()
}

// Add the previously tested plugins to the composer-prod.json file
func released() {
	release = solicit("Enter the current release number: ")
	checkout(relbranch)
	wpackagist()
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
			execute("composer", "require", plugin, "-W", "--no-install")
		} else {
			execute("composer", "require", plugin, "--no-install")
		}
	}
}

// Iterate through the Args array and assign plugin and ticket values
func sift() {
	for i := 2; i < inputs; i++ {
		plugin = passed[i]
		i++
		ticket = passed[i]
		require()
		commit()
	}
}
