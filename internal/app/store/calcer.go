package store

// CalcSumOfBytes calculate sum of bytes in word
func CalcSumOfBytes(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(s[i])
	}

	return sum
}
