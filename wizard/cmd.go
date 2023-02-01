package wizard

import (
	"log"
	"os/exec"
	"strings"
)

type cmd struct {
	id            string
	args          []string
	convertInput  func(interface{}) []string
	convertOutput func([]byte) interface{}
	step          step
}

func (c *cmd) Exec(input interface{}) {
	params := []string{}
	if c.convertInput != nil {
		params = c.convertInput(input)
	}
	gitParams := append(c.args, params...)
	cmd := exec.Command("git", gitParams...)

	println(cmd.String())

	stdout, err := cmd.Output()

	if err != nil {
		log.Panic(err.Error())
	}

	if c.step == nil {
		return
	}

	if c.convertOutput != nil {
		output := c.convertOutput(stdout)
		c.step.Exec(output)
	} else {
		c.step.Exec(nil)
	}
}

func importStringSlice(input interface{}) []string {
	b, ok := input.([]string)
	if !ok {
		log.Fatal("boomie")
	}
	return b
}

// func importBranchList(input interface{}) []string {
// 	b, ok := input.(*branch)
// 	if !ok {
// 		log.Fatal("boom")
// 	}
// 	return []string{b.refname}
// }

func exportBranchList(output []byte) interface{} {
	s := strings.Split(strings.TrimSpace(string(output)), "\n")
	b := []branch{}

	for _, a := range s {
		b = append(b, branch{refname: a})
	}

	return &branches{branches: b}
}
