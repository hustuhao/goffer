package ipqueue

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/facebookgo/ensure"
	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	c := 100
	pq := NewQueue[int]()

	for i := 0; i < c+1; i++ {
		pq.Offer(&Item[int]{Value: i, Priority: int64(i)})
	}
	assert.Equal(t, pq.Len(), c+1)
	for i := 0; i < c+1; i++ {
		item := pq.Poll()
		assert.Equal(t, item.Value, i)
	}
}

func TestUnsortedInsert(t *testing.T) {
	c := 100
	pq := NewQueue[int]()
	ints := make([]int, 0, c)

	for i := 0; i < c; i++ {
		v := rand.Int()
		ints = append(ints, v)
		pq.Offer(&Item[int]{Value: i, Priority: int64(v)})
	}
	assert.Equal(t, pq.Len(), c)

	sort.Sort(sort.IntSlice(ints))

	for i := 0; i < c; i++ {
		item := pq.Poll()
		assert.Equal(t, int(item.Priority), ints[i])
	}
}

func TestRemove(t *testing.T) {
	c := 100
	pq := NewQueue[string]()

	for i := 0; i < c; i++ {
		v := rand.Int()
		pq.Offer(&Item[string]{Value: "test", Priority: int64(v)})
	}

	for i := 0; i < 10; i++ {
		pq.Poll()
	}

	lastPriority := pq.Poll().Priority
	for i := 0; i < (c - 10 - 1); i++ {
		item := pq.Poll()
		ensure.DeepEqual(t, lastPriority < item.Priority, true)
		lastPriority = item.Priority
	}
}

func TestPriorityQueueWithZeroCapacity(t *testing.T) {
	pq := NewQueue[int](&Item[int]{Value: 1, Priority: 1})
	ensure.DeepEqual(t, cap(*pq), 1)
}
