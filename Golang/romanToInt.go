func romanToInt(s string) int {
	var curr int
	var next int
	m := make(map[byte]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000
	sol := 0
	for i := 0; i < len(s); i++ {
			curr = m[s[i]]
			if i + 1 < len(s) {
					next = m[s[i + 1]]
			}
			if curr < next {
					sol += next - curr
					i += 1
			} else {
					sol += curr
			}
			next = 0
	}
	return sol
}
