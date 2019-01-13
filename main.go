package main

import (
	"flag"
	"fmt"

	"github.com/rsjethani/sysinfo"
)

var debug bool
var path string
var threshold uint
var notif bool
var notificon bool

func init() {
	flag.BoolVar(&debug, "d", false, "enable debug mode")
	flag.UintVar(&threshold, "t", 20, "threshold of critical situation")
	flag.BoolVar(&notif, "n", false, "enable send notification")
	flag.BoolVar(&notificon, "i", false, "enable send notification with icon")
}

func main() {
	flag.Parse()

	info, err := sysinfo.GetInfo("hardware", "battery")

	if err != nil && debug {
		fmt.Println(fmt.Errorf("an error on get battery info occured: %s", err.Error()))
	}

	if notif {
		c, _ := info.Attribute(0, "CAPACITY")
		cap, _ := c.(uint)
		s, _ := info.Attribute(0, "STATUS")
		state, _ := s.(string)

		err = notification(cap, state, notificon)
		if err != nil && debug {
			fmt.Println(fmt.Errorf("an error on displaying notification occured: %s", err.Error()))
		}
	}
}
