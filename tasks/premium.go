package tasks

import (
	"encoding/json"
	"io"
	"os"
	"os/exec"
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

// Premium contains a sequential list of tasks to run to complete the program
func Premium() {
	jsonParse()
	assign(os.Args[2], os.Args[3])
	norm.Version, odd.Version = number[1], number[1]
	if strings.Contains(folder[1], "event") {
		if odd.Name+":"+odd.Version == plugin {
			execute()
		}
	} else if norm.Name+":"+norm.Version == plugin {
		execute()
	} else {
		Errors("Plugin name does not match composer.json entry - program halted")
	}
}

// A sequential list of tasks run to complete the program
func execute() {
	Prepare()
	checkout(upbranch)
	script()
	jsonWrite()
	commit()
	tags()
}

// Read the composer.json file and store the results in the WordPress structure
func jsonParse() {
	current, _ := os.Open("composer.json")
	byteValue, _ := io.ReadAll(current)
	defer current.Close()
	json.Unmarshal(byteValue, &norm)
	json.Unmarshal(byteValue, &odd)
}

// Run the update script
func script() {
	exec.Command("/bin/bash", "-c", "scripts/update.sh ~/Downloads/"+folder[1]+"/").Run()
}

// Convert the WordPress structure back into json and overwrite the composer.json file
func jsonWrite() {
	var updated []byte
	if strings.Contains(odd.Name, "event") {
		updated, _ = json.MarshalIndent(odd, "", "    ")
	} else {
		updated, _ = json.MarshalIndent(norm, "", "    ")
	}
	os.WriteFile("composer.json", updated, 0644)
}

// Tag the version so Satis can package it
func tags() {
	exec.Command("git", "tag", "v"+norm.Version).Run()
	exec.Command("git", "push", "origin", "--tags").Run()
}
