package main

func main() {

}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	count := 0

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0' // 标记访问过

		// 上下左右四个方向
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				dfs(r, c)
			}
		}
	}

	return count
}

func numIslands2(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	count := 0

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				grid[r][c] = '0'
				queue := [][]int{{r, c}}

				for len(queue) > 0 {
					cell := queue[0]
					queue = queue[1:]

					for _, d := range dirs {
						nr, nc := cell[0]+d[0], cell[1]+d[1]
						if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '1' {
							grid[nr][nc] = '0'
							queue = append(queue, []int{nr, nc})
						}
					}
				}
			}
		}
	}

	return count
}
