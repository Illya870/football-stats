package main

import (
	"football-stats/cli"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		cli.PrintHelp()
		os.Exit(2)
	}
	cli.Cli(os.Args[1:])
}
