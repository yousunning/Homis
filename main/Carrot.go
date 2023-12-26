package main

import (
	"fmt"
)

func main() {
	a := 56
	if a > 30 {
		fmt.Println("맞다면 당근을 흔들어")
	} else if a < 30 {
		fmt.Println("틀려도 당근을 흔들어")
	} else {
		fmt.Println("당근을 냅다 흔들어")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
