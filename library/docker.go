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
		fmt.Println("📟 Stop and Remove Container")
		snr := fmt.Sprintf("docker stop %s && docker rm %s", *flags.ID, *flags.ID)
		cmd := [...]string{"bash", "-c", snr}
		fmt.Println(ExecCommand(cmd[:]...))
	}
	/**
	 * Compose restart (down & up)
	 * - Equivalent to : `docker-compose down && docker-compose up -d`
	 */
	if *flags.DockerComposeRestart {
		fmt.Println("📟 Compose restart (down & up)")
		filename := *flags.Filename
		dockercmd := ""
		if len(*flags.Filename) > 0 {
			dockercmd = fmt.Sprintf("docker-compose -f %s down && docker-compose -f %s up -d;", filename[0], filename[0])
		} else {
			dockercmd = "docker-compose down && docker-compose up -d;"
		}
		fmt.Println(dockercmd)
		cmd := [...]string{"bash", "-c", dockercmd}
		fmt.Println(ExecCommand(cmd[:]...))
	}
}
