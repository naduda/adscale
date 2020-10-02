package main

import (
	"adscale-tools/config"
	"adscale-tools/repo"
	"adscale-tools/web"
	"flag"
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

func main2() {
	//repoPath := "/Users/pr/projects/git.public"
	//configPath := "/Users/Shared/configuration/easyleads.conf"
	repoPath := flag.String("repo", "/Users/pr/projects/git.public", "set your repo path")
	configPath := flag.String("config", "/Users/Shared/configuration/easyleads.conf", "select your easyleads.conf")
	flag.Parse()

	config := config.Config{}
	config.Init(*configPath)
	config.Format()

	repo := repo.Repository{}
	repo.Init(*repoPath)

	repo.CheckProperties(&config)

	for k, v := range config.Props {
		if !v.Status {
			fmt.Println(k)
		}
	}
}
