package iterations

func Repeat(s string, cnt int) (result string) {
	for i := 0; i < cnt; i++ {
		result += s
	}
	return result
}
