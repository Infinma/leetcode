func isPalindrome(x int) bool {
	if x < 0 {
			return false
	}
	temp := x
	comp := 0
	for temp > 0 {
			comp = comp * 10 + temp % 10
			temp /= 10
	}
	if comp == x {
			return true
	} else {
			return false
	}
}
