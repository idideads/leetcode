package leetcode

func addStrings(num1, num2 string) string {
	answer := make([]byte, 0)
	for len(num1) > 0 || len(num2) > 0 {
		var charN byte
		switch {
		case len(num1) > 0 && len(num2) > 0:
			charN = num1[len(num1)-1] + num2[len(num2)-1] - 96
			num1 = num1[:len(num1)-1]
			num2 = num2[:len(num2)-1]
		case len(num1) > 0:
			charN = num1[len(num1)-1] - 48
			num1 = num1[:len(num1)-1]
		case len(num2) > 0:
			charN = num2[len(num2)-1] - 48
			num2 = num2[:len(num2)-1]
		}
		answer = append([]byte{charN}, answer...)
	}
	answer = append([]byte{0}, answer...)
	for i := len(answer) - 1; i > 0; i-- {
		if answer[i] > 9 {
			answer[i-1] += 1
			answer[i] += 38
		} else {
			answer[i] += 48
		}
	}
	if answer[0] == 0 {
		answer = answer[1:]
	}
	if answer[0] == 1 {
		answer[0] = 49
	}

	return string(answer)
}
