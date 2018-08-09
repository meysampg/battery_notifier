package main

import (
	"regexp"
	"os/exec"
	"fmt"
	"strconv"
)

func batteryPercentage(path string) (uint, error) {
	// first get information of battery
	out, err := getBatteryInfo(path)
	if err != nil {
		return 0, err
	}

	// then extract percentage of battery
	percentage, err := extractBatteryInfo(out)
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

func getBatteryInfo(path string) (string, error) {
	out, err := exec.Command("upower", "-i", path).Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func extractBatteryInfo(info string) (string, error) {
	var re = regexp.MustCompile(`(?i)percentage:\s+(\d{1,3})%`)
	var result = re.FindStringSubmatch(info)

	if len(result) < 2 {
		return "", fmt.Errorf("can't extract percentage of battery")
	}

	return result[1], nil
}