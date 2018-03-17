package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	var n int
	var input string

	fmt.Scanf("%d", &n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	parts := strings.Split(input, " ")

	for i := n-1; i>0; i--{
		fmt.Print(parts[i] + " ")
	}
	// To avoid extra space at the end
	fmt.Print(parts[0])
}