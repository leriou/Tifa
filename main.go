package main

import (
	"Tifa/manager"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	root := flag.Arg(0)
	fm := manager.NewFm()
	c := fm.Tool.GetTimeStamp()
	fmt.Println(c)
	// fm.Remove("dd")
	fm.Scan(root)

	fmt.Println(fm.Tool.GetTimeStamp() - c)
	//root = "/Users/lixiumeng/data"
	//fm.Apply(root)
	//fm.ClearPath(root)
}
