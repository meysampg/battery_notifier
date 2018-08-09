package main

import (
	"os/exec"
	"fmt"
)

func notification(percentage uint) error {
	cmd := exec.Command("notify-send", "-u", "critical", "-t", "5000", notificationTitle(percentage), notificationSummary(percentage))
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

func notificationSummary(percentage uint) string  {
	switch {
	case percentage <= threshold:
		return "Please Connect Charger to Device"
	default:
		return ""
	}
}