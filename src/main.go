package main

import (
	"flag"
	"sephiroth/utils"
	"strconv"
	"tifa/src/manager"
)

func main() {
	flag.Parse()
	comm := flag.Arg(0)
	root := flag.Arg(1)
	fm := manager.NewFm()
	tu := utils.NewTimeUtil()
	c := tu.GetTimeStamp()
	logger := utils.NewLogger()
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
	logger.Info(strconv.FormatInt((tu.GetTimeStamp()-c)/1000, 10))
}
