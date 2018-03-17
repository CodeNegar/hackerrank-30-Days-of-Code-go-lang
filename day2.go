package main

import(
	"fmt"
	"math"
)

func main(){
	var mealCost float64
	var tipPercent float64
	var taxPercent float64
	var total float64

	fmt.Scanf("%f", &mealCost)
	fmt.Scanf("%f", &tipPercent)
	fmt.Scanf("%f", &taxPercent)
	total = (mealCost*(tipPercent/100)) + (mealCost*(taxPercent/100)) + mealCost

	if total - math.Floor(total) > 0.5{
		total = math.Ceil(total)
	}else{
		total = math.Floor(total)
	}

	fmt.Println("The total meal cost is", total, "dollars.")
}