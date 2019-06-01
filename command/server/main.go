package main

import (
	"fmt"

	"github.com/eiichiroi/ping"
)

func main() {
	fmt.Fprintf("Server! " + ping.Version + "\n")
}
