package main

import (
	"flag"
	"fmt"
)

var debug bool
var path string
var threshold uint
var notif bool
var notificon bool

func init() {
	flag.BoolVar(&debug, "d", false, "enable debug mode")
	flag.StringVar(&path, "p", "/org/freedesktop/UPower/devices/battery_BAT1", "path of battery upower battery device (run upower -e)")
	flag.UintVar(&threshold, "t", 20, "threshold of critical situation")
	flag.BoolVar(&notif, "n", false, "enable send notification")
	flag.BoolVar(&notificon, "i", false, "enable send notification with icon")
}

func main() {
	flag.Parse()

	info, err := getBatteryInfo(path)
	if err != nil && debug {
		fmt.Println(fmt.Errorf("an error on get battery info occured: %s", err.Error()))
	}

	percentage, err := batteryPercentage(info)
	if err != nil && debug {
		fmt.Println(fmt.Errorf("an error on generating percentage occured: %s", err.Error()))
	}

	chargingState, err := batteryChargingState(info)
	if err != nil && debug {
		fmt.Println(fmt.Errorf("an error on generating charging state occured: %s", err.Error()))
	}

	if notif {
		if notificon {
			err = notification(percentage, chargingState, true)
		} else {
			err = notification(percentage, chargingState, false)
		}
		if err != nil && debug {
			fmt.Println(fmt.Errorf("an error on displaying notification occured: %s", err.Error()))
		}
	}
}
