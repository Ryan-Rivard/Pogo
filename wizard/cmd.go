package wizard

import (
	"log"
	"os/exec"
)

// Composite Pattern
// Leaf Component
// Does the actual work

type Cmd struct {
	Id  string
	arg []string
}

func (c *Cmd) GetId() string {
	return c.Id
}

func (c *Cmd) Execute() {
	cmd := exec.Command("git", c.arg...)
	stdout, err := cmd.Output()

	if err != nil {
		log.Panic(err.Error())
	}

	println(string(stdout))
}
