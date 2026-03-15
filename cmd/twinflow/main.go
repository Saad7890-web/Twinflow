package main

import (
	"log"

	"github.com/Saad7890-web/twinflow/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}