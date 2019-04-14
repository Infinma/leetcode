func letterCombinations(digits string) []string {
	m := make(map[rune][]string)
	m['2'] = []string{"a", "b", "c"}
	m['3'] = []string{"d", "e", "f"}
	m['4'] = []string{"g", "h", "i"}
	m['5'] = []string{"j", "k", "l"}
	m['6'] = []string{"m", "n", "o"}
	m['7'] = []string{"p", "q", "r", "s"}
	m['8'] = []string{"t", "u", "v"}
	m['9'] = []string{"w", "x", "y", "z"}
	
	var prev []string
	var sol []string
	for i, num := range digits {
			sol = nil
			if i == 0 {
					sol = append(sol, m[num]...)
			} else {
					for _, pre := range prev {
							for _, char := range m[num] {
									sol = append(sol, pre + char)
							}
					}
			}
			prev = nil
			prev = append(prev, sol...)
	}
	return sol
}
