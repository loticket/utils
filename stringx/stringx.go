package stringx

func Len(str string) int {
	return utf8.RuneCountInString(str)
}
