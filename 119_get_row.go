package leetcode

func getRow(rowIndex int) []int {
	return generate(rowIndex + 1)[rowIndex]
}
