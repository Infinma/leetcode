func reverse(x int) int {
	sol := 0
	curr := x
	pos := true
	if curr < 0 {
			pos = false
			curr *= -1
	}
	for curr > 0 {
			sol = (sol * 10) + (curr % 10)
			curr /= 10
			if (sol > math.MaxInt32) {
					return 0
			}
	}
	if !pos {
			sol *= -1
	}
	return sol
}
