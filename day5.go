package main
import "fmt"

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	for i := 1; i <= 10; i++{
		fmt.Println(n, "x", i, "=", i*n)
	}
}