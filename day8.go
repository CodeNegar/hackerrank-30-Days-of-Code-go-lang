package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	n := 0
	input := ""
	book := make(map[string] string)
	fmt.Scanf("%d", &n)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i<n; i++{
		scanner.Scan()
		input = scanner.Text()
		parts := strings.Split(input, " ")
		book[parts[0]] = parts[1]
	}

	for{
		scanner.Scan()
		input = scanner.Text()
		if input == ""{
			break;
		}
		result, exists := book[input]
		if exists{
			fmt.Println(input + "=" + result)
		}else{
			fmt.Println("Not found")
		}
	}
}