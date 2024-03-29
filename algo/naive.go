package algo

func BruteForceSearch(text string, pattern string) (occurrences []int) {
	if !isMatchable(text, pattern) {
		return
	}

	for pos := 0; pos <= len(text)-len(pattern); pos++ {
		isMismatch := false
		for sPos := pos; sPos < pos+len(pattern); sPos++ {
			if text[sPos] != pattern[sPos-pos] {
				isMismatch = true
				break
			}
		}
		if !isMismatch {
			occurrences = append(occurrences, pos)
		}
	}
	return
}
