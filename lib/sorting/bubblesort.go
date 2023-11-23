package sorting

import "log/slog"

func BubbleSort(list []int) {

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

				swap(&list[j], &list[j+1])
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

// swap exchanges values between two items.
func swap(i, j *int) {
	temp := *i
	*i = *j
	*j = temp
}
