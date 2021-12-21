package main

import (
	"fmt"
	"math/rand"
	"ratemeter"
	"time"
)

func main() {
	m := ratemeter.NewRateMeter(time.Second, 60)
	m.SetFlushHook(func(ticker int) {
		fmt.Printf("단위시간동안 %d번의 이벤트 발생\n", ticker)
		fmt.Printf("최근 1분간 이벤트 발생 횟수: %d\n", m.Sum())
	})

	rng := rand.New(rand.NewSource(time.Now().Unix()))
	for {
		m.Incr()
		sleep := rng.Intn(1000) + 1
		fmt.Printf("sleep: %d\n", sleep)
		time.Sleep(time.Millisecond * time.Duration(sleep))
	}
}
