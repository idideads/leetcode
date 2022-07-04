package leetcode

func reverseStr(s string, k int) string {
	bytes := []byte(s)
	if len(s) <= k {
		for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
			bytes[i], bytes[j] = bytes[j], bytes[i]
		}
		return string(bytes)
	}
	index := 0
	for index+k < len(bytes) {
		bytes2 := bytes[index : index+k]
		for i, j := 0, len(bytes2)-1; i < j; i, j = i+1, j-1 {
			bytes2[i], bytes2[j] = bytes2[j], bytes2[i]
		}

		index += 2 * k
	}
	// fmt.Printf("index=%d\n", index)
	if index < len(bytes) {
		if index+k > len(bytes) {
			bytes2 := bytes[index:]
			for i, j := 0, len(bytes2)-1; i < j; i, j = i+1, j-1 {
				bytes2[i], bytes2[j] = bytes2[j], bytes2[i]
			}
		} else {
			bytes2 := bytes[index : index+k]
			for i, j := 0, len(bytes2)-1; i < j; i, j = i+1, j-1 {
				bytes2[i], bytes2[j] = bytes2[j], bytes2[i]
			}
		}
	}
	return string(bytes)
}
