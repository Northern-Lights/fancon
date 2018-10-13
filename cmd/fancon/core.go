package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// Temp is an int representing a temperature
type Temp int

// GetCoreTemp returns the greatest temperature over all the cores
func GetCoreTemp() Temp {
	var max, n int
	for _, path := range corePaths {
		var (
			buf  [10]byte
			temp = 99999 // In case something goes wrong, as a safety
		)
		f, err := os.Open(path)
		if err != nil {
			log.Printf("Couldn't open temp for %s: %v", path, err)
			goto end
		}
		defer f.Close()

		n, err = f.Read(buf[:])
		if err != nil {
			log.Printf("Couldn't read temp for %s: %v", path, err)
			goto end
		}

		temp, err = strconv.Atoi(strings.TrimSpace(string(buf[:n])))
		if err != nil {
			log.Printf("Couldn't convert temp to int for %s: %v", path, err)
			goto end
		}

		if temp > max {
			max = temp
		}
	end:
	}

	if max < int(stateLo.Lo) {
		max = 99999
	}

	return Temp(max)
}
