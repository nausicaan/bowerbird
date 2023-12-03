package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func convert() string {
	lfv, err := strconv.Atoi(history.Issues[0].Fields.FixVersions[0].Name)
	inspect(err)
	return fmt.Sprint(lfv)
}

// Calculate the difference between two ISO 8601 formatted units of time
func subtract(bigger, smaller string) (time.Duration, error) {
	Minuend, err := time.Parse(time.RFC3339Nano, bigger)
	inspect(err)
	Subtrahend, err := time.Parse(time.RFC3339Nano, smaller)
	inspect(err)
	return Minuend.Sub(Subtrahend), nil
}

// Get the amount of time since a ticket was last updated
func amount(lastUpdated string) time.Duration {
	currentTime := time.Now().Format(time.RFC3339Nano)
	duration, err := subtract(currentTime, lastUpdated)
	inspect(err)
	return duration
}

// Get user input via screen prompt
func solicit(prompt string) string {
	fmt.Print(prompt)
	response, _ := reader.ReadString('\n')
	return strings.TrimSpace(response)
}

// Write a passed variable to a named file
func document(name string, d []byte) {
	inspect(os.WriteFile(name, d, 0644))
}

// Read any file and return the contents as a byte variable
func read(file string) []byte {
	outcome, problem := os.ReadFile(file)
	inspect(problem)
	return outcome
}

// Open a file for reading and return an os.File variable
func expose(file string) *os.File {
	outcome, err := os.Open(file)
	inspect(err)
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

// Record a list of files in a folder
func ls(folder string) []string {
	var content []string
	dir := expose(folder)

	files, err := dir.ReadDir(0)
	inspect(err)

	for _, f := range files {
		content = append(content, f.Name())
	}
	return content
}

// Run a terminal command using flags to customize the output
func execute(variation, task string, args ...string) []byte {
	osCmd := exec.Command(task, args...)
	switch variation {
	case "-e":
		exec.Command(task, args...).CombinedOutput()
	case "-c":
		both, _ := osCmd.CombinedOutput()
		return both
	case "-v":
		osCmd.Stdout = os.Stdout
		osCmd.Stderr = os.Stderr
		err := osCmd.Run()
		inspect(err)
	}
	return nil
}

// Remove files or directories
func cleanup(cut ...string) {
	inspect(os.Remove(cut[0.]))
}

// Check for errors, print the result if found
func inspect(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
