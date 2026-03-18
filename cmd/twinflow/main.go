package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Saad7890-web/twinflow/cli"
)

var version = "v0.1.0"

func main() {

	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println("TwinFlow version:", version)
		return
	}
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}