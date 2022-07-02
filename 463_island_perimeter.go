package leetcode

func islandPerimeter(grid [][]int) int {
    perimeter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				perimeter += 4
				if j < len(grid[i])-1 && grid[i][j+1] == 1 {
					perimeter -= 2
				}
				if i < len(grid)-1 && grid[i+1][j] == 1 {
					perimeter -= 2
				}
			}
		}
	}
	return perimeter
}