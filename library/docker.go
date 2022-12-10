package library

import (
	"flag"
	"fmt"
)

var (
	/** Command */
	DockerStopandRemoveFlag  = flag.Bool("docker", false, "docker mode")
	DockerPruneContainerFlag = flag.Bool("prune", false, "docker stop and remove container")
	DockerComposeRestartFlag = flag.Bool("docker-compose-restart", false, "docker compose restart")

	/** Identifier */
	DockerContainerIDFlag = flag.String("id", "", "docker container identifier")
)

/** Initiate Docker Function */
func InitiateDockerFunction() {
	/**
	 * Stop and Remove Container
	 * - Equivalent to : `docker stop {identifier} && docker rm {identifier}`
	 */
	if *DockerStopandRemoveFlag && *DockerPruneContainerFlag && *DockerContainerIDFlag != "" {
		fmt.Println("Stop and Remove Container")
		snr := fmt.Sprintf("docker stop %s && docker rm %s", *DockerContainerIDFlag, *DockerContainerIDFlag)
		cmd := [...]string{"bash", "-c", snr}
		fmt.Println(ExecCommand(cmd[:]...))
	}
	/**
	 * Compose restart (down & up)
	 * - Equivalent to : `docker-compose down && docker-compose up -d`
	 */
	if *DockerComposeRestartFlag {
		fmt.Println("Compose restart (down & up)")
		cmd := [...]string{"bash", "-c", "docker-compose down && docker-compose up -d;"}
		fmt.Println(ExecCommand(cmd[:]...))
	}
}
