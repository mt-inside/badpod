package data

import (
	"net/http"

	"github.com/mt-inside/badpod/pkg/util"
)

func GetData(r *http.Request) map[string]string {
	d := make(map[string]string) //TODO: strongly type me with a struct. Esp for (optional) sections

	d = util.AppendMap(d, getBuildData())
	d = util.AppendMap(d, getSessionData())
	d = util.AppendMap(d, getSettingsData())
	d = util.AppendMap(d, getProcData())

	return d
}
