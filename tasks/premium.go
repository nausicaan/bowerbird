package tasks

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// WordPress structure to hold the contents of the composer.json file
type WordPress struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`
	Require struct {
		ComposerInstallers string `json:"composer/installers"`
	} `json:"require"`
}

// Events structure to hold the contents of the composer.json file
type Events struct {
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
	wordpress.Version = number[1]
	if wordpress.Name+":"+wordpress.Version == plugin {
		execute()
	} else {
		fmt.Println("plugin name does not match composer.json entry - program halted")
	}
}

// A sequential list of tasks run to complete the program
func execute() {
	prepare()
	checkout(upbranch)
	script()
	jsonWrite()
	commit()
	tags()
	push(upbranch)
}

// Read the composer.json file and store the results in the WordPress structure
func jsonParse() {
	current, _ := os.Open("composer.json")
	defer current.Close()
	byteValue, _ := io.ReadAll(current)
	json.Unmarshal(byteValue, &wordpress)
	json.Unmarshal(byteValue, &events)
}

// Run the update script
func script() {
	exec.Command("/bin/bash", "-c", "scripts/update.sh ~/Downloads/"+folder[1]+"/").Run()
}

// Convert the WordPress structure back into json and overwrite the composer.json file
func jsonWrite() {
	var updated []byte
	if strings.Contains(events.Name, "event") {
		updated, _ = json.MarshalIndent(events, "", "    ")
	} else {
		updated, _ = json.MarshalIndent(wordpress, "", "    ")
	}
	os.WriteFile("composer.json", updated, 0644)
}

// Tag the version so Satis can package it
func tags() {
	exec.Command("git", "tag", "v"+wordpress.Version).Run()
	exec.Command("git", "push", "origin", "--tags").Run()
}
