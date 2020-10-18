package web

import (
	"adscale-tools/docker"
	"adscale-tools/model"
	"adscale-tools/web/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
)

type App struct {
	handlers map[string]http.HandlerFunc
}

func NewApp(disableCors bool) App {
	var settings model.Settings
	if err := settings.Read(); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Settings not exists...")
		} else {
			panic(err)
		}
	}

	var d docker.AdscaleContainer
	if err := d.Init(); err != nil {
		panic(err)
	}

	h := handlers.Handlers{
		Settings: settings,
		Docker:   d,
	}

	return App{
		handlers: h.GetHandlers("api", disableCors),
	}
}

func (a *App) Serve(port string) error {
	for path, handler := range a.handlers {
		http.Handle(path, handler)
	}

	log.Printf("Web server is available on port %s\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
