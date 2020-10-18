package main

import (
	"adscale-tools/web"
	"fmt"
	"log"
	"os"
)

func main() {
	cors := os.Getenv("profile") == "prod"
	app := web.NewApp(!cors)
	fmt.Println(app)
	port := "8088"
	if err := app.Serve(port); err != nil {
		log.Println("Error", err)
	}
}
