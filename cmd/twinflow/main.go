package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("twinflow: a traffic capture and replay tool")
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Println("  twinflow record   Start traffic recorder")
		fmt.Println("  twinflow replay   Replay captured traffic")
		fmt.Println("  twinflow analyze  Analyze results")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "record":
		fmt.Println("Starting record mode...")
	case "replay":
		fmt.Println("Starting replay mode...")
	case "analyze":
		fmt.Println("Starting analyze mode...")
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}