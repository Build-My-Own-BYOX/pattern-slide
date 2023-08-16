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

func SingleSlidingWindowSearch(text string, pattern string) (occurrences []int) {
	if !isMatchable(text, pattern) {
		return
	}

	firstPatternChr := pattern[0]

	for pos := 0; pos <= len(text)-len(pattern); pos++ {
		isMismatch := false
		jmp := 0
		for sPos := pos; sPos < pos+len(pattern); sPos++ {
			// find next potential matched position
			if jmp > 0 && text[sPos] == firstPatternChr {
				jmp = sPos
			}

			if text[sPos] != pattern[sPos-pos] {
				isMismatch = true
				break
			}
		}
		if !isMismatch {
			occurrences = append(occurrences, pos)
		}

		if jmp > 0 {
			// jump directly to next potential matched position
			pos = jmp - 1
		}
	}
	return
}
