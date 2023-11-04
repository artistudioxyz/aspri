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
		line = strings.Join(strings.Fields(line), " ")

		// Get the view count
		re := regexp.MustCompile(`\b([\d,]+) views\b`)
		match := re.FindStringSubmatch(line)
		views := ""
		viewsRaw := ""
		if len(match) > 1 {
			views = strings.Replace(match[1], ",", "", -1)
			viewsRaw = match[1]
		}

		// Get the title
		re = regexp.MustCompile(`^(.*?) views`)
		match = re.FindStringSubmatch(line)
		title := ""
		if len(match) > 1 {
			title = strings.Replace(match[1], viewsRaw, "", -1)
		}

		// Get the release date
		re = regexp.MustCompile(`(\d+ [a-zA-Z]+ ago)`)
		match = re.FindStringSubmatch(line)
		release := ""
		if len(match) > 1 {
			release = match[1]
		}

		// Get the length
		re = regexp.MustCompile(`(\d+ [a-zA-Z]+ https)`)
		match = re.FindStringSubmatch(line)
		length := ""
		if len(match) > 1 {
			length = strings.Replace(match[1], "https", "", -1)
		}

		// Get the link
		re = regexp.MustCompile(`\bhttps://www.youtube.com/watch\?v=([\w-]+)\b`)
		match = re.FindStringSubmatch(line)
		link := ""
		if len(match) > 1 {
			link = "https://www.youtube.com/watch?v=" + match[1]
		}

		// Debug
		//fmt.Println(title)
		//fmt.Println(views)
		//fmt.Println(viewsRaw)
		//fmt.Println(release)
		//fmt.Println(length)
		//fmt.Println(link)

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
