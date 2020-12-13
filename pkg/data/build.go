package data

import (
	"fmt"
)

const Binary = "badpod"

var (
	Version   string
	GitCommit string
	BuildTime string
)

func getBuildData() map[string]string {
	data := map[string]string{}

	data["Version"] = Version
	data["GitCommit"] = GitCommit
	data["BuildTime"] = BuildTime

	return data
}

// For logging
func RenderBuildData() string {
	return fmt.Sprintf("%s %s: git %s, built at %s", Binary, Version, GitCommit, BuildTime)
}
