package lstructs

// For log2(n) >= 0
func fastLog2(n int) int {
	if n <= 0 {
		return -1
	}

	res := 0
	for n >= 2 {
		n /= 2
		res++
	}

	return res
}

// For n := [0, inf]
func fastPow2(n int) int {
	return 1 << n
}
