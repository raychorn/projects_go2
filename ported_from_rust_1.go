package main

import (
    "fmt"
    "time"
)

func main() {
    start := time.Now().Nanosecond()
    var sumOfMultiples int64 = 0

    for i := 0; i < 1000; i++ {
        if (isMultiple(i)) {
            sumOfMultiples = sumOfMultiples + int64(i)
        };
    }
    fmt.Printf("Sum is %d\n", sumOfMultiples)
    now := time.Now().Nanosecond()
    durationMs  := int(((now - start) % 1e6) / 1e3)
    fmt.Printf("Time elapsed is: %d\n", durationMs)
}

func isMultiple(num int) (result bool) {
	return ( (num % 5 == 0) || (num % 3 == 0) )
}

