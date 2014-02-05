package main

import (
	"BallClock/stack"
	"fmt"
)

func main() {
	for i := 27; i <= 127; i++ {
		lcm := getLCM(i) / 2
		fmt.Println(i, "balls cycle after ", lcm, " days")
	}

	return
}

func getLCM(size int) int {
	gcdNeeded := size
	cycles := 0
	
	gcdIndex := make([]int, size)
	
	var ones stack.Stack
	var fives stack.Stack
	var hours stack.Stack

	hours.Push(-1)

	queue := make(chan int, size)

	for i := 1; i <= size; i++ {
		queue <- i
	}

	for gcdNeeded > 0 {
		ones.Push(<- queue)
		if ones.Len() < 5 { 
			continue 
		}

		item, err := ones.Pop()
		if(err != nil){ break }
		fives.Push(item)
		
		for ones.Len() > 0 {
			item, err := ones.Pop()
			if(err != nil){ break }
			queue <- item.(int)
		}
		
		if fives.Len() < 12 {
			continue
		}
		
		hour, err := fives.Pop()
		if(err != nil){
			break
		}
		for fives.Len() > 0 {
			item, err := fives.Pop()
			if(err != nil){
				break
			}
			queue <- item.(int)
		}
		
		
		if hours.Len() < 12 {
			hours.Push(hour)
		} else {
			for hours.Len() > 1 {
				item, err := hours.Pop()
				if(err != nil){
					break
				}
				queue <- item.(int)
			}
			
			queue <- hour.(int)
			
			cycles++
			
			gcdNeeded = CheckOrder(queue, gcdIndex, cycles, gcdNeeded)
		}
	}
	
	lcdSet := map[int]int { }
	
	for i := 0; i < len(gcdIndex); i++ {
		v := gcdIndex[i]
		if v != 1 {
			lcdSet[v] = 1
		}
	}
	
	result := 1
	for k,_ := range lcdSet {
		result = LCM(result, k)
	}
	
	return result
}

func LCM(a int, b int) int {
	max := 0
	min := 0
	
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}

	result := min
	
	for i := 1; i <= min; i++ {
		if (max * i) % min == 0 {
			result = i * max
			break
		}
	}
	
	return result
}

func CheckOrder(queue chan int, gcdIndex []int, cycles int, remaining int) int {
	if remaining == 0 { return remaining }
	
	item := 0
	for i := 0; i < len(gcdIndex); i++ {
		item = <- queue
		
		if item == i + 1 && gcdIndex[i] == 0 {
			remaining--
			gcdIndex[i] = cycles
		}
		
		queue <- item
	}

	return remaining
}
