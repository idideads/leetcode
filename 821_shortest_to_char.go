package leetcode

func shortestToChar(s string, c byte) []int {
	size := len(s)
	answer := make([]int, size)
	lastPosition, m1, m2 := -1, 0, 0
	for i, letter := range s {
		if byte(letter) == c {
			// -1 表示之前没找到过，需要升序计算距离
			if lastPosition == -1 {
				for j := 0; j < i; j++ {
					if byte(s[j]) != c {
						answer[j] = i - j
					}
				}
			} else {
				// 之前有找到过，从中间往两边计算距离
				m1, m2 = (lastPosition+i)/2, (lastPosition+i+1)/2
				for j, k := m1, m2; j > lastPosition && k < i; j, k = j-1, k+1 {
					if byte(s[j]) != c {
						answer[j] = j - lastPosition
					}
					if byte(s[k]) != c {
						answer[k] = i - k
					}
				}
			}
			// 刷新上次找到的位置
			lastPosition = i
		} else {
			// 往后一直找不到的，距离递增就好，一开始用负1计算的距离不用在意，会被覆盖的
			answer[i] = i - lastPosition
		}
	}

	return answer
}
