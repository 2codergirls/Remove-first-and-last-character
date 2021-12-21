package kata

func RemoveChar(word string) string {
	lastCharacter := len(word) - 1
	return word[1:lastCharacter]
}
