package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rsjethani/sysinfo"
)

var argWatch bool
var argThreshold uint
var argLowInterval time.Duration
var argNormalInterval time.Duration

func init() {
	flag.BoolVar(&argWatch, "w", false, "continously watch battery status")
	flag.DurationVar(&argLowInterval, "l", time.Minute*2, "battery check interval during low battery")
	flag.DurationVar(&argNormalInterval, "n", time.Minute*10, "battery check interval during good battery")
	flag.UintVar(&argThreshold, "t", 20, "threshold below which battery capacty would be considered critical")
}

func main() {
	flag.Parse()

	for {

		info, err := sysinfo.GetInfo("hardware", "battery")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		c, _ := info.Attribute(0, "CAPACITY")
		capacity, _ := c.(uint)
		s, _ := info.Attribute(0, "STATUS")
		state, _ := s.(string)

		log.Println(capacity, state)

		if argWatch {
			sleepInterval := argNormalInterval
			if state == "Discharging" && capacity < argThreshold {
				sleepInterval = argLowInterval
				err = sendNotification(capacity, state)
				if err != nil {
					fmt.Println("an error on displaying notification occured: %s", err)
					os.Exit(2)
				}
			}
			time.Sleep(sleepInterval)
			continue
		}

		break
	}
}
