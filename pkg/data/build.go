package data

import (
	"fmt"
	"runtime"
)

const Binary = "badpod"

var (
	Version   string
	BuildTime string
)

func getBuildData() map[string]string {
	data := map[string]string{}

	data["Version"] = Version
	data["BuildTime"] = BuildTime

	return data
}

// For logging
func RenderBuildData() string {
	return fmt.Sprintf("%s %s, built at %s with %s", Binary, Version, BuildTime, runtime.Version())
}
