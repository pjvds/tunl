package version

import "fmt"

var (
	Version   = "<unknown>"
	BuildDate = ""
	GitCommit = ""
)

// String returns the version string.
func String() string {
	version := Version
	if GitCommit != "" && BuildDate != "" {
		version += fmt.Sprintf(" (%s, %s)", GitCommit, BuildDate)
	}
	return version
}
