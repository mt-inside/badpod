package renderers

import (
	"strings"

	"github.com/mt-inside/badpod/pkg/data"
)

const lorumIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

func bufferOutput(n int) string {
	var s strings.Builder

	ch := make(chan string)
	go makeOutput(n, ch)

	for part := range ch {
		s.WriteString(part)
	}

	return s.String()
}

func makeOutput(n int, ch chan<- string) {
	/* Really just want pointers to these but go makes that difficult */
	ss := []string{
		data.RenderBuildData() + "\n",
		data.RenderSessionData() + "\n",
		lorumIpsum,
	}
	i := 0

	/* TODO deal with -1 == inf */
	for {
		if n <= len(ss[i]) {
			ch <- ss[i][:n]
			close(ch)
			return
		} else {
			ch <- ss[i][:]
			n = n - len(ss[i])

			if i != len(ss)-1 {
				i = i + 1
			}
		}
	}
}

func RenderLorumIpsum(_ map[string]string) (bs []byte) {
	//return []byte(bufferOutput (len(lorumIpsum) * 2 /*TODO*/))
	/* FIXME more flushing */
	// TODO set http response content-length header (except if len is inf)
	return []byte(bufferOutput(100 /*TODO*/))
}
