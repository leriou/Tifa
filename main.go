package main

import (
	"Tifa/manager"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	root := flag.Arg(1)
	comm := flag.Arg(0)
	fm := manager.NewFm()
	c := fm.Tool.GetTimeStamp()
	switch comm {
		case "scan":
			fm.Scan(root)
			break
		case "remove":
			fm.Remove(root)
			break
		case "apply":
			fm.Apply(root)
			break
		case "clear":
			fm.ClearPath(root)
			break
	}
	fmt.Println(fm.Tool.GetTimeStamp() - c)
}
