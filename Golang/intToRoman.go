func intToRoman(num int) string {
	M := [4]string{"", "M", "MM", "MMM"}
	C := [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	X := [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	I := [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	return M[num/1000] + C[(num % 1000)/100] + X[(num % 100)/10] + I[(num % 10)]
}
