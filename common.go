package main

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
	"time"
)

// Desso holds the values needed to move updates from testing to release
type Desso struct {
	Issues []struct {
		Fields struct {
			Project struct {
				ID string `json:"id"`
			} `json:"project"`
			Updated []string `json:"updated"`
			Labels  []string `json:"labels"`
			Status  struct {
				ID string `json:"id"`
			} `json:"status"`
		}
	}
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
	flag      = os.Args[1]
	hmdr, _   = os.UserHomeDir()
	reader    = bufio.NewReader(os.Stdin)
	common    = hmdr + "/Documents/common/"
	bitbucket = hmdr + "/Documents/bitbucket/"
)

func discovery(filepath string) {
	goals := read(filepath)
	updates = strings.Split(string(goals), " ")
	updates = updates[:len(updates)-1]
	inputs = len(updates)
}

// Read any file and return the contents as a byte variable
func read(file string) []byte {
	outcome, problem := os.ReadFile(file)
	inspect(problem)
	return outcome
}

// Open a file and append a string
func atf(name, content string) {
	// Open a file for appending, create it if it doesn't exist
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	inspect(err)
	defer file.Close()

	// Write the content to the file
	_, err = file.WriteString(content)
	inspect(err)
}

// Grab the ticket information from Jira in order to extract the DESSO-XXXX identifier
func apiget(ticket string) {
	/* Test method to aquire data for the result variable */
	result := read(common + "db/search.json")
	// result := execute("-c", "curl", "-X", "GET", "-H", "Authorization: Bearer "+jira.Token, "-H", "Content-Type: application/json", jira.Base+"search?jql=summary~%27"+ticket+"%27")
	json.Unmarshal(result, &desso)
}

func difference(past, present string) (time.Duration, error) {
	Subtrahend, err := time.Parse(time.RFC3339Nano, past)
	inspect(err)
	Minuend, err := time.Parse(time.RFC3339Nano, present)
	inspect(err)
	return Minuend.Sub(Subtrahend), nil
}

func amount() time.Duration {
	currentTime := time.Now().Format(time.RFC3339Nano)
	lastUpdated := "2023-11-01T12:14:09.920-07:00"
	duration, err := difference(lastUpdated, currentTime)
	inspect(err)
	return duration
}
