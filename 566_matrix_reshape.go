package leetcode

func matrixReshape(mat [][]int, r int, c int) [][]int {
    m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}

	tmpSlice := make([]int, m*n)
	index := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tmpSlice[index] = mat[i][j]
			index++
		}
	}

	reShape := make([][]int, r)
	for i := range reShape {
		reShape[i] = make([]int, c)
	}

	index = 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			reShape[i][j] = tmpSlice[index]
			index++
		}
	}
	return reShape
}