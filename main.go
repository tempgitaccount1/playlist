package main

import (
	"fmt"

	"github.com/tempgitaccount1/playlist/server"
)

func main() {
	server.DoServerStuff()
	g := server.New()
	g.Start()
	fmt.Println("Hello, world!")
}
