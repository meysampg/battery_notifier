package main

import (
	"regexp"
	"os/exec"
	"fmt"
	"strconv"
)

func batteryPercentage(info string) (uint, error) {
	// extract percentage of battery
	percentage, err := extractBatteryPercentage(info)
	if err != nil {
		return 0, err
	}

	// convert it to integer
	p, err := strconv.Atoi(percentage)
	if err != nil {
		return 0, err
	}

	// and finally return it on proper type
	return uint(p), nil
}

func batteryChargingState(info string) (uint, error) {
	// extract charging state of battery
	state, err := extractBatteryChargingState(info)
	if err != nil {
		return 0, err
	}

	if state == "charging" {
		return 1, nil
	} else {
		return 0, nil
	}
}

func getBatteryInfo(path string) (string, error) {
	// get battery information
	out, err := exec.Command("upower", "-i", path).Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func extractBatteryPercentage(info string) (string, error) {
	var re = regexp.MustCompile(`(?i)percentage:\s+(\d{1,3})%`)
	var result = re.FindStringSubmatch(info)

	if len(result) < 2 {
		return "", fmt.Errorf("can't extract percentage of battery")
	}

	return result[1], nil
}

func extractBatteryChargingState(info string) (string, error) {
	var re = regexp.MustCompile(`(?i)state:\s+(\w+)`)
	var result = re.FindStringSubmatch(info)

	if len(result) < 2 {
		return "", fmt.Errorf("can't extract charging state of battery")
	}
	
	return result[1], nil
}