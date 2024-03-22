package queue

import "sync"

type Queue[T any] struct {
	Items []T
	lock  sync.Mutex
}

func (q *Queue[T]) Push(i T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.Items = append(q.Items, i)
}

func (q *Queue[T]) Pop() T {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.Items) == 0 {
		var zero T
		return zero
	}

	item := q.Items[0]
	q.Items = q.Items[1:]

	return item
}
