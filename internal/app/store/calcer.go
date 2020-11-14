package store

// CalcSumOfBytes ...
func CalcSumOfBytes(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(s[i])
	}

	return sum
}
