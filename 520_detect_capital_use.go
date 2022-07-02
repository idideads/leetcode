package leetcode

func detectCapitalUse(word string) bool {
	if len(word) < 2 {
		return true
	}
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	capitalMap := make(map[byte]struct{}, 26)
	for i := range alphabet {
		capitalMap[alphabet[i]] = struct{}{}
	}
	_, fristCapital := capitalMap[word[0]]
	_, secondCapital := capitalMap[word[1]]
	if !fristCapital && secondCapital {
		return false
	}
	for i := 2; i < len(word); i++ {
		_, currentCapital := capitalMap[word[i]]
		switch {
		case fristCapital && secondCapital && !currentCapital:
			return false
		case fristCapital && !secondCapital && currentCapital:
			return false
		case !fristCapital && !secondCapital && currentCapital:
			return false
		}
	}
	return true
}
