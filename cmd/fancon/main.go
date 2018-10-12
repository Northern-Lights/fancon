package main

import (
	"log"
	"math"
	"time"
)

var (
	stateLo = State{
		Lo:    0,
		Hi:    stateMed.Lo - 1,
		Speed: 2000, // TODO: use fanX_min
	}
	stateMed = State{
		Lo:    60000,
		Hi:    stateHi.Lo - 1,
		Speed: 4000, // TODO: use avg of min & max
	}
	stateHi = State{
		Lo:    70000,
		Hi:    Temp(math.MaxInt32),
		Speed: 6200, // TODO: use fanX_max
	}
)

var (
	corePaths = []string{
		"/sys/devices/platform/applesmc.768/temp9_input",
		"/sys/devices/platform/applesmc.768/temp10_input",
	}
)

func run() {

	// state := GetState(GetCoreTemp())
	state := stateHi
	log.Println("Current state:", state)
	err := SetByState(state)
	if err != nil {
		log.Fatalln("Couldn't set initial fan speed:", err)
	}

	// TODO: stop this on interrupt
	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {
		temp := GetCoreTemp()
		log.Println("Current temp:", temp)
		next := state.NextState(temp)
		log.Println("Next state:", next)
		if next != state {
			log.Println("Setting fan due to state change")
			err := SetByState(next)
			if err != nil {
				log.Println("Unable to set fan speed:", err)
			}
			state = next
		}
	}
}

func main() {
	log.Println("Setting fan control to manual mode...")
	err := SetManual(true)
	if err != nil {
		log.Fatalf("Couldn't take manual control of fan %s: %v", fanManual, err)
	}
	defer SetManual(false) // TODO: really need to log or something if failure

	log.Println("Starting fan control...")
	run()
}
