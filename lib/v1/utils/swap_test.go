package utils

import "testing"

func TestSwap(t *testing.T) {

	i := 10
	j := 20

	Swap(&i, &j)

	if i != 20 {
		t.Fail()
	}

	if j != 10 {
		t.Fail()
	}
}
