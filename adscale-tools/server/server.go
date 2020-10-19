package main

import (
	"adscale-tools/web"
	"flag"
	"log"
	"os"
)

func main() {
	port := flag.Int("p", 8088, "dev server port")
	flag.Parse()

	cors := os.Getenv("profile") == "prod"
	app := web.NewApp(!cors)

	if err := app.Serve(*port); err != nil {
		log.Println("Error", err)
	}
}
