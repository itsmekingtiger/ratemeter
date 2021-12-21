package gorate

import "time"

type Rater struct {
	TimeFrame     time.Duration
	NumberOfFrame int
	circularQueue CircularQueue
	ticker        int
}

func NewRater(timeFrame time.Duration, numberOfFrame int) *Rater {
	return &Rater{
		TimeFrame:     timeFrame,
		NumberOfFrame: numberOfFrame,
		circularQueue: NewCircularQueue(numberOfFrame),
	}
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

func (r *Rater) tick() {
	r.circularQueue.Push(r.ticker)
	r.ticker = 0
}
