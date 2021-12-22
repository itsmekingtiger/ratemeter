package main

import (
	"fmt"
	"math/rand"
	"ratemeter"
	"time"
)

func main() {
	m := ratemeter.NewRateMeter(time.Second, 10)
	m.SetFlushHookBefore(func(ticker int) {
		fmt.Println(m.CircularQueue)
	})

	rng := rand.New(rand.NewSource(time.Now().Unix()))
	for {
		count := rng.Intn(100)
		for i := 0; i < count; i++ {
			m.Incr()
		}
		time.Sleep(time.Second)
	}
}
