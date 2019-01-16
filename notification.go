package main

import (
	"fmt"

	"github.com/esiqveland/notify"
	"github.com/godbus/dbus"
)

func sendNotification(percentage uint, chargingState string) error {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	urgency := 1 // normal
	if percentage < 5 {
		urgency = 2 //critical
	}

	n := notify.Notification{
		AppName:       "Battery Notifier",
		ReplacesID:    0,
		AppIcon:       notificationIcon(percentage, chargingState),
		Body:          notificationSummary(percentage),
		Summary:       notificationTitle(percentage),
		Hints:         map[string]dbus.Variant{"urgency": dbus.MakeVariant(urgency)},
		ExpireTimeout: int32(3000),
	}

	_, err = notify.SendNotification(conn, n)
	return err
}

func notificationTitle(percentage uint) string {
	return fmt.Sprintf("Battery %s (%d%%)", batteryAdjective(percentage), percentage)
}

func batteryAdjective(percentage uint) string {
	switch {
	case percentage <= argThreshold:
		return "LOW"
	case percentage <= 50:
		return "is OK"
	case percentage <= 70:
		return "is fine"
	default:
		return "kisses you"
	}
}

func notificationSummary(percentage uint) string {
	switch {
	case percentage <= argThreshold:
		return "Please Connect Charger to Device"
	default:
		return ""
	}
}

func notificationIcon(percentage uint, chargingState string) string {
	switch {
	case chargingState == "Charging":
		return chargingIcon(percentage)
	case chargingState == "Discharging":
		return nonChargingIcon(percentage)
	default:
		return ""
	}
}

func chargingIcon(percentage uint) string {
	switch {
	case percentage <= 1:
		return "battery-empty"
	case percentage <= argThreshold:
		return "battery-low-charging"
	case percentage == 100:
		return "battery-full-charged"
	default:
		return "battery-good-charging"
	}
}

func nonChargingIcon(percentage uint) string {
	switch {
	case percentage <= 1:
		return "battery-empty"
	case percentage <= argThreshold:
		return "battery-low"
	case percentage == 100:
		return "battery-full"
	default:
		return "battery-good"
	}
}
