package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	var table [6][6]int64
	var input string
	var max , temp int64
	scanner := bufio.NewScanner(os.Stdin)
	max = -99

	for i :=0; i<6; i++{
		scanner.Scan()
		input = scanner.Text()
		row := strings.Split(input, " ")
		for j := 0; j<6; j++{
			row[j] = strings.TrimSpace(row[j])
			table[i][j], _ = strconv.ParseInt(row[j], 10, 64)
		}
	}

	for i :=0; i<4; i++{
		for j :=0; j<4; j++{
			temp = table[i][j] + table[i][j+1] + table[i][j+2] + table[i+1][j+1] + table[i+2][j] + table[i+2][j+1] + table[i+2][j+2]

			if temp > max {
				max = temp
			}
		}
	}

	fmt.Println(max)
}