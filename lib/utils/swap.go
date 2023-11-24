package utils

// Swap exchanges values between two items.
func Swap[V int | float64 | string](i, j *V) {
	temp := *i
	*i = *j
	*j = temp
}
