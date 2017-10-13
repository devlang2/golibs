package strings

func StrPadRight(str string, length int, pad string) string {
	return str + times(pad, length-len(str))
}

func StrPadLeft(str string, length int, pad string) string {
	return times(pad, length-len(str)) + str
}

func times(str string, n int) (out string) {
	for i := 0; i < n; i++ {
		out += str
	}
	return
}
