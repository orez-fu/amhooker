package main

import (
	"amhooker/cli"
	"log"
	"os"
)

func main() {
	log.Println("Start Alert Manager Hooker...")

	if err := cli.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
