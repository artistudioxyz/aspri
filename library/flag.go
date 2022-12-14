package library

import "flag"

/** Flag Struct */
type Flag struct {
	/** Mode */
	Docker               *bool
	Git                  *bool
	DockerComposeRestart *bool
	QuoteofTheDay        *bool
	SearchandReplace     *bool
	WPPluginBuild        *bool
	WPPluginBuildCheck   *bool
	WPThemeBuild         *bool
	WPThemeBuildCheck    *bool
	WPRefactor           *bool
	SelfUpdate           *bool

	/** Bool Parameters */
	Production *bool
	Prune      *bool
	Version    *bool

	/** String Parameters */
	ID      *string
	From    *string
	Message *string
	Path    *string
	To      *string
	Type    *string
}

/** Get Flag */
func GetFlag() Flag {
	flags := Flag{
		/** Mode */
		Docker:               flag.Bool("docker", false, "Docker Mode"),
		DockerComposeRestart: flag.Bool("docker-compose-restart", false, "Docker Compose Restart"),
		Git:                  flag.Bool("git", false, "Git Mode"),
		WPPluginBuild:        flag.Bool("wp-plugin-build", false, "WP Build Plugin Comply"),
		WPPluginBuildCheck:   flag.Bool("wp-plugin-build-check", false, "WP Check Plugin Comply with WordPress.org (Version Check)"),
		WPThemeBuild:         flag.Bool("wp-theme-build", false, "WP Theme Plugin Comply"),
		WPThemeBuildCheck:    flag.Bool("wp-theme-build-check", false, "WP Check Theme Comply with WordPress.org (Version Check)"),
		WPRefactor:           flag.Bool("wp-refactor", false, "Refactor Library"),
		QuoteofTheDay:        flag.Bool("quote-of-the-day", false, "show quote of the day"),
		SearchandReplace:     flag.Bool("search-replace", false, "do search and replace"),
		SelfUpdate:           flag.Bool("self-update", false, "self update"),

		/** Bool Parameters */
		Production: flag.Bool("production", false, "Production (WP Mode): Production Environment"),
		Prune:      flag.Bool("prune", false, "Prune (Docker Mode): Container"),
		Version:    flag.Bool("version", false, "show current version"),

		/** String Parameters */
		ID:      flag.String("id", "", "Identifier (Docker Mode): Container"),
		Path:    flag.String("path", "", "Refactor : Path to Directory"),
		From:    flag.String("from", "", "Refactor Text From"),
		Message: flag.String("m", "", "Message (Git Mode): Commit Message"),
		To:      flag.String("to", "", "Refactor Text To"),
		Type:    flag.String("type", "", "Build type (WordPress)"),
	}

	return flags
}
