package main

import (
	"fmt"
)

func main() {
	// sc := bufio.NewScanner(os.Stdin)
	// sc.Scan()
	// inputs := strings.Split(sc.Text(), " ")

	var h, w, n int
	fmt.Scanf("%d %d %d", &h, &w, &n)

	d := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	di := 0
	now := [2]int{0, 0}
	canvas := make([][]bool, h)
	for i := range canvas {
		canvas[i] = make([]bool, w)
	}

	for i := 0; i < n; i++ {
		// fmt.Println(now[0])
		// fmt.Println(now[1])
		if canvas[now[0]][now[1]] {
			di += 4
			di--
			di %= 4
		} else {
			di += 4
			di++
			di %= 4
		}
		canvas[now[0]][now[1]] = !canvas[now[0]][now[1]]
		now[0] += h + d[di][0]
		now[0] %= h
		now[1] += w + d[di][1]
		now[1] %= w
	}
	for _, ca := range canvas {
		for _, c := range ca {
			if c {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}
