// Package clock provides create, add and subtract methods on Clock type
package clock

import (
	"fmt"
)

// Clock is our custom clock
type Clock struct {
	minutes int
}

const minutesInADay = 24 * 60
const minutesInAHour = 60

// New creates new Clock
func New(hour int, minute int) Clock {
	m := minute + hour*60
	m %= 24 * 60
	if m < 0 {
		m += 24 * 60
	}
	return Clock{m}
}

// String prints current time
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/minutesInAHour, c.minutes%minutesInAHour)
}

// Add provides add functionality on our custom Clock type
func (c Clock) Add(minutes int) Clock {
	c.minutes = c.minutes + minutes%minutesInADay
	for c.minutes >= minutesInADay {
		c.minutes = c.minutes - minutesInADay
	}
	for c.minutes < 0 {
		c.minutes = c.minutes + minutesInADay
	}
	return c
}

// Subtract provides subtract functionality on our custom Clock type
func (c Clock) Subtract(minutes int) Clock {
	c.minutes = c.minutes - minutes%minutesInADay
	for c.minutes < 0 {
		c.minutes = c.minutes + minutesInADay
	}
	return c
}
