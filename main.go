package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countingValleys(int32(8), string("UDDDUDUU")))
}

func countingValleys(steps int32, path string) int32 {
	chars := strings.Split(path, "")
	var condition bool
	down := string("D")
	sum := int32(0)
	valleysCount := int32(0)
	for i := int32(0); i < steps; i++ {
		condition = chars[i] == down
		sum += int32(1)
		if condition {
			sum -= int32(2)
		}
		if i > int32(0) && sum == int32(0) && !condition {
			valleysCount += int32(1)
		}
	}
	return valleysCount
}
