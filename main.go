package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rsjethani/sysinfo"
)

var threshold uint
var notif bool

func main() {
	flag.UintVar(&threshold, "t", 20, "threshold of critical situation")
	flag.BoolVar(&notif, "n", false, "enable send notification")
	flag.Parse()

	info, err := sysinfo.GetInfo("hardware", "battery")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c, _ := info.Attribute(0, "CAPACITY")
	capacity, _ := c.(uint)
	s, _ := info.Attribute(0, "STATUS")
	state, _ := s.(string)

	fmt.Println(capacity, state)

	if notif {
		err = sendNotification(capacity, state)
		if err != nil {
			fmt.Println("an error on displaying notification occured: %s", err)
			os.Exit(2)
		}
	}
}
