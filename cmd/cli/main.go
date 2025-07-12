package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	helloCmd := flag.NewFlagSet("hello", flag.ExitOnError)
	name := helloCmd.String("name", "world", "the name to greet")

	if len(os.Args) < 2 {
		fmt.Println("expected 'hello' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "hello":
		if err := helloCmd.Parse(os.Args[2:]); err != nil {
			fmt.Printf("error parsing flags: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Hello, %s!\n", *name)
	default:
		fmt.Println("expected 'hello' subcommand")
		os.Exit(1)
	}
}
