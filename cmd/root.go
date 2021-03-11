package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected 'generate' or 'sync' subcommands")
		os.Exit(1)
	}

	var err error
	switch os.Args[1] {
	case "generate":
		generate()
	case "sync":
		err = sync()
	default:
		err = fmt.Errorf("subcommand %q not supported", os.Args[1])
	}
	if err != nil {
		log.Fatal(err)
	}
}
