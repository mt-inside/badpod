package middleware

import (
	"net/http"

	"github.com/mt-inside/badpod/pkg/data"
)

// TODO next should be an http.Handler?
func MiddlewareStack(next func(map[string]string) []byte, mime string) http.Handler {
	//return /*recoveryMiddleware( */
	/* errorMiddleware( */
	/* latencyMiddleware( */
	return bandwidthMiddleware(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", mime)

			//TODO: wind envbin back to the point this actually worked...
			bs := next(data.GetData(r)) // TODO: refactor. This really shouldn't be here. next should be passed in partially applied to the data? and as a Handler[Func]?

			// Templates can be executed straight into writers, so we could pump the template into the httpResponseWriter. Problem is, it only flushes on the boundaries into and out of {{}} template substitutions, which makes the output sporadic. So we dump into a string and write that one byte at a time.
			// TODO: make the chan here, pass in, get a string at a time to avoid filling memory for inf body len
			for i := 0; i < len(bs); i++ {
				w.Write(bs[i : i+1])
			}
		}),
	)
	//)
	//),
	//)
}
