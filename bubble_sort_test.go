package sort_test

import (
	"github.com/aceakash/sort"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestSortSingleElement(t *testing.T) {
	list := []int{2}
	sort.BubbleSort(list)
	assert.Equal(t, []int{2}, list)
}

func TestSortShortList(t *testing.T) {
	list := []int{4, 3, 2}
	sort.BubbleSort(list)
	assert.Equal(t, []int{2, 3, 4}, list)
}

func BenchmarkName(b *testing.B) {
	list := bigList(10 * 1000)
	for i := 0; i < b.N; i++ {
		sort.BubbleSort(list)
	}
}

func bigList(size int) []int {
	rand.Seed(time.Now().UnixNano())

	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = rand.Int()
	}
	return list
}
