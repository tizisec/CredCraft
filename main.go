package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	// Check if -w and -t flags are used
	if len(os.Args) < 5 || os.Args[1] != "-w" || os.Args[3] != "-t" {
		log.Fatal("Invalid arguments. Usage: go run main.go -w <cr_names> -t <cr_templates>")
	}

	crNamesFilePath := os.Args[2]
	crTemplatesFilePath := os.Args[4]

	// Read cr_names lines into a slice
	crNamesLines, err := readFileLines(crNamesFilePath)
	if err != nil {
		log.Fatalf("Failed to read cr_names file: %s", err)
	}

	// Read cr_templates lines into a slice
	crTemplatesLines, err := readFileLines(crTemplatesFilePath)
	if err != nil {
		log.Fatalf("Failed to read cr_templates file: %s", err)
	}

	// Store the modified lines in the result variable
	var result []string
	for _, template := range crTemplatesLines {
		// Replace "{}" in cr_templates lines with each line from cr_names
		for _, name := range crNamesLines {
			modifiedLine := strings.Replace(template, "{}", name, 1)
			result = append(result, modifiedLine)
		}
	}

	// Print the result
	for _, line := range result {
		println(line)
	}

	log.Println("Replacement completed successfully.")
}

// Read file lines into a slice
func readFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
