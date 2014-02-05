package main

import (
	"BallClock/stack"
	"fmt"
)

func main() {
	for i := 27; i <= 127; i++ {
		c := NewClock(i)
		days := c.getLCM() / 2
		fmt.Println(i, "balls cycle after ", days, " days")
	}

	return
}
