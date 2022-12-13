package library

import (
	"fmt"
)

/** Initiate Docker Function */
func InitiateDockerFunction(flags Flag) {
	/**
	 * Stop and Remove Container
	 * - Equivalent to : `docker stop {identifier} && docker rm {identifier}`
	 */
	if *flags.Docker && *flags.Prune && *flags.ID != "" {
		fmt.Println("Stop and Remove Container")
		snr := fmt.Sprintf("docker stop %s && docker rm %s", *flags.ID, *flags.ID)
		cmd := [...]string{"bash", "-c", snr}
		fmt.Println(ExecCommand(cmd[:]...))
	}
	/**
	 * Compose restart (down & up)
	 * - Equivalent to : `docker-compose down && docker-compose up -d`
	 */
	if *flags.DockerComposeRestart {
		fmt.Println("Compose restart (down & up)")
		cmd := [...]string{"bash", "-c", "docker-compose down && docker-compose up -d;"}
		fmt.Println(ExecCommand(cmd[:]...))
	}
}
