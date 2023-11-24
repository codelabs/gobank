package sorting

import (
	"log/slog"
)

// InsertionSort is a simple sorting algorithm that works similar to the way you sort playing cards in your hands.
// The array is virtually split into a sorted and an unsorted part.
// Values from the unsorted part are picked and placed at the correct position in the sorted part.
func InsertionSort(list []int) {

	n := len(list)
	slog.Debug("", "listSize", n)
	slog.Debug("Before Sort", "list", list)

	for i := 1; i < n; i++ {

		key := list[i]
		slog.Debug("", "i", i, "key", key)

		for j := i - 1; j >= 0; j-- {
			if list[j] > key {
				slog.Debug("greater than key found", "j", j)
				list[j+1] = list[j]
				list[j] = key
			}
		}
	}

	slog.Debug("After Sort", "list", list)
}
