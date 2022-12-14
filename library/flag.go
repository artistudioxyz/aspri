package library

import (
	flag "github.com/spf13/pflag"
)

/** Flag Struct */
type Flag struct {
	/** Mode */
	ChatGPT              *bool
	Dir                  *bool
	Docker               *bool
	File                 *bool
	Git                  *bool
	GitResetCache        *bool
	DockerComposeRestart *bool
	ListClass            *bool
	ListFunction         *bool
	Markdown             *bool
	PHP                  *bool
	QuoteofTheDay        *bool
	Remove               *bool
	RemoveLink           *bool
	RemoveFunction       *bool
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
	API_KEY      *string
	ID           *string
	Dirname      *[]string
	Ext          *[]string
	Except       *[]string
	Filename     *[]string
	FunctionName *[]string
	From         *string
	Message      *string
	Regex        *string
	Path         *string
	To           *string
	Type         *string
}

/** Get Flag */
func GetFlag() Flag {
	flags := Flag{
		/** Mode */
		ChatGPT:              flag.Bool("chatgpt", false, "Chat with GPT-3"),
		Dir:                  flag.Bool("dir", false, "Directory Mode"),
		Docker:               flag.Bool("docker", false, "Docker Mode"),
		DockerComposeRestart: flag.Bool("docker-compose-restart", false, "Docker Compose Restart"),
		File:                 flag.Bool("file", false, "File Mode"),
		Git:                  flag.Bool("git", false, "Git Mode"),
		GitResetCache:        flag.Bool("git-reset-cache", false, "Git Reset Cache"),
		ListClass:            flag.Bool("list-class", false, "List Class"),
		ListFunction:         flag.Bool("list-function", false, "List Function"),
		Markdown:             flag.Bool("md", false, "Markdown Mode"),
		PHP:                  flag.Bool("php", false, "PHP Mode"),
		Remove:               flag.Bool("remove", false, "Remove Mode for Dir and File"),
		RemoveLink:           flag.Bool("remove-link", false, "Remove Link from File"),
		RemoveFunction:       flag.Bool("remove-function", false, "Remove Link from File"),
		QuoteofTheDay:        flag.Bool("quote-of-the-day", false, "show quote of the day"),
		SearchandReplace:     flag.Bool("search-replace", false, "do search and replace"),
		SelfUpdate:           flag.Bool("self-update", false, "self update"),
		WPPluginBuild:        flag.Bool("wp-plugin-build", false, "WP Build Plugin Comply"),
		WPPluginBuildCheck:   flag.Bool("wp-plugin-build-check", false, "WP Check Plugin Comply with WordPress.org (Version Check)"),
		WPThemeBuild:         flag.Bool("wp-theme-build", false, "WP Theme Plugin Comply"),
		WPThemeBuildCheck:    flag.Bool("wp-theme-build-check", false, "WP Check Theme Comply with WordPress.org (Version Check)"),
		WPRefactor:           flag.Bool("wp-refactor", false, "Refactor Library"),

		/** Bool Parameters */
		Production: flag.Bool("production", false, "Production (WP Mode): Production Environment"),
		Prune:      flag.Bool("prune", false, "Prune (Docker Mode): Container"),
		Version:    flag.Bool("version", false, "show current version"),

		/** String Parameters */
		API_KEY:      flag.String("api-key", "", "API Key"),
		ID:           flag.String("id", "", "Identifier (Docker Mode): Container"),
		Dirname:      flag.StringArray("dirname", []string{}, "Directory Name (Dir Mode): Directory Name"),
		Ext:          flag.StringArray("ext", []string{}, "File extensions to include"),
		Except:       flag.StringArray("except", []string{}, "File to exclude"),
		Filename:     flag.StringArray("filename", []string{}, "Filenames"),
		FunctionName: flag.StringArray("functionname", []string{}, "Function Name"),
		From:         flag.String("from", "", "Refactor Text From"),
		Message:      flag.StringP("message", "m", "", "Message (Git Mode): Commit Message"),
		Path:         flag.String("path", "", "Refactor : Path to Directory"),
		Regex:        flag.String("regex", "", "Regex"),
		To:           flag.String("to", "", "Refactor Text To"),
		Type:         flag.String("type", "", "Build type (WordPress)"),
	}
	flag.Parse()

	return flags
}
