func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
			return ""
	} else if len(strs) == 1 {
			return strs[0]
	}
	
	var c string
	for i := range strs[0] {
			c = strs[0][i:i+1]
			for _, str := range strs {
					if (i == len(str) || str[i:i+1] != c) {
							return strs[0][:i]
					}
			}
	}
	return strs[0]
}
