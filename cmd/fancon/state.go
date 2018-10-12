package main

// Temperatures = 1000 per degree Celcius
const delta = 4 * 1000

// State is defined by a range of temperatures and an associated fan speed
type State struct {
	Lo    Temp
	Hi    Temp
	Delta Temp
	Speed int
}

// Lower reports the lowest temperature for a state to prevent noise from
// constantly throttling the temperature
func (s State) Lower() Temp {
	return s.Lo - delta
}

// NextState determines the next state we should be in, according to the
// temperature
func (s State) NextState(t Temp) State {
	var prev, next State

	switch s {
	case stateLo:
		prev = stateLo
		next = stateMed

	case stateMed:
		prev = stateLo
		next = stateHi

	case stateHi:
		prev = stateMed
		next = stateHi

	default:
		return stateHi
	}

	if t > s.Hi {
		return next
	} else if t < s.Lower() {
		return prev
	}

	return s
}

// GetState gets the state based on the temperature
func GetState(t Temp) State {
	if t < stateMed.Lo {
		return stateLo
	}
	if t < stateHi.Lo {
		return stateMed
	}
	return stateHi
}

func (s State) String() string {
	switch s {
	case stateLo:
		return "lo"
	case stateMed:
		return "med"
	case stateHi:
		return "hi"
	}
	return "unknown"
}
