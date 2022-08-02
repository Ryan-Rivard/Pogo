package git

import (
	"strings"

	"github.com/Ryan-Rivard/pogo/git/terminal"
)

func ListChangedFiles() []string {
	output := terminal.Execute("git", "ls-files", "-m", "-d", "-o", "--exclude-standard")

	return strings.Split(strings.TrimSpace(output), "\n")
}
