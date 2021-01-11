package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	cli "github.com/jawher/mow.cli"
	"github.com/mt-inside/badpod/pkg/data"
	"github.com/mt-inside/badpod/pkg/handlers"
	"github.com/mt-inside/badpod/pkg/middleware"
	"github.com/mt-inside/badpod/pkg/renderers"
)

func main() {
	log.Println(data.RenderBuildData())
	log.Println(data.RenderSessionData())

	app := cli.App(data.Binary, "A programme that behaves, sometimes, badly")
	app.Spec = "[ADDR]"
	addr := app.StringArg("ADDR", ":8080", "Listen address")

	app.Action = func() { appMain(addr) }

	app.Run(os.Args)
}

func appMain(addr *string) {
	rootMux := mux.NewRouter()
	rootMux.Use(middleware.LoggingMiddleware)

	handlers.HandleApi(rootMux.PathPrefix("/api").Subrouter()) //TODO rename our package away from handlers
	handlers.HandleMisc(rootMux)
	handlers.HandleProbes(rootMux)

	rootMux.Path("/").Handler(middleware.MiddlewareStack(renderers.RenderLorumIpsum, "text/plain"))

	log.Printf("Listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, rootMux))

	// TODO: graceful shutdown (lower readiness)
}
