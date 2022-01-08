package main

import (
	"fmt"
	"log"
	"os"

	"github.com/puoklam/gohelper/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
