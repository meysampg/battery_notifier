package main

import (
	"fmt"
	"os/exec"
)

func notification(percentage uint, chargintState string, withIcon bool) error {
	command := "notify-send"
	var args []string
	if withIcon {
		args = []string{"-u", "critical", "-t", "5000", notificationTitle(percentage), notificationSummary(percentage), "-i", notificationIcon(percentage, chargintState)}
	} else {
		args = []string{"-u", "critical", "-t", "5000", notificationTitle(percentage), notificationSummary(percentage)}
	}
	cmd := exec.Command(command, args...)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func notificationTitle(percentage uint) string {
	return fmt.Sprintf("Battery %s (%d%%)", batteryAdjective(percentage), percentage)
}

func batteryAdjective(percentage uint) string {
	switch {
	case percentage <= threshold:
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
	case percentage <= threshold:
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
	case percentage <= threshold:
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
	case percentage <= threshold:
		return "battery-low"
	case percentage == 100:
		return "battery-full"
	default:
		return "battery-good"
	}
}
