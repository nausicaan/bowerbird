package main

// A sequential list of tasks run to complete the program
func wpackagist(content []string) {
	tracking("Composer Update")
	execute("composer", "update")
	tracking("Plugin Update")
	sift(content)
	tracking("Git Push")
	push()
}

// Add the previously tested plugins to the composer-prod.json file
func releases(content []string) {
	release = solicit("Enter the current release number: ")
	checkout(relbranch)
	sift(content)
}

// Add the Developer tested plugins to the composer-prod.json file
func inhouse(content []string) {
	sift(content)
	tracking("Git Push")
	push()
}

// Iterate through the Issues JSON object and assign plugin and ticket values
func sift(content []string) {
	for i := 0; i < len(content); i++ {
		plugin = content[i]
		i++
		ticket = content[i]
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
