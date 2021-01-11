package middleware

import (
	"log"
	"net/http"

	"github.com/mt-inside/badpod/pkg/data"
	"github.com/mxk/go-flowrate/flowrate"
)

type slowResponseWriter struct {
	rw    http.ResponseWriter
	fw    *flowrate.Writer
	oldBw int64
}

func newSlowResponseWriter(rw http.ResponseWriter) slowResponseWriter {
	bandwidth := data.GetBandwidth() //TODO lazy in Write()
	fw := flowrate.NewWriter(rw, bandwidth)
	return slowResponseWriter{rw, fw, bandwidth}
	//defer fw.Close()
}

func (sr slowResponseWriter) Header() http.Header {
	return sr.rw.Header()
}

func (sr slowResponseWriter) Write(b []byte) (written int, err error) {
	if sr.oldBw != data.GetBandwidth() {
		// FIXME: not thread safe
		sr.oldBw = data.GetBandwidth()
		sr.fw.SetLimit(sr.oldBw)
		log.Println("adjusted writer bw to ", sr.oldBw)
	}
	written, err = sr.fw.Write(b)
	sr.rw.(http.Flusher).Flush() // TODO defer?
	return
}

func (sr slowResponseWriter) WriteHeader(statusCode int) {
	sr.rw.WriteHeader(statusCode)
}

func bandwidthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(newSlowResponseWriter(w), r)
	})
}
