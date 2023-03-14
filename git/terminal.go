package git

import (
	"log"
	"os/exec"
)

func ExecuteGitCommand(args ...string) string {
	cmd := exec.Command("git", args...)

	println(cmd.String())
	stdout, err := cmd.Output()

	if err != nil {
		log.Panic(err.Error())
	}

	output := string(stdout)
	println("Output: " + output)

	return output
}
