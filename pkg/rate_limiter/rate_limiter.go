package rate_limiter

import "time"

var (
	count   = 0
	lastSec = 0
)

const LIMIT_THRESHOLD = 10
const TIME_WINDOW = 10

func LimitReached() bool {
	secondStore := make(map[int]int)
	// sliding window
	if count == LIMIT_THRESHOLD {
		return true
	}
	currSec := int(time.Now().Unix())
	count++
	secondStore[currSec]++
	// window has to be changed when currSec > lastSec
	// limitation: if no request in a particular second, then it won't subtract
	// change required here
	if currSec > lastSec {
		count -= secondStore[currSec-TIME_WINDOW]
		lastSec = currSec
	}
	return false
}
