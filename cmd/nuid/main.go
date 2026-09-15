package main

import (
	"fmt"
	"os"

	"github.com/nats-io/nuid"
)

const usage = `Usage:
  nuid               print a new NUID
  nuid parse <nuid>  check that <nuid> is a valid NUID
  nuid help          show this help

Aliases: p = parse; -h, --help = help
`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(nuid.Next())
		return
	}
	switch os.Args[1] {
	case "help", "-h", "--help":
		fmt.Print(usage)
	case "parse", "p":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "usage: nuid parse <nuid>")
			os.Exit(1)
		}
		_, err := nuid.Parse(os.Args[2])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("ok")
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n\n%s", os.Args[1], usage)
		os.Exit(1)
	}
}
