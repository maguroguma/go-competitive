package lake_counting

import "fmt"

var count int

func countLake(garden [][]rune) int {
	count = 0
	n, m := len(garden), len(garden[0])
	// すべてのマス目について調査
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if garden[i][j] == 'W' {
				// 池を見つけるたびにインクリメントする（DFSの中心であるここでのみインクリメントする）
				count++
				// 池を8方向に進行しながら埋め立てていく
				dfs(garden, i, j, n, m)
			}
		}
	}
	return count
}

func dfs(garden [][]rune, i, j, n, m int) {
	/* 再帰関数の終了条件: 進行先が埋め立てられているか、庭の範囲外のとき */

	// まずは現在の場所を埋め立てる
	garden[i][j] = '.'
	// 8方向に進行
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			x, y := j+dx, i+dy
			if 0 <= x && x < m && 0 <= y && y < n && garden[y][x] == 'W' {
				// 埋め立ての途中経過を表示
				showGarden(garden, n, m)
				dfs(garden, y, x, n, m)
			}
		}
	}
}

func showGarden(garden [][]rune, n, m int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%c", garden[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("--------------------\n")
}
