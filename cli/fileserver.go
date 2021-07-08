package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			fmt.Println(`---Network file server---
This program is meant to pass a file 
between two devices over the network.
Usage:
./fileserver -f <filename> -p <port_to_listen>`)
			return
		}
	}
}
