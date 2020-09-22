package main

import (
	"fmt"

	"github.com/tempgitaccount1/playlist/server"
)

func main() {
	g := server.New()
	g.Start()
	fmt.Println("Hello, world!")
}
