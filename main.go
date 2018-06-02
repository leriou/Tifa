package main

import (
	"filetarven/manager"
	"flag"
	// "fmt"
)

func main() {
	flag.Parse()
	root := flag.Arg(0)
	fm := manager.NewFm()
	//fm.Scan(root)
	root = "/Users/lixiumeng/data"
	fm.Apply(root)
}
