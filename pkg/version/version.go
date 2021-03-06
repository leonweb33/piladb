// Package version provides ways to get the version of the
// program.
package version

import (
	"os/exec"
	"strings"
)

// VERSION defines piladb version
const VERSION = "0.1.0"

// Version returns piladb version given a v version. If v is empty,
// defaults to CommitHash.
func Version(v string) string {
	if v == "" {
		return CommitHash()
	}

	return v
}

// CommitHash returns the commit hash of the repository.
func CommitHash() string {
	cmd := exec.Command("git", []string{"rev-parse", "HEAD"}...)

	b, err := cmd.Output()
	if err != nil {
		return "master"
	}
	return strings.Replace(string(b), "\n", "", -1)
}
