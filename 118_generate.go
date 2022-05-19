package leetcode

func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i := range triangle {
		triangle[i] = make([]int, i+1)
	}
	triangle[0][0] = 1
	if numRows > 1 {
		triangle[1][0], triangle[1][1] = 1, 1
	}
	if numRows > 2 {
		for i := 2; i < numRows; i++ {
			for j := 0; j < len(triangle[i]); j++ {
				if j == 0 || j == len(triangle[i])-1 {
					triangle[i][j] = 1
				} else {
					triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
				}
			}
		}
	}
	return triangle
}
