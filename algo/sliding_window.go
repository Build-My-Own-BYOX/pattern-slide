package algo

func isMatchable(text string, pattern string) bool {
	textLen := len(text)
	patternLen := len(pattern)
	if textLen <= 0 || patternLen <= 0 {
		return false
	}
	if textLen < patternLen {
		return false
	}
	return true
}

func SingleSlidingWindow(text string, pattern string) (occurrences []int) {
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
