package ratemeter

import (
	"time"
)

// RateMeter는 이벤트 발생 빈도를 측정하는 기능을 가지고 있다.
// 이벤트가 발생할 때  Incr를 통해 ticker를 증가시키고,
// TimeFrame 주기로 flushTicker를 호출해 ticker를 circularQueue에 저장한다.
//
// 이 구현은 Thread-safe하지 않으며, 대략의 추세를 측정하기 위해 사용한다.
//
type RateMeter struct {
	TimeFrame time.Duration
	CircularQueue
	flushHook func(int)
	ticker    int
	dispose   bool
}

// NewRateMeter는 주어진 타임프레임과 숫자 프레임을 가지는 RateMeter를 생성한다.
// 주기적으로 호출되는 함수는 주어진 프레임 단위로 호출되며,
// 호출되는 함수는 주어진 프레임 단위로 시간을 지연하고,
func NewRateMeter(timeFrame time.Duration, numberOfFrame int) *RateMeter {
	r := &RateMeter{
		TimeFrame:     timeFrame,
		CircularQueue: NewCircularQueue(numberOfFrame),
	}

	go func() {
		for {
			if r.dispose {
				return
			}
			time.Sleep(timeFrame)
			r.flushTicker()
		}
	}()

	return r
}

func (r *RateMeter) Incr() {
	r.ticker++
}

func (r *RateMeter) Sum() int {
	sum := r.ticker
	for _, v := range r.CircularQueue.quque {
		sum += v
	}
	return sum
}

func (r *RateMeter) Dispose() {
	r.dispose = true
}

func (r RateMeter) Size() int {
	return r.CircularQueue.size
}

func (r *RateMeter) Clear() {
	r.ticker = 0
	r.CircularQueue = NewCircularQueue(r.CircularQueue.size)
}

func (r *RateMeter) SetFlushHook(hook func(ticker int)) {
	r.flushHook = hook
}

func (r *RateMeter) flushTicker() {
	if r.flushHook != nil {
		r.flushHook(r.ticker)
	}
	r.CircularQueue.Push(r.ticker)
	r.ticker = 0
}
