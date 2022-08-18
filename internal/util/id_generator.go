package util

import (
	"math/rand"
	"sync"
	"time"

	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

var rg = struct {
	sync.Mutex
	rand *rand.Rand
}{

	rand: rand.New(rand.NewSource(time.Now().UnixNano())),
}

func init() {

	var f sonyflake.Settings
	f.StartTime = time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC)
	sf = sonyflake.NewSonyflake(f)
	if sf == nil {

		panic("id generator init error.")
	}
}

func Int63nRange(min, max int64) int64 {

	rg.Lock()
	defer rg.Unlock()
	return rg.rand.Int63n(max-min) + min
}

func GetNextId() uint64 {

	ret, err := sf.NextID()
	if err != nil {

		ret = uint64(Int63nRange(1926425572, 1926425572223607))
	}
	return ret
}
