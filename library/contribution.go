package library

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// Date format used in the markdown files
const dateFormat = "2006-01-02" // Default date format

// Initiate Contribution Function
func InitiateContribution(flags Flag) {
	// Calculate contributions in file date range
	if *flags.Contribution {
		count, err := calculateContributions(*flags.Path, *flags.Text, *flags.DateStart, *flags.DateEnd, *flags.Exclude)
		if err != nil {
			fmt.Println("❌ Error calculating contributions:", err)
			return
		}

		fmt.Println("🐙 There are", count, "files containing", *flags.Text)
	}
}

// Calculate the date range (Monday-Sunday of the previous week)
func getLastWeekRange() (time.Time, time.Time) {
	now := time.Now()
	weekday := now.Weekday()

	// Calculate the previous Sunday's date
	daysSinceSunday := int(weekday+1) % 7
	sundayLastWeek := now.AddDate(0, 0, -daysSinceSunday)

	// Calculate the previous Monday's date
	mondayLastWeek := sundayLastWeek.AddDate(0, 0, -5)

	return mondayLastWeek, sundayLastWeek
}

// Function to traverse the markdown directory and aggregate contributions
func calculateContributions(dirPath string, text string, start string, end string, exclude []string) (int, error) {
	var count int
	contributorPattern := regexp.MustCompile(`- \[\[(.*?)\]\]: (\d{4}-\d{2}-\d{2})`)

	// Get the Monday-Sunday range of last week
	startOfLastWeek, endOfLastWeek := getLastWeekRange()

	// Check if a start date is provided
	startDate, err := time.Parse(dateFormat, start)
	if err != nil {
		startDate = startOfLastWeek
	}

	// Check if an end date is provided
	endDate, err := time.Parse(dateFormat, end)
	if err != nil {
		endDate = endOfLastWeek
	}

	// Walk through the directories in the provided path
	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("❌ Error:", err)
			return err
		}

		// Check if the directory should be ignored
		for _, excludeDir := range exclude {
			if strings.Contains(path, excludeDir) {
				return nil
			}
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				// Check if the line contains the contributor's name
				line := scanner.Text()
				if strings.Contains(line, text) {
					if match := contributorPattern.FindStringSubmatch(line); match != nil {
						dateStr := match[2]
						contributionDate, err := time.Parse(dateFormat, dateStr)
						if err != nil {
							return err
						}

						// Only include contributions within the Monday-Sunday of last week
						if contributionDate.After(startDate) && contributionDate.Before(endDate.AddDate(0, 0, 1)) {
							count++
							break
						}
					}
				}
			}
		}

		return nil
	})

	return count, err
}
