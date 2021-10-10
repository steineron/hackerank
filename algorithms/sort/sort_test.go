package sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func sortedList(size int) []int {
	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = i
	}
	return list
}

func unsortedList(size int) []int {
	return rand.Perm(size)
}

func randomList(size int) []int {
	list := make([]int, size)
	rnd := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < size; i++ {
		list[i] = rnd.Int()
	}
	return list
}

func TestSorting(t *testing.T) {
	var sorters = []Sorter{
		&bubbler{},
		&insertion{},
	}

	type test struct {
		input    []int
		expected []int
	}

	tests := []test{
		{
			input:    nil,
			expected: nil,
		}, {
			input:    []int{},
			expected: []int{},
		}, {
			input:    []int{1},
			expected: []int{1},
		}, {
			input:    []int{1, 2},
			expected: []int{1, 2},
		}, {
			input:    []int{2, 1},
			expected: []int{1, 2},
		}, {
			input:    sortedList(200),
			expected: sortedList(200),
		}, {
			input:    unsortedList(200),
			expected: sortedList(200),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			for _, sorter := range sorters {
				var in []int
				if test.input != nil {
					in = make([]int, len(test.input))
					copy(in, test.input)
				}

				sort := sorter.Sort(in)
				if !reflect.DeepEqual(test.expected, sort) {
					t.Errorf("%v did not sort correctly: %v", sorter, sort)
				} else {
					t.Logf("%v sort correctly: %v", sorter, sort)

				}

			}
		})
	}
}
