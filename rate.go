package gorate

import (
	"time"
)

// Rater는 TimeFrame 주기로 ticker를 circularQueue에 저장한다.
type Rater struct {
	TimeFrame     time.Duration
	circularQueue CircularQueue
	ticker        int
	dispose       bool
}

// NewRater는 주어진 타임프레임과 숫자 프레임을 가지는 Rater를 생성한다.
// 주기적으로 호출되는 함수는 주어진 프레임 단위로 호출되며,
// 호출되는 함수는 주어진 프레임 단위로 시간을 지연하고,
func NewRater(timeFrame time.Duration, numberOfFrame int) *Rater {
	r := &Rater{
		TimeFrame:     timeFrame,
		circularQueue: NewCircularQueue(numberOfFrame),
	}

	go func() {
		for {
			if r.dispose {
				return
			}
			time.Sleep(timeFrame)
			r.tick()
		}
	}()

	return r
}

func (r *Rater) Incr() {
	r.ticker++
}

func (r *Rater) Sum() int {
	sum := r.ticker
	for _, v := range r.circularQueue.quque {
		sum += v
	}
	return sum
}

func (r *Rater) Dispose() {
	r.dispose = true
}

func (r Rater) Size() int {
	return r.circularQueue.size
}

func (r *Rater) tick() {
	r.circularQueue.Push(r.ticker)
	r.ticker = 0
}
