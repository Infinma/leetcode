func myAtoi(str string) int {
	hasStarted := false
	end := false
	pos := true
	sol := 0
	for _, v := range str {
			switch {
					case v >= '0' && v <= '9':
					sol = sol * 10 + int(v - '0')
					hasStarted = true
					case v == '-':
					if hasStarted {
							end = true
					} else {
							pos = false
							hasStarted = true
					}
					case v == '+':
					if hasStarted {
							end = true
					}
					hasStarted = true
					case v == ' ':
					if hasStarted {
							end = true
					}
					default:
					if !hasStarted {
							return 0
					} else {
							end = true
					}
			}
			temp := sol
			if !pos {
					temp *= -1
			}
			if temp > math.MaxInt32 {
					return math.MaxInt32
			} else if temp < math.MinInt32 {
					return math.MinInt32
			} else if end {
					break
			}
	}
	if (!pos) {
			sol *= -1
	}
	return sol
}