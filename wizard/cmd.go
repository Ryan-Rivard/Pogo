package wizard

import (
	"log"
	"os/exec"
)

// Composite Pattern
// Leaf Component
// Does the actual work

type Cmd struct {
	arg string
}

func (w *Cmd) Execute() {
	cmd := exec.Command("git", w.arg)
	stdout, err := cmd.Output()

	if err != nil {
		log.Panic(err.Error())
	}

	println(string(stdout))
}
