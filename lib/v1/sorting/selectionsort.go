package sorting

import (
	"fmt"
	"github.com/codelabs/gobank/lib/v1/utils"
	"log/slog"
)

// SelectionSort ...
func SelectionSort(list []int) {

	slog.Info("Before Sort", "list", list)
	n := len(list)

	for i := 0; i < n-1; i++ {

		// Controls the starting point for the sublist
		start := i
		// Assuming the smallest item is at beginning of the sublist
		minItem := list[start]
		// Position or index of the smallest element found.
		position := start

		// Finding the minimum item from the sublist, with its position
		for j := start + 1; j < n; j++ {
			if minItem > list[j] {
				minItem = list[j]
				position = j
			}
		}

		utils.Swap(&list[start], &list[position])

		iteration := fmt.Sprintf("iteration-%d", i+1)
		slog.Debug(iteration, "min", minItem, "pos", position, "list", list)
	}

	slog.Info("After Sort", "list", list)
}
