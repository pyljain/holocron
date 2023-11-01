package priorityqueue

import (
	"holocron/pkg/pubsub"
	"sort"
)

type priorityQueue struct {
	elements []element
	topK     int
}

type element struct {
	distance      float64
	MatchedVector *pubsub.Event
}

type byDistance []element

func (a byDistance) Len() int           { return len(a) }
func (a byDistance) Less(i, j int) bool { return a[i].distance < a[j].distance }
func (a byDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func NewQueue(topK int) *priorityQueue {
	return &priorityQueue{
		topK: topK,
	}
}

func (q *priorityQueue) Push(distance float64, vector *pubsub.Event) {
	if len(q.elements) < q.topK {
		q.elements = append(q.elements, element{
			distance:      distance,
			MatchedVector: vector,
		})

		sort.Sort(byDistance(q.elements))

		return
	}

	if q.elements[len(q.elements)-1].distance < distance {
		return
	}

	q.elements[len(q.elements)-1].distance = distance
	q.elements[len(q.elements)-1].MatchedVector = vector

	sort.Sort(byDistance(q.elements))
}

func (q *priorityQueue) Get() []element {
	return q.elements
}
