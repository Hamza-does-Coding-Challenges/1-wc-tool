package main

import (
	"fmt"
	"os"
	"strings"
)

func parseCommandLineArgs() (string, []string, error) {
	argsOfProgram := os.Args

	if len(argsOfProgram) == 0 {
		return "", nil, fmt.Errorf("No arguments provided")
	}

	var fileName string = ""
	var options []string = []string{}
	for i := 1; i < len(argsOfProgram); i++ {

		splitted := strings.Split(argsOfProgram[i], ".")
		if len(splitted) == 2 {
			fileName = argsOfProgram[i]
		} else if strings.HasPrefix(argsOfProgram[i], "-") {
			options = append(options, argsOfProgram[i])
		}
	}

	if fileName == "" {
		return "", nil, fmt.Errorf("No file name provided")
	}

	return fileName, options, nil
}

// Simply read file and return contents
func readFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return data
}

// Used when "-c" option is provided
func countFileBytes(path string) int {
	data := readFile(path)
	if data == nil {
		return 0
	}
	return len(data)
}

// Used when "-l" option is provided
func countFileLines(path string) int {
	data := readFile(path)
	if data == nil {
		return 0
	}
	lines := strings.Split(string(data), "\n")

	return len(lines)
}

// Use when "-w" option is provided
func countFileWords(path string) int {
	data := readFile(path)

	if data == nil {
		return 0
	}

	// Split data by whitespace and return number of words
	// Using strings.Fields() here to account for multiple
	// white spaces between words
	words := strings.Fields(string(data))

	return len(words)
}

func outputParser(fileName string, options []string) string {
	var finalOutput string = ""

	for i := 0; i < len(options); i++ {
		if options[i] == "-c" {
			finalOutput += fmt.Sprintf("%d", countFileBytes(fileName)) + " "
		}

		if options[i] == "-l" {
			finalOutput += fmt.Sprintf("%d", countFileLines(fileName)) + " "
		}

		if options[i] == "-m" {
			finalOutput += fmt.Sprintf("%d", countFileWords(fileName)) + " "
		}
	}

	return finalOutput + fileName
}

func main() {
	fileName, options, err := parseCommandLineArgs()

	if err != nil {
		fmt.Println("Error parsing command line arguments:", err)
		return
	}

	fmt.Println("File name: ", fileName)
	fmt.Println("Options: ", options)

	fmt.Println(outputParser(fileName, options))
}
