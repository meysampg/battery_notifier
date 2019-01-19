package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/esiqveland/notify"
	"github.com/godbus/dbus"
	"github.com/rsjethani/sysinfo"
)

const appVersion = "v2.1.0"

var argWatch bool
var argThreshold uint
var argShowVersion bool
var argLowInterval time.Duration
var argNormalInterval time.Duration

func init() {
	flag.BoolVar(&argShowVersion, "v", false, "Show application version")
	flag.DurationVar(&argLowInterval, "l", time.Minute*2, "battery check interval during low (< threshold) battery")
	flag.DurationVar(&argNormalInterval, "n", time.Minute*5, "battery check interval during good/normal (> threshold) battery")
	flag.BoolVar(&argWatch, "w", false, "continuously watch battery level at preset interval. The interval depends on values of '-n' and '-l'")
	flag.UintVar(&argThreshold, "t", 20, "battery percentage threshold, below which the battery will be condiered as *low* and the user will start getting desktop notifications about low battery.")
}

func getBatteryStatus() (uint, string, error) {
	info, err := sysinfo.GetInfo("hardware", "battery")
	if err != nil {
		return 0, "", err
	}
	c, _ := info.Attribute(0, "CAPACITY")
	capacity, _ := c.(uint)
	s, _ := info.Attribute(0, "STATUS")
	state, _ := s.(string)

	return capacity, state, nil
}

func sendNotification(percentage uint, chargingState string) error {
	conn, err := dbus.SessionBus()
	if err != nil {
		return err
	}

	urgency := 1 // normal
	icon := "battery-low"
	if percentage <= 5 {
		urgency = 2 //critical
		icon = "battery-empty"
	}

	n := notify.Notification{
		AppName:       "Battery Notifier",
		ReplacesID:    0,
		AppIcon:       icon,
		Summary:       fmt.Sprintf("Only %v%% battery remaining !!!", percentage),
		Hints:         map[string]dbus.Variant{"urgency": dbus.MakeVariant(urgency)},
		ExpireTimeout: int32(3000),
	}

	_, err = notify.SendNotification(conn, n)
	return err
}

func main() {
	flag.Parse()

	if argShowVersion {
		fmt.Println("battery_notifier", appVersion)
		return
	}

	for {
		capacity, state, err := getBatteryStatus()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Current Capacity: %v%%, Current State: %v\n", capacity, state)

		// If -w not given, exit
		if !argWatch {
			break
		}

		sleepInterval := argNormalInterval
		if state == "Discharging" && capacity < argThreshold {
			sleepInterval = argLowInterval
			err = sendNotification(capacity, state)
			if err != nil {
				fmt.Println("Error while sending desktop notification: %v", err)
				os.Exit(2)
			}
		}
		time.Sleep(sleepInterval)
	}
}
