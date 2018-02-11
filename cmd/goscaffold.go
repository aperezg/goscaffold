package main

import (
	"github.com/aperezg/goscaffold"
)

func main() {
	s := goscaffold.Client.Stdin()
	g := goscaffold.NewGenerator(s)
	g.Scaffold()
}
