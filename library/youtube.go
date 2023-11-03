package library

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strings"
)

/** Initiate Syncthing Function */
func InitiateYouTubeFunction(flags Flag) {
	/** remove Sync Conflict Files older Than x days */
	if *flags.YouTube && *flags.Extract && *flags.Path != "" {
		ExtractYouTubeData(*flags.Path)
	}
}

/** Extract YouTube Data */
func ExtractYouTubeData(inputFilePath string) {
	// Provide the output file path
	outputFilePath := "output.csv"

	// Read the data from the input file
	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create the output CSV file
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Process each line of the input file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimLeft(line, " ")

		// Get the title
		re := regexp.MustCompile(`^(.*?)\d`)
		match := re.FindStringSubmatch(line)
		title := ""
		if len(match) > 1 {
			title = strings.Join(strings.Fields(match[1]), " ")
		}

		// Get the view count
		re = regexp.MustCompile(`\b([\d,]+) views\b`)
		match = re.FindStringSubmatch(line)
		views := ""
		if len(match) > 1 {
			views = match[1]
		}

		// Get the release date
		re = regexp.MustCompile(`(\d+ [a-zA-Z]+ ago)`)
		match = re.FindStringSubmatch(line)
		release := ""
		if len(match) > 1 {
			release = match[1]
		}

		// Get the length
		re = regexp.MustCompile(`(\d+ [a-zA-Z]+, \d+ [a-zA-Z]+)`)
		match = re.FindStringSubmatch(line)
		length := ""
		if len(match) > 1 {
			length = match[1]
		}

		// Get the link
		re = regexp.MustCompile(`\bhttps://www.youtube.com/watch\?v=([\w-]+)\b`)
		match = re.FindStringSubmatch(line)
		link := ""
		if len(match) > 1 {
			link = "https://www.youtube.com/watch?v=" + match[1]
		}

		// Transform the data into a CSV record
		record := []string{
			title,
			views,
			release,
			length,
			link,
		}

		// Write the record to the CSV file
		err := writer.Write(record)
		if err != nil {
			fmt.Println("❌ Error writing to CSV:", err)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("❌ Error reading file:", err)
		return
	}

	fmt.Println("✅ CSV file created successfully.")
}
