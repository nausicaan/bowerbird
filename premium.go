package main

import (
	"encoding/json"
	"os"
	"strings"
)

// Satis structure captures the contents of the composer.json file for typical premium plugins
type Satis struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`
}

// Event structure captures the contents of the composer.json file for Events Calendar related items
type Event struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`
	Require struct {
		EventsCalendar string `json:"wpackagist-plugin/the-events-calendar"`
	} `json:"require"`
}

// A sequential list of tasks run to complete the program
func quarterback() {
	prepare()
	checkout(upbranch)
	tracking("Update Script")
	script()
	correct()
	commit()
	tracking("Tagging to Satis")
	// tags()
	// tracking("Git Push")
	// push()
}

// Premium directs the preliminary actions to determine if the program can continue
func premium() {
	plugin, ticket = updates[0], updates[1]
	assign()
	os.Chdir(hmdr + "/Documents/bitbucket/" + folder[1])
	serialize()
	satis.Version, event.Version = number[1], number[1]

	if strings.Contains(folder[1], "event") {
		if event.Name+":"+event.Version == plugin {
			quarterback()
		}
	} else if satis.Name+":"+satis.Version == plugin {
		quarterback()
	} else {
		alert("Plugin name does not match composer.json entry -")
	}
}

// Read the composer.json file and store the results in a structure
func serialize() {
	current, _ := os.ReadFile("composer.json")
	err := json.Unmarshal(current, &satis)
	inspect(err)
	err = json.Unmarshal(current, &event)
	inspect(err)
}

// Split the supplied arguments and assign them to variables
func assign() {
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
	if strings.Contains(event.Name, "event") {
		updated, _ = json.MarshalIndent(event, "", "    ")
	} else {
		updated, _ = json.MarshalIndent(satis, "", "    ")
	}
	document("composer.json", updated)
}

// Tag the version so Satis can package it
func tags() {
	execute("git", "tag", "v"+satis.Version)
	execute("git", "push", "origin", "--tags")
}
