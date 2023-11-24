package sorting

import (
	"github.com/codelabs/gobank/lib/v1/utils"
	"log/slog"
)

// BubbleSort is the simplest sorting algorithm that works by repeated swapping the adjacent elements if they are in
// the wrong order.
func BubbleSort[V int | float64](list []V) {

	n := len(list)
	slog.Info("list size", "size", n)

	// swapCounter tracks number of swaps performed.
	swapCounter := 0

	for i := 0; i < n-1; i++ {

		swapped := false
		scanSize := n - 1 - i
		slog.Debug("scan size", "scanSize", scanSize)

		for j := 0; j < scanSize; j++ {

			if list[j] > list[j+1] {

				utils.Swap(&list[j], &list[j+1])
				swapCounter++
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	slog.Debug("num swaps", "swapCount", swapCounter)

	if swapCounter == 0 {
		slog.Warn("already sorted")
	}
}
