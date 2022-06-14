package leetcode

func canConstruct(ransomNote, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	magazineMap := make(map[rune]int)
	for _, runeMagazine := range magazine {
		magazineMap[runeMagazine]++
	}
	for _, runeRansomNote := range ransomNote {
		if _, ok := magazineMap[runeRansomNote]; ok {
			magazineMap[runeRansomNote]--
		} else {
			return false
		}
	}
	for _, n := range magazineMap {
		if n < 0 {
			return false
		}
	}
	return true
}
