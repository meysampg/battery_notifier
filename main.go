package main

import (
	"flag"
	"fmt"
)

var debug bool
var path string
var threshold uint

func init() {
	flag.BoolVar(&debug, "d", false, "enable debug mode")
	flag.StringVar(&path, "p", "/org/freedesktop/UPower/devices/battery_BAT1", "path of battery upower battery device (run upower -e)")
	flag.UintVar(&threshold, "t", 20, "threshold of critical situation")
}

func main() {
	flag.Parse()

	percentage, err := batteryPercentage(path)
	if err != nil && debug {
		fmt.Println(fmt.Errorf("an error on generating percentage occured: %s", err.Error()))
	}

	err = notification(percentage)
	if err != nil && debug {
		fmt.Println(fmt.Errorf("an error on displaying notification occured: %s", err.Error()))
	}
}
