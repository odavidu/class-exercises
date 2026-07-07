package main

import (
	"math"
	"strconv"
	"strings"
)

func Map(key string, value string) []KVPair {

	output := make([]KVPair, 0)

	// TODO: This loop iterates over each line of the "value" string
	// You will want to parse out the date and temperature from each line and add it to the "output" slice
	for _, line := range strings.Split(strings.TrimSuffix(value, "\n"), "\n") {
		parts := strings.Split(line, ",")

		year := strings.Split(parts[1], "-")[0]
		temp := parts[2]

		output = append(output, KVPair{key: year, value: temp})
	}

	return output
}

func Reduce(key string, value []string) float64 {

	// Converting from a string to float may be useful
	max := math.NaN()
	for _, str := range value {
		val, err := strconv.ParseFloat(str, 64)

		if err != nil {
			panic(err)
		}
		if math.IsNaN(max) || val > max {
			max = val
		}
	}

	return max
}
