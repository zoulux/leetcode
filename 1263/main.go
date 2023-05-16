package main

func main() {

}

func minPushBox(grid [][]byte) int {
	m, n := len(grid), len(grid[0]) // m 行 n 列
	var si, sj, bi, bj int          // 玩家的起始位置，箱子的起始位置
	for i, row := range grid {
		for j, c := range row {
			if c == 'S' {
				si, sj = i, j
			} else if c == 'B' {
				bi, bj = i, j
			}
		}
	}

	check := func(i, j int) bool {
		// 如果没有越界，如果没有撞墙
		return i >= 0 && i < m && j >= 0 && j < n && grid[i][j] != '#'
	}
	q := [][5]int{{si, sj, bi, bj, 0}} // 开始人的位置，箱子的位置，步骤为0
	vis := make(map[[4]int]bool)       // 访问数组
	vis[[4]int{si, sj, bi, bj}] = true //其实状态位置反问过
	dirs := [5]int{-1, 0, 1, 0, -1}    // 方向数组
	for len(q) > 0 {
		p := q[0]                                         // 队列的头部
		q = q[1:]                                         //
		si, sj, bi, bj, d := p[0], p[1], p[2], p[3], p[4] // 还原位置
		if grid[bi][bj] == 'T' {                          // 找到 T
			return d
		}
		for k := 0; k < 4; k++ {
			sx, sy := si+dirs[k], sj+dirs[k+1] //人的方向
			if !check(sx, sy) {
				// 位置无效
				continue
			}
			if sx == bi && sy == bj {
				// 人的位置与箱子的位置重叠则算推动一次
				bx, by := bi+dirs[k], bj+dirs[k+1] // 箱子的方向
				if !check(bx, by) || vis[[4]int{sx, sy, bx, by}] {
					// 如果箱子的位置有效，并且之前没有访问过
					continue
				}

				vis[[4]int{sx, sy, bx, by}] = true           // 人的位置，箱子的位置 访问过
				q = append(q, [5]int{sx, sy, bx, by, d + 1}) // 人和箱子都移动到新位置，看下继续怎么走
			} else if !vis[[4]int{sx, sy, bi, bj}] {
				vis[[4]int{sx, sy, bi, bj}] = true
				q = append([][5]int{{sx, sy, bi, bj, d}}, q...) // 人动，箱子不动，推动次数不动
			}
		}
	}
	return -1
}

func minPushBox2(grid [][]byte) int {
	m, n := len(grid), len(grid[0]) // m 行 n 列
	var si, sj, bi, bj int          // 玩家的起始位置，箱子的起始位置
	for i, row := range grid {
		for j, c := range row {
			if c == 'S' {
				si, sj = i, j
			} else if c == 'B' {
				bi, bj = i, j
			}
		}
	}

	// 将二维坐标转成一维坐标
	f := func(i, j int) int {
		return i*n + j
	}
	check := func(i, j int) bool {
		// 如果没有越界，如果没有撞墙
		return i >= 0 && i < m && j >= 0 && j < n && grid[i][j] != '#'
	}
	q := [][3]int{{f(si, sj), f(bi, bj), 0}} // 开始人的位置，箱子的位置，步骤为0
	vis := make([][]bool, m*n)               // 访问数组
	for i := range vis {
		vis[i] = make([]bool, m*n)
	}
	vis[f(si, sj)][f(bi, bj)] = true //其实状态位置反问过
	dirs := [5]int{-1, 0, 1, 0, -1}  // 方向数组
	for len(q) > 0 {
		p := q[0]                                       // 队列的头部
		q = q[1:]                                       //
		si, sj, bi, bj = p[0]/n, p[0]%n, p[1]/n, p[1]%n // 还原位置
		d := p[2]
		if grid[bi][bj] == 'T' { // 找到 T
			return d
		}
		for k := 0; k < 4; k++ {
			sx, sy := si+dirs[k], sj+dirs[k+1] //人的方向
			if !check(sx, sy) {
				// 位置无效
				continue
			}
			if sx == bi && sy == bj {
				// 人的位置与箱子的位置重叠则算推动一次
				bx, by := bi+dirs[k], bj+dirs[k+1] // 箱子的方向
				if !check(bx, by) || vis[f(sx, sy)][f(bx, by)] {
					// 如果箱子的位置有效，并且之前没有访问过
					continue
				}
				vis[f(sx, sy)][f(bx, by)] = true                   // 人的位置，箱子的位置 访问过
				q = append(q, [3]int{f(sx, sy), f(bx, by), d + 1}) // 人和箱子都移动到新位置，看下继续怎么走
			} else if !vis[f(sx, sy)][f(bi, bj)] {
				vis[f(sx, sy)][f(bi, bj)] = true
				q = append([][3]int{{f(sx, sy), f(bi, bj), d}}, q...) // 人动，箱子不动，推动次数不动
			}
		}
	}
	return -1
}
