package ratemeter

import "errors"

type CircularQueue struct {
	quque []int
	head  int
	tail  int
	size  int
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		quque: make([]int, size),
		size:  size,
	}
}

func (q *CircularQueue) Push(i int) {
	q.quque[q.head] = i
	q.moveHead()
}

// Pop은 tail의 위치에 있는 값을 꺼내고,
// 이 값을 리턴한다.
// 만약 큐가 비어있다면 에러를 리턴한다.
func (q *CircularQueue) Pop() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	v := q.quque[q.tail]
	q.moveTail()
	return v, nil
}

func (q CircularQueue) IsEmpty() bool {
	return q.head == q.tail
}

// 큐의 복사본을 반환한다.
func (q CircularQueue) Queue() []int {
	dump := make([]int, q.size)
	copy(dump, q.quque)
	return dump
}

// moveHead는 head를 다음 위치로 이동하고
// 만약 head가 size보다 크면 head를 처음 위치로 이동한다.
// 만약 head가 tail과 같으면 tail을 다음 위치로 이동한다.
func (q *CircularQueue) moveHead() {
	q.head = (q.head + 1) % q.size

	if q.head == q.tail {
		q.moveTail()
	}
}

// moveTail는 tail을 다음 위치로 이동하고
// 만약 tail이 size보다 크면 tail을 처음 위치로 이동한다.
func (q *CircularQueue) moveTail() {
	q.tail = (q.tail + 1) % q.size
}
