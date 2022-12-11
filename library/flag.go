package library

import "flag"

/** Flag Struct */
type Flag struct {
	/** Mode */
	Docker               *bool
	DockerComposeRestart *bool
	WPRefactor           *bool

	/** Additional Mode */
	Prune *bool

	/** Parameters */
	ID *string
}

/** Get Flag */
func GetFlag() Flag {
	flags := Flag{}

	/** Mode */
	flags.Docker = flag.Bool("docker", false, "Docker Mode")
	flags.DockerComposeRestart = flag.Bool("docker-compose-restart", false, "Docker Compose Restart")
	flags.WPRefactor = flag.Bool("wp-refactor", false, "Refactor Library")

	/** Parameters */
	flags.Prune = flag.Bool("prune", false, "Prune (Docker Mode): Container")
	flags.ID = flag.String("id", "", "Identifier (Docker Mode): Container")

	return flags
}
