func threeSum(nums []int) [][]int {
	var neg []int
	var pos []int
	var sol [][]int
	var sum int
	m := make(map[int]int)
	for _, v := range nums {
			_, ok := m[v]
			if ok {
					m[v] += 1
			} else {
					m[v] = 1
			}
	}
	
	for key := range m {
			if key < 0 {
					neg = append(neg, key)
			} else {
					pos = append(pos, key)
			}
	}
	
	if zero, ok := m[0]; ok {
			if zero >= 3 {
					arr := []int{0, 0, 0}
					sol = append(sol, arr)
			}
	}
	
	for _, p := range pos {
			for _, n := range neg {
					sum = - p - n
					if count, ok := m[sum]; ok {
							if (sum == p || sum == n) && count > 1 {
									arr := []int{n, sum, p}
									sol = append(sol, arr)
							}
							if sum > p {
									arr := []int{n, p, sum}
									sol = append(sol, arr)
							}
							if sum < n {
									arr := []int{sum, n, p}
									sol = append(sol, arr)
							}
					}
			}
	}
	return sol
}
