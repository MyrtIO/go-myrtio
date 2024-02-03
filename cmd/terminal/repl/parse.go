package repl

import (
	"strconv"
	"strings"
)

func ParsePayload(input string) ([]byte, error) {
	symbols := strings.Fields(input)
	result := make([]byte, 0, len(symbols))
	for _, symbol := range symbols {
		value, err := strconv.Atoi(symbol)
		if err != nil {
			return nil, err
		}
		result = append(result, byte(value))
	}
	return result, nil
}
