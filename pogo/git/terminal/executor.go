package terminal

import (
	"log"
	"os/exec"
)

func Execute(program string, args ...string) string {
	cmd := exec.Command(program, args...)
	stdout, err := cmd.Output()

	if err != nil {
		log.Panic(err.Error())
	}

	return string(stdout)
}
