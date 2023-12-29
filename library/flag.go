package library

import (
	flag "github.com/spf13/pflag"
)

// Flag Structure
type Flag struct {
	// Mode
	ChatGPT            *bool
	Dir                *bool
	Docker             *bool
	DockerCompose      *bool
	DryRun             *bool
	Extract            *bool
	ExtractUrl         *bool
	File               *bool
	Git                *bool
	Gone               *bool
	Help               *bool
	Install            *bool
	ListClass          *bool
	ListFunction       *bool
	ListFunctionCall   *bool
	Minify             *bool
	Markdown           *bool
	NoIP               *bool
	OlderThan          *bool
	PHP                *bool
	PHPCS              *bool
	QuoteofTheDay      *bool
	Remove             *bool
	RemoveConflicts    *bool
	RemoveLink         *bool
	RemoveFunction     *bool
	ResetCache         *bool
	Stats              *bool
	SearchandReplace   *bool
	Syncthing          *bool
	WPClean            *bool
	WPPluginBuild      *bool
	WPPluginBuildCheck *bool
	WPThemeBuild       *bool
	WPThemeBuildCheck  *bool
	WPRefactor         *bool
	SelfUpdate         *bool
	Tree               *bool
	Update             *bool
	YouTube            *bool

	// Bool Parameters
	Count      *bool
	Production *bool
	Prune      *bool
	Restart    *bool
	Version    *bool

	// String Parameters
	API_KEY      *string
	ID           *string
	Days         *int
	Dirname      *[]string
	Ext          *[]string
	Except       *[]string
	Filename     *[]string
	FunctionName *[]string
	From         *string
	Hostname     *string
	IgnoreDirs   *[]string
	Message      *string
	Path         *string
	Password     *string
	Regex        *string
	Text         *string
	To           *string
	Type         *string
	Url          *string
	Username     *string
}

// Get Flag
func GetFlag() Flag {
	flags := Flag{
		// Mode
		ChatGPT:            flag.Bool("chatgpt", false, "Chat with GPT-3"),
		Dir:                flag.Bool("dir", false, "Directory Mode"),
		Docker:             flag.Bool("docker", false, "Docker Mode"),
		DockerCompose:      flag.Bool("docker-compose", false, "Docker Compose Mode"),
		DryRun:             flag.Bool("dry-run", false, "Dry Run Mode"),
		Extract:            flag.Bool("extract", false, "Extract Mode"),
		ExtractUrl:         flag.Bool("extract-url", false, "Extract URL Mode"),
		File:               flag.Bool("file", false, "File Mode"),
		Git:                flag.Bool("git", false, "Git Mode"),
		Gone:               flag.Bool("gone", false, "Gone Mode"),
		Help:               flag.Bool("help", false, "Help Mode"),
		Install:            flag.Bool("install", false, "Install Mode"),
		ListClass:          flag.Bool("list-class", false, "List Class"),
		ListFunction:       flag.Bool("list-function", false, "List Function"),
		ListFunctionCall:   flag.Bool("list-function-call", false, "List Function Call"),
		Markdown:           flag.Bool("md", false, "Markdown Mode"),
		Minify:             flag.Bool("minify", false, "Minify Mode"),
		NoIP:               flag.Bool("noip", false, "No-IP Mode"),
		OlderThan:          flag.Bool("older-than", false, "Older Than Mode"),
		PHP:                flag.Bool("php", false, "PHP Mode"),
		PHPCS:              flag.Bool("phpcs", false, "PHP Code Sniffer Mode"),
		Remove:             flag.Bool("remove", false, "Remove Mode for Dir and File"),
		RemoveConflicts:    flag.Bool("remove-conflicts", false, "Remove Conflicts"),
		RemoveLink:         flag.Bool("remove-link", false, "Remove Link from File"),
		RemoveFunction:     flag.Bool("remove-function", false, "Remove Link from File"),
		ResetCache:         flag.Bool("reset-cache", false, "Git Reset Cache"),
		Syncthing:          flag.Bool("syncthing", false, "Syncthing Mode"),
		QuoteofTheDay:      flag.Bool("quote-of-the-day", false, "show quote of the day"),
		Restart:            flag.Bool("restart", false, "Restart (Docker Mode): Container"),
		SearchandReplace:   flag.Bool("search-replace", false, "do search and replace"),
		SelfUpdate:         flag.Bool("self-update", false, "self update"),
		Stats:              flag.Bool("stats", false, "show stats"),
		Tree:               flag.Bool("tree", false, "Tree Mode"),
		Update:             flag.Bool("update", false, "update"),
		WPClean:            flag.Bool("wp-clean", false, "WP Clean Project Files for Production"),
		WPPluginBuild:      flag.Bool("wp-plugin-dist", false, "WP Build Plugin Comply"),
		WPPluginBuildCheck: flag.Bool("wp-plugin-dist-check", false, "WP Check Plugin Comply with WordPress.org (Version Check)"),
		WPThemeBuild:       flag.Bool("wp-theme-dist", false, "WP Theme Plugin Comply"),
		WPThemeBuildCheck:  flag.Bool("wp-theme-dist-check", false, "WP Check Theme Comply with WordPress.org (Version Check)"),
		WPRefactor:         flag.Bool("wp-refactor", false, "Refactor Library"),
		YouTube:            flag.Bool("youtube", false, "YouTube Mode"),

		// Bool Parameters
		Count:      flag.Bool("count", false, "Count Mode"),
		Production: flag.Bool("production", false, "Production (WP Mode): Production Environment"),
		Prune:      flag.Bool("prune", false, "Prune (Docker Mode): Container"),
		Version:    flag.Bool("version", false, "show current version"),

		// String Parameters
		API_KEY:      flag.String("api-key", "", "API Key"),
		ID:           flag.String("id", "", "Identifier (Docker Mode): Container"),
		Days:         flag.Int("days", 0, "Days (Older Than Mode): Days"),
		Dirname:      flag.StringArray("dirname", []string{}, "Directory Name (Dir Mode): Directory Name"),
		Ext:          flag.StringArray("ext", []string{}, "File extensions to include"),
		Except:       flag.StringArray("except", []string{}, "File to exclude"),
		Filename:     flag.StringArrayP("filename", "f", []string{}, "Filenames"),
		FunctionName: flag.StringArray("functionname", []string{}, "Function Name"),
		From:         flag.String("from", "", "Refactor Text From"),
		Hostname:     flag.String("hostname", "", "Hostname"),
		IgnoreDirs:   flag.StringArray("ignoredirs", []string{}, "Directory Name to Ignored"),
		Message:      flag.StringP("message", "m", "", "Message (Git Mode): Commit Message"),
		Path:         flag.String("path", "", "Refactor : Path to Directory"),
		Password:     flag.StringP("password", "p", "", "Password"),
		Regex:        flag.String("regex", "", "Regex"),
		Text:         flag.String("text", "", "Text"),
		To:           flag.String("to", "", "Refactor Text To"),
		Type:         flag.String("type", "", "Build type (WordPress)"),
		Username:     flag.StringP("username", "u", "", "Username"),
		Url:          flag.String("url", "", "Url"),
	}
	flag.Parse()

	return flags
}
