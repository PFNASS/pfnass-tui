package main

import (
	"log"

	"github.com/PFNASS/pfnass-tui/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
