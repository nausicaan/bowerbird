package main

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
	"time"
)

// Atlassian builds a list of jira tokens and api addresses
type Atlassian struct {
	Base  string `json:"base"`
	Token string `json:"token"`
}

// Desso contains the JSON parameters for a new Jira ticket
type Desso struct {
	Total  int64 `json:"total"`
	Issues []struct {
		Key    string `json:"key,omitempty"`
		Fields struct {
			Issuetype struct {
				ID string `json:"id,omitempty"`
			} `json:"issuetype,omitempty"`
			Status struct {
				Category struct {
					ID   int64  `json:"id,omitempty"`
					Name string `json:"name,omitempty"`
				} `json:"statusCategory,omitempty"`
			} `json:"status,omitempty"`
			Project struct {
				ID  string `json:"id,omitempty"`
				Key string `json:"key,omitempty"`
			} `json:"project,omitempty"`
			Labels      []string     `json:"labels,omitempty"`
			Updated     string       `json:"updated,omitempty"`
			Summary     string       `json:"summary,omitempty"`
			FixVersions []FixVersion `json:"fixVersions,omitempty"`
		} `json:"fields,omitempty"`
	} `json:"issues,omitempty"`
}

// FixVersion holds the information about a release
type FixVersion struct {
	Name        string `json:"name,omitempty"`
	Released    bool   `json:"released,omitempty"`
	ReleaseDate string `json:"releaseDate,omitempty"`
}

const (
	bv        string  = "2.1"
	benchmark float64 = 168.00
	upbranch  string  = "update/"
	relbranch string  = "release/"
	halt      string  = "program halted "
	zero      string  = "Insufficient arguments supplied -"
)

var (
	inputs    int
	desso     Desso
	event     Event
	satis     Satis
	plugin    string
	release   string
	ticket    string
	folder    []string
	number    []string
	updates   []string
	jira      Atlassian
	flag      = os.Args[1]
	hmdr, _   = os.UserHomeDir()
	reader    = bufio.NewReader(os.Stdin)
	common    = hmdr + "/Documents/common/"
	bitbucket = hmdr + "/Documents/bitbucket/"
)

// Parse files created by Silkworm
func discovery(filepath string) {
	goals := read(filepath)
	updates = strings.Split(string(goals), " ")
	updates = updates[:len(updates)-1]
	inputs = len(updates)
}

// Grab the ticket information from Jira in order to extract the DESSO-XXXX identifier
func apiget(summary string) {
	/* Test method to aquire data for the result variable */
	result := read(common + summary)
	// swimlane1 := execute("-c", "curl", "-X", "GET", "-H", "Authorization: Bearer "+jira.Token, "-H", "Content-Type: application/json", jira.Base+todo)
	// swimlane2 := execute("-c", "curl", "-X", "GET", "-H", "Authorization: Bearer "+jira.Token, "-H", "Content-Type: application/json", jira.Base+testing)
	// result := execute("-c", "curl", "-X", "GET", "-H", "Authorization: Bearer "+jira.Token, "-H", "Content-Type: application/json", jira.Base+"search?jql=summary~%27"+summary+"%27")
	json.Unmarshal(result, &desso)
}

// Calculate the difference between two ISO 8601 formatted units of time
func subtract(bigger, smaller string) (time.Duration, error) {
	Minuend, err := time.Parse(time.RFC3339Nano, bigger)
	inspect(err)
	Subtrahend, err := time.Parse(time.RFC3339Nano, smaller)
	inspect(err)
	return Minuend.Sub(Subtrahend), nil
}

// Get the current time and
func amount(lastUpdated string) time.Duration {
	currentTime := time.Now().Format(time.RFC3339Nano)
	// lastyUpdated := "2023-11-01T12:14:09.920-07:00"
	duration, err := subtract(currentTime, lastUpdated)
	inspect(err)
	return duration
}
