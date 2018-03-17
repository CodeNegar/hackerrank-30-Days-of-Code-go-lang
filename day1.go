package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	var num1 uint64
	var num2 float64

	fmt.Scanf("%d", &num1)
	fmt.Println(i + num1)

	fmt.Scanf("%f", &num2)
	num2 +=d
	fmt.Printf("%.1f", num2)
	fmt.Println("")

	scanner.Scan()
	fmt.Println(s + scanner.Text())
}