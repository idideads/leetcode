package leetcode

func findWords(words []string) []string {
    letters := "qwertyuiopQWERTYUIOPasdfghjklASDFGHJKLzxcvbnmZXCVBNM"
	letterMap := make(map[byte]int, 26*2)
	for i := range letters {
		switch {
		case i < 20:
			letterMap[letters[i]] = 1
		case i < 38:
			letterMap[letters[i]] = 2
		default:
			letterMap[letters[i]] = 3
		}
	}
	sameRowWords := make([]string, 0)
	for i := 0; i < len(words); i++ {
		isTheSameRow := true
		rowNum := letterMap[words[i][0]]
		for j := 1; j < len(words[i]); j++ {
			if letterMap[words[i][j]] != rowNum {
				isTheSameRow = false
				break
			}
		}
		if isTheSameRow {
			sameRowWords = append(sameRowWords, words[i])
		}
	}
	return sameRowWords
}