package priorityqueue

import (
	"holocron/pkg/pubsub"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	queue := NewQueue(3)
	distances := []float64{10, 10.3, 9, 5, 10.1, 4}

	for _, distance := range distances {
		queue.Push(distance, &pubsub.Event{})
	}

	result := queue.Get()

	assert.Equal(t, 3, len(result))
	assert.Equal(t, float64(4), result[0].distance)
	assert.Equal(t, float64(5), result[1].distance)
	assert.Equal(t, float64(9), result[2].distance)
}
