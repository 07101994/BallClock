package main

import (
	"BallClock/stack"
)

type Clock struct {
	ones stack.Stack
	fives stack.Stack
	hours stack.Stack
	cycles int
	queue chan int
	index [] int
}

func NewClock(size int) *Clock {
	var clock Clock
	clock.index = make([]int, size)
	clock.queue = make(chan int, size)
	
	clock.hours.Push(-1)

	for i := 1; i <= size; i++ {
		clock.queue <- i
	}
	
	return &clock
}

func (c *Clock) getLCM() int {
	// track how many balls haven't repeated yet
	remaining := len(c.index)
	
	// increment clock until we've logged the repeat frequency of each ball
	for remaining > 0 {
		remaining -= c.increment()
	}
	
	set := map[int]int { }
	for i := 0; i < len(c.index); i++ {
		v := c.index[i]
		if v != 1 {
			set[v] = 1
		}	
	}
	
	result := 1
	for k, _ := range set {
		result = LCM(result, k)
	}
	
	return result
}

func (c *Clock) increment() int {
	repeats := 0
	
	// increment by one minute
	c.ones.Push(<- c.queue)
	if c.ones.Len() < 5 { return repeats }
	
	// top minute becomes a five-minute ball
	minute, _ := c.ones.Pop()
	c.fives.Push(minute)
	
	// re-queue the remaining 4 balls
	for c.ones.Len() > 0 {
		item, _ := c.ones.Pop()
		c.queue <- item.(int)
	}
	
	// 12 five-minute balls == hour
	if c.fives.Len() < 12 { return repeats }
	
	// top five-minute ball becomes an hour
	hour, _ := c.fives.Pop()
	
	// re-queue the remaining 11 balls
	for c.fives.Len() > 0 {
		item, _ := c.fives.Pop()
		c.queue <- item.(int)
	}
	
	if c.hours.Len() < 12 {
		c.hours.Push(hour)
	} else {
		// 12 hours == cycle
		c.cycles++
		
		for c.hours.Len() > 1 {
			item, _ := c.hours.Pop()
			c.queue <- item.(int)
		}
		
		// the 13th hour gets re-queued last
		c.queue <- hour.(int)
		
		// update the number of balls that repeated
		repeats = c.logRepeats()
	}
	
	// return balls that repated in this cycle
	return repeats
}

func (c *Clock) logRepeats() int {
	found := 0
	item := 0
	
	// inspect each ball and its order on the queue
	for i := 0; i < len(c.index); i++ {
		// inspect the newest ball in the queue
		item = <- c.queue
		
		if item == i + 1 && c.index[i] == 0 {
			found++
			// set the cycle the repeat was logged
			c.index[i] = c.cycles
		}
		
		// return the ball to the end of the queue
		c.queue <- item
	}
	
	// return the number of balls that repeated in the current cycle
	return found
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