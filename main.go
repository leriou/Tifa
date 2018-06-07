package main

import (
	"Tifa/manager"
	"flag"
	// "fmt"
)

func main() {
	flag.Parse()
	root := flag.Arg(0)
	fm := manager.NewFm()
	// fm.Remove("dd")
	fm.Scan(root)
	//root = "/Users/lixiumeng/data"
	//fm.Apply(root)
	//fm.ClearPath(root)
}
