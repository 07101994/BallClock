package main

import (
	"fmt"
)

func main() {
	var i int

	for  {
		_, err := fmt.Scanf("%d\n", &i)
		
		if err != nil {  fmt.Println("ERROR: ", err); break }
		if i == 0 { break }
		
		if i >= 27 && i <= 127 {
			c := NewClock(i)
			if days, e := c.getLCM() / 2; e == nil {
				fmt.Println(i, "balls cycle after ", days, " days")
			} else {
				fmt.Println("ERROR: %v", e)
			}
		}
	}

	return
}
