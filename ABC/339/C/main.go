package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	// inputs := strings.Fields(sc.Text())
	a := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		a[i], _ = strconv.Atoi(sc.Text())
	}
	now := 0
	mini := 0
	for _, v := range a {
		now += v
		mini = min(now, mini)
	}
	if mini < 0 {
		now = -mini
	} else {
		now = 0
	}
	for _, v := range a {
		now += v
	}
	fmt.Println(now)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
