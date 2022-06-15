package leetcode

func findTheDifference(s, t string) byte {
	// if len(s) == 0 {
	// 	return t[0]
	// }
	countMap := make(map[byte]int)
	for i := range s {
		countMap[s[i]]++
	}
	for i := range t {
		countMap[t[i]]--
	}
	for k, v := range countMap {
		if v < 0 {
			return k
		}
	}
	return t[0]
}
