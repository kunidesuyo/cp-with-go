package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	sc := bufio.NewScanner(os.Stdin)

	var g = make([][]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		tmp := make([]int, n)
		for j := 0; j < n; j++ {
			tmp[j], _ = strconv.Atoi(inputs[j])
		}
		g[i] = tmp
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if g[i][j] == 1 {
				fmt.Printf("%d ", j+1)
			}
		}
		fmt.Println()
	}

}
