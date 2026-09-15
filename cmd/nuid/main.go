package main

import (
	"fmt"
	"os"

	"github.com/nats-io/nuid"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "parse" || os.Args[1] == "p") {
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
		return
	}
	fmt.Println(nuid.Next())
}
