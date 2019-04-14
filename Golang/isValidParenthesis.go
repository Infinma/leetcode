func isValid(s string) bool {
	var stack []rune
	for _, char := range s {
			switch {
					case char == '(' || char == '[' || char == '{':
					stack = append(stack, char)
					case char == ')':
					if len(stack) < 1 || stack[len(stack)-1] != '(' {
							return false
					}
					stack = stack[:len(stack)-1]
					case char == ']':
					if len(stack) < 1 || stack[len(stack)-1] != '[' {
							return false
					}
					stack = stack[:len(stack)-1]
					case char == '}':
					if len(stack) < 1 || stack[len(stack)-1] != '{' {
							return false
					}
					stack = stack[:len(stack)-1]
			}
	}
	if len(stack) > 0 {
			return false
	}
	return true
}
