package leetcode

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
*/
func isAnagram(s string, t string) bool {
	size, size2 := len(s), len(t)
	if size != size2 {
		return false
	}
	mapS, mapT := make(map[byte]int, size), make(map[byte]int, size)
	for i := 0; i < size; i++ {
		mapS[s[i]]++
		mapT[t[i]]++
	}
	for k, v := range mapS {
		byteT, ok := mapT[k]
		if !ok || byteT != v {
			return false
		}
	}
	return true
}
