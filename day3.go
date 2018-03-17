package main
import ("fmt")

func main() {
	n := 0
	fmt.Scanf("%d", &n)

	if n%2 != 0 {
		fmt.Println("Weird")
	}else {
		if n >=2 && n <= 5 {
			fmt.Println("Not Weird")
		}
		if n >=6 && n <= 20 {
			fmt.Println("Weird")
		}
		if n >20 {
			fmt.Println("Not Weird")
		}
	}
}