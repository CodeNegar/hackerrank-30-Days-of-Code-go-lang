package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	var tests, length int
	var input, even, odd string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Scanf("%d", &tests)

	for i :=0; i<tests; i++{
		even = ""
		odd = ""
		scanner.Scan()
		input = scanner.Text()
		length = len(input)
		chars := strings.Split(input, "")

		for j:=0; j<length; j++{
			if j%2 == 0{
				even += chars[j]
			}else{
				odd += chars[j]
			}
		}

		fmt.Println(even, odd)
	}
}