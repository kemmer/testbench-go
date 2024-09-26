package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, world!")
}

func Int64ToInt(val int64) (int, error) {
	if val > int64(math.MaxInt) {
		return 0, errors.New(fmt.Sprintf("value provided exceeds the maximum int size of %v", math.MaxInt))
	}

	return int(val), nil
}
