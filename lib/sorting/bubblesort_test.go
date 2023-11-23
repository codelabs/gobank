package sorting

import (
	"log/slog"
	"os"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	t.Run("unsorted list", func(t *testing.T) {
		list := []int{45, 22, 10, 32, 67, 5, 18, 52, 4}
		BubbleSort(list)
		expected := []int{4, 5, 10, 18, 22, 32, 45, 52, 67}
		if !isListsEqual(list, expected) {
			t.Fail()
		}
	})

	t.Run("sorted list", func(t *testing.T) {
		list := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
		BubbleSort(list)
		expected := list
		if !isListsEqual(list, expected) {
			t.Fail()
		}
	})
}

func isListsEqual(a, b []int) bool {

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
