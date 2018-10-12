package main

import (
	"fmt"
	"os"
	"strconv"
)

// Temp levels are celcius x 1000
const (
	fanPrefix = "/sys/devices/platform/applesmc.768/fan1_"
	fanManual = fanPrefix + "manual"
	fanSet    = fanPrefix + "output"
)

// SetManual switches the fan to manual control if true; automatic, otherwise
func SetManual(set bool) error {
	f, err := os.Create(fanManual)
	if err != nil {
		return err
	}
	defer f.Close()

	val := "0"
	if set {
		val = "1"
	}
	_, err = f.WriteString(val)
	if err != nil {
		return err
	}

	return nil
}

// SetByState sets the fan speed appropriate for the state
func SetByState(s State) error {
	if s.Speed < stateLo.Speed {
		return fmt.Errorf("Speed %d is lower than allowed (%d)",
			s.Speed, stateLo.Speed)
	}

	val := strconv.Itoa(s.Speed)
	f, err := os.Create(fanSet)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(val)
	if err != nil {
		return err
	}

	return nil
}
