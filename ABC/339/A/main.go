package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	S := sc.Text()
	parts := strings.Split(S, ".")
	fmt.Println(parts[len(parts)-1])
}
