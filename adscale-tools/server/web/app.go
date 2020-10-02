package web

import (
	"adscale-tools/web/handlers"
	"fmt"
	"log"
	"net/http"
)

type App struct {
	handlers map[string]http.HandlerFunc
}

func NewApp(disableCors bool) App {
	return App{
		handlers: handlers.GetHandlers("api", disableCors),
	}
}

func (a *App) Serve(port string) error {
	for path, handler := range a.handlers {
		http.Handle(path, handler)
	}
	log.Printf("Web server is available on port %s\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
