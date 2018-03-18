package main
import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int64
	var max string
	fmt.Scanf("%d", &n)

	bin := strconv.FormatInt(n, 2)
	parts := strings.Split(string(bin), "0")
	for _, val := range parts {
		if len(val) > len(max){
			max = val
		}
	}

	fmt.Println(len(max))
}