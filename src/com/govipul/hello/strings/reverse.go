package strings

func Reverse(value string) string {
	runes := []rune(value)
	reversedRunes := reverseRune(runes)
	return string(reversedRunes)
}
