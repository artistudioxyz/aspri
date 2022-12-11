package library

import "flag"

/** Flag Struct */
type Flag struct {
	/** Mode */
	Docker               *bool
	Git                  *bool
	DockerComposeRestart *bool
	WPRefactor           *bool

	/** Additional Mode */
	Prune *bool
	Push  *bool

	/** Parameters */
	ID      *string
	Message *string
}

/** Get Flag */
func GetFlag() Flag {
	flags := Flag{}

	/** Mode */
	flags.Docker = flag.Bool("docker", false, "Docker Mode")
	flags.DockerComposeRestart = flag.Bool("docker-compose-restart", false, "Docker Compose Restart")
	flags.Git = flag.Bool("git", false, "Git Mode")
	flags.WPRefactor = flag.Bool("wp-refactor", false, "Refactor Library")

	/** Parameters */
	flags.ID = flag.String("id", "", "Identifier (Docker Mode): Container")
	flags.Prune = flag.Bool("prune", false, "Prune (Docker Mode): Container")
	flags.Push = flag.Bool("push", false, "Push (Git Mode): Commit and Push")
	flags.Message = flag.String("m", "", "Message (Git Mode): Commit Message")

	return flags
}
