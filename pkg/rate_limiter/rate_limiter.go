package rate_limiter

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gagandeepahuja09/redisgo/pkg/redisClient"
	"github.com/go-redis/redis/v8"
)

var (
	count   = 0
	lastSec = 0
)

const LIMIT_THRESHOLD = 10
const TIME_WINDOW = 10

// Sliding Window Implementation
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

// Leaky Bucket Implementation
func LimitReachedRedis() bool {
	ctx := context.Background()

	// use redis to save this instead
	countVal, err := redisClient.Rdb.Get(ctx, "count").Result()
	if err == redis.Nil {
		countVal = "0"
	} else if err != nil {
		panic(err)
	}
	redisCount, err := strconv.Atoi(countVal)
	if err != nil {
		panic(err)
	}
	// sliding window
	if redisCount >= LIMIT_THRESHOLD {
		return true
	}
	// should we do both together in a pipeline?
	currSec := int(time.Now().Unix())
	redisClient.Rdb.Incr(ctx, "count")
	redisClient.Rdb.Incr(ctx, fmt.Sprintf("%d:count", currSec))
	redisClient.Rdb.Expire(ctx, fmt.Sprintf("%d:count", currSec), time.Second*TIME_WINDOW+time.Millisecond)
	// get the value of the count which was 10 seconds ago and decrement it.
	prevValue, err := redisClient.Rdb.Get(ctx, fmt.Sprintf("%d:count", currSec-TIME_WINDOW)).Result()
	if err == redis.Nil {
		prevValue = "0"
	} else if err != nil {
		panic(err)
	}
	prevValueInt, err := strconv.Atoi(prevValue)
	if err != nil {
		panic(err)
	}
	redisClient.Rdb.DecrBy(ctx, "count", int64(prevValueInt))
	return false
}
