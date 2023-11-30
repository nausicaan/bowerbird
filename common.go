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

// Desso holds the values needed to move updates from testing to release
type Desso struct {
	Key    string `json:"key"`
	Fields struct {
		Project struct {
			ID string `json:"id"`
		} `json:"project"`
		Updated string `json:"updated"`
		Status  struct {
			Category struct {
				ID   int64  `json:"id"`
				Name string `json:"name"`
			} `json:"statusCategory"`
		} `json:"status"`
		Labels      []string     `json:"labels"`
		Summary     string       `json:"summary"`
		FixVersions []FixVersion `json:"fixVersions"`
	} `json:"fields"`
}

// FixVersion holds the information about a release
type FixVersion struct {
	Name        string `json:"name"`
	Released    bool   `json:"released"`
	ReleaseDate string `json:"releaseDate"`
}

const (
	bv        string = "2.1"
	upbranch  string = "update/"
	relbranch string = "release/"
	halt      string = "program halted "
	zero      string = "Insufficient arguments supplied -"
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
func apiget(ticket string) {
	/* Test method to aquire data for the result variable */
	result := read(common + "db/d1510.json")
	// result := execute("-c", "curl", "-X", "GET", "-H", "Authorization: Bearer "+jira.Token, "-H", "Content-Type: application/json", jira.Base+"search?jql=summary~%27"+ticket+"%27")
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
