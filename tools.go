package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

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

// Open a file for reading and return an os.File variable
func expose(file string) *os.File {
	outcome, err := os.Open(file)
	inspect(err)
	return outcome
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

// Run standard terminal commands and display the output
func execute(task string, args ...string) {
	osCmd := exec.Command(task, args...)
	osCmd.Stdout = os.Stdout
	osCmd.Stderr = os.Stderr
	err := osCmd.Run()
	inspect(err)
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
