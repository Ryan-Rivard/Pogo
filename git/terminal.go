package git

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
)

func ExecuteGitCommand(args ...string) {
	cmd := exec.Command("git", args...)

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	// Execute the command
	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}

	// fmt.Println(stdBuffer.String())

	// cmd := exec.Command(program, args...)
	// stdout, err := cmd.Output()

	// if err != nil {
	// 	log.Panic(err.Error())
	// }

	// return string(stdout)
}
