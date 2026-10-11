package main

import "fmt"

func estimateBrewTime(cupsQuantity int, secondsPerCup int) (totalTimeInSeconds int, info string) {
	totalTimeInSeconds = cupsQuantity * secondsPerCup
	info = "Estimated Coffee brew time:"
	return // naked return
}

func main() {
	// 12 cups, 20 seconds per cupconst
	// 12*20 = 240 seconds

	brewTime, info := estimateBrewTime(12, 20)

	fmt.Println(info, brewTime)


}
