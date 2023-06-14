package tasks

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)

// Satis structure to hold the contents of the composer.json file
type Satis struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`
	Require struct {
		ComposerInstallers string `json:"composer/installers"`
	} `json:"require"`
}

// Event structure to hold the contents of the composer.json file
type Event struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`
	Require struct {
		ComposerInstallers                string `json:"composer/installers"`
		WpackagistPluginTheEventsCalendar string `json:"wpackagist-plugin/the-events-calendar"`
	} `json:"require"`
}

// A sequential list of tasks run to complete the program
func quarterback() {
	Prepare()
	checkout(upbranch)
	tracking("Update Script")
	script()
	correct()
	commit()
	tracking("Tagging to Satis")
	tags()
	tracking("Git Push")
	push()
}

// Premium directs the preliminary actions to determine if the program can continue
func Premium() {
	learn()
	assign(os.Args[2], os.Args[3])
	norm.Version, odd.Version = number[1], number[1]
	if strings.Contains(folder[1], "event") {
		if odd.Name+":"+odd.Version == plugin {
			quarterback()
		}
	} else if norm.Name+":"+norm.Version == plugin {
		quarterback()
	} else {
		Alert("Plugin name does not match composer.json entry - program halted")
	}
}

// Read the composer.json file and store the results in a structure
func learn() {
	current, _ := os.Open("composer.json")
	byteValue, _ := io.ReadAll(current)
	err := json.Unmarshal(byteValue, &norm)
	inspect(err)
	err = json.Unmarshal(byteValue, &odd)
	inspect(err)
	err = current.Close()
	inspect(err)
}

// Split the supplied arguments and assign them to variables
func assign(p, t string) {
	plugin, ticket = p, t
	number = strings.Split(plugin, ":")
	folder = strings.Split(number[0], "/")
}

// Run the update script on downloaded content
func script() {
	execute("sh", "-c", "scripts/update.sh ~/Downloads/"+folder[1]+"/")
}

// Convert the structure back into json and overwrite the composer.json file
func correct() {
	var updated []byte
	if strings.Contains(odd.Name, "event") {
		updated, _ = json.MarshalIndent(odd, "", "    ")
	} else {
		updated, _ = json.MarshalIndent(norm, "", "    ")
	}
	err := os.WriteFile("composer.json", updated, 0644)
	inspect(err)
}

// Tag the version so Satis can package it
func tags() {
	execute("git", "tag", "v"+norm.Version)
	execute("git", "push", "origin", "--tags")
}
