package countdown

import (
	"fmt"
	"time"
)

type Countdown struct {
	times        []int64
	lastCount    int64
	lastEstimate int64
	lastCheck    int64
	count        int64
}

func (c *Countdown) Start(length int) {
	c.times = []int64{}
	c.lastEstimate = 0
	c.lastCheck = time.Now().Unix()
	c.count = int64(length)
	c.lastCount = time.Now().UnixNano()
}

func (c *Countdown) Count() {
	c.times = append(c.times, time.Now().UnixNano()-c.lastCount)
	c.lastCount = time.Now().UnixNano()
	c.count--
	if time.Now().Unix() > c.lastCheck+3 {
		c.lastCheck = time.Now().Unix()
		total := int64(0)
		for _, time := range c.times {
			total = total + time
		}
		averageTime := total / int64(len(c.times))
		timeRemaining := averageTime * c.count
		timeRemainingSeconds := timeRemaining / 1000000000
		if timeRemainingSeconds != c.lastEstimate {
			if timeRemainingSeconds > 60 {
				timeRemainingMinutes := timeRemainingSeconds / 60
				fmt.Printf("Estimated time remaining: %d minutes\n", timeRemainingMinutes)
			} else {
				fmt.Printf("Estimated time remaining: %d seconds\n", timeRemainingSeconds)
			}
			c.lastEstimate = timeRemainingSeconds
		}
	}
}
