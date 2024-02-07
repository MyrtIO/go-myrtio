package serial

import (
	"go.bug.st/serial/enumerator"
)

func findCandidates(ports []*enumerator.PortDetails) []string {
	highPriority := make([]string, 0)
	lowPriority := make([]string, 0)
	for _, port := range ports {
		if !port.IsUSB {
			continue
		}
		if isKnownDevice(port.VID + ":" + port.PID) {
			highPriority = append(highPriority, port.Name)
			continue
		}
		if isKnownDevice(port.VID) {
			lowPriority = append(lowPriority, port.Name)
			continue
		}
	}
	highPriority = append(highPriority, lowPriority...)
	return highPriority
}

// Discover finds known serial devices
func Discover() ([]string, error) {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		return nil, err
	}
	if len(ports) == 0 {
		return []string{}, nil
	}
	return findCandidates(ports), nil
}
