package wizard

import (
	"log"
	"os/exec"
	"strings"
)

type cmd struct {
	id           string
	args         []string
	formatOutput func([]byte) []string
	step         step
}

func (c *cmd) Exec(params []string) {
	gitParams := append(c.args, params...)
	cmd := exec.Command("git", gitParams...)

	stdout, err := cmd.Output()

	if err != nil {
		log.Panic(err.Error())
	}

	if c.formatOutput != nil {
		c.step.Exec(c.formatOutput(stdout))
	}

	c.step.Exec(nil)
}

func branchListFormat(output []byte) []string {
	return strings.Split(strings.TrimSpace(string(output)), "\n")
}
