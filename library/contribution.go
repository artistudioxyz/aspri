package library

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Date format used in the markdown files
const dateFormat = "2006-01-02"

// Calculate the date range (Monday-Sunday of the previous week)
func getLastWeekRange() (time.Time, time.Time) {
	now := time.Now()
	weekday := now.Weekday()

	// Calculate the previous Sunday's date
	daysSinceSunday := int(weekday+1) % 7
	sundayLastWeek := now.AddDate(0, 0, -daysSinceSunday)

	// Calculate the previous Monday's date
	mondayLastWeek := sundayLastWeek.AddDate(0, 0, -6)

	return mondayLastWeek, sundayLastWeek
}

// Contributor struct to store the name and contribution date
type Contributor struct {
	Name string
	Date time.Time
}

// Function to parse the contribution log from a markdown file
func parseContributionLog(filePath string, start, end time.Time) ([]Contributor, error) {
	println(filePath)

	var contributors []Contributor

	// Open the markdown file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	print(file)

	scanner := bufio.NewScanner(file)
	inContributorsSection := false
	//contributorPattern := regexp.MustCompile(`- \[\[(.*?)\]\]: (\d{4}-\d{2}-\d{2})`)

	// Scan through each line of the markdown file
	for scanner.Scan() {
		// Check if we're in the "Contributors" section
		if strings.Contains(scanner.Text(), "## Contributors") {
			inContributorsSection = true
			continue
		}

		if inContributorsSection {
			print("MASUK")
		}

		//// If in the contributors section, parse the log entries
		//if inContributorsSection {
		//	if match := contributorPattern.FindStringSubmatch(line); match != nil {
		//		name := match[1]
		//		dateStr := match[2]
		//		contributionDate, err := time.Parse(dateFormat, dateStr)
		//		if err != nil {
		//			return nil, err
		//		}
		//
		//		// Only include contributions within the Monday-Sunday of last week
		//		if contributionDate.After(start) && contributionDate.Before(end.AddDate(0, 0, 1)) {
		//			contributors = append(contributors, Contributor{Name: name, Date: contributionDate})
		//		}
		//	}
		//}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return contributors, nil
}

// Function to traverse the markdown directory and aggregate contributions
func calculateContributions(dirPath string, start, end time.Time) (map[string]int, error) {
	contributionsByUser := make(map[string]int)

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Process only .md files
		if filepath.Ext(path) == ".md" {
			contributors, err := parseContributionLog(path, start, end)
			if err != nil {
				return err
			}

			// Count contributions for each user
			for _, contributor := range contributors {
				contributionsByUser[contributor.Name]++
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return contributionsByUser, nil
}

// Function to generate a report of contributions
func generateReport(contributionsByUser map[string]int, start, end time.Time) {
	fmt.Printf("Contributions from %s to %s:\n\n", start.Format(dateFormat), end.Format(dateFormat))
	for contributor, count := range contributionsByUser {
		fmt.Printf("%s: %d contributions\n", contributor, count)
	}
}

// Main function
func InitiateContribution(flags Flag) {
	/** Directory Stats */
	if *flags.Contribution {
		// Get the Monday-Sunday range of last week
		startOfLastWeek, endOfLastWeek := getLastWeekRange()

		// Check if a start date is provided
		if *flags.DateStart == "" {
			*flags.DateStart = startOfLastWeek.Format(dateFormat)
		}

		// Check if an end date is provided
		if *flags.DateEnd == "" {
			*flags.DateEnd = endOfLastWeek.Format(dateFormat)
		}

		// Calculate contributions in that range
		contributionsByUser, err := calculateContributions(*flags.Path, startOfLastWeek, endOfLastWeek)
		if err != nil {
			fmt.Printf("Error calculating contributions: %v\n", err)
			return
		}

		print(contributionsByUser)

		// Generate and print the report
		//generateReport(contributionsByUser, startOfLastWeek, endOfLastWeek)
	}
}
