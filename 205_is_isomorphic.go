package leetcode

func isIsomorphic(s string, t string) bool {
	lengthS, lengthT := len(s), len(t)
	if lengthS != lengthT {
		return false
	}
	keySvalueT, keyTvalueS := make(map[byte]byte), make(map[byte]byte)
	for i := 0; i < lengthS; i++ {
		byteT, ok := keySvalueT[s[i]]
		if ok && byteT != t[i] {
			return false
		}
		keySvalueT[s[i]] = t[i]

		byteS, ok := keyTvalueS[t[i]]
		if ok && byteS != s[i] {
			return false
		}
		keyTvalueS[t[i]] = s[i]
	}
	return true
}
