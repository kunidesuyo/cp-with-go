package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var a, b int
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	a, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	b, _ = strconv.Atoi(sc.Text())
	for i := 0; i <= 9; i++ {
		if i != a+b {
			fmt.Println(i)
			return
		}
	}
}
