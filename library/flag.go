package library

import "flag"

/** Flag Struct */
type Flag struct {
	/** Mode */
	Docker                    *bool
	Git                       *bool
	DockerComposeRestart      *bool
	QuoteofTheDay             *bool
	SearchandReplaceDirectory *bool
	WPPluginBuildCheck        *bool
	WPRefactor                *bool

	/** Bool Parameters */
	Production *bool
	Prune      *bool

	/** String Parameters */
	ID      *string
	From    *string
	Message *string
	Path    *string
	To      *string
}

/** Get Flag */
func GetFlag() Flag {
	flags := Flag{
		/** Mode */
		Docker:                    flag.Bool("docker", false, "Docker Mode"),
		DockerComposeRestart:      flag.Bool("docker-compose-restart", false, "Docker Compose Restart"),
		Git:                       flag.Bool("git", false, "Git Mode"),
		WPPluginBuildCheck:        flag.Bool("wp-plugin-build-check", false, "WP Check Plugin Comply with Directory"),
		WPRefactor:                flag.Bool("wp-refactor", false, "Refactor Library"),
		QuoteofTheDay:             flag.Bool("quote-of-the-day", false, "show quote of the day"),
		SearchandReplaceDirectory: flag.Bool("search-replace-directory", false, "do search and replace"),

		/** Bool Parameters */
		Production: flag.Bool("production", false, "Production (WP Mode): Production Environment"),
		Prune:      flag.Bool("prune", false, "Prune (Docker Mode): Container"),

		/** String Parameters */
		ID:      flag.String("id", "", "Identifier (Docker Mode): Container"),
		Path:    flag.String("path", "", "Refactor : Path to Directory"),
		From:    flag.String("from", "", "Refactor Text From"),
		Message: flag.String("m", "", "Message (Git Mode): Commit Message"),
		To:      flag.String("to", "", "Refactor Text To"),
	}

	return flags
}
