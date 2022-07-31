package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
)

func Execute(gitCmd ...string) {
	cmd := exec.Command("git", gitCmd...)

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	// Execute the command
	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}

	// automatically see it on the terminal
	// log.Println(stdBuffer.String())
}
