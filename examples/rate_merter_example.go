package main

import (
	"fmt"
	"math/rand"
	"ratemeter"
	"time"
)

func main() {
	m := ratemeter.NewRateMeter(time.Second, 60)

	rng := rand.New(rand.NewSource(time.Now().Unix()))
	for {
		m.Incr()
		time.Sleep(time.Millisecond * time.Duration((rng.Intn(5)+5)*100))
		fmt.Println(m)
		fmt.Println(m.Sum())
	}
}
