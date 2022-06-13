package leetcode

func reverseVowels(s string) string {
    sbyte := []byte(s)
	length := len(sbyte)
	if length == 0 {
		return ""
	}
	empty := struct{}{}
	dictMap := map[byte]struct{}{
		'a': empty, 'e': empty, 'i': empty, 'o': empty, 'u': empty,
		'A': empty, 'E': empty, 'I': empty, 'O': empty, 'U': empty,
	}
	i, j := 0, length-1
	for i <= j {
		_, ok_i := dictMap[sbyte[i]]
		_, ok_j := dictMap[sbyte[j]]

		if ok_i && ok_j {
			sbyte[i], sbyte[j] = sbyte[j], sbyte[i]
			i++
			j--
		}

		if !ok_i {
			i++
		}
		if !ok_j {
			j--
		}
		
	}
	return string(sbyte)
}