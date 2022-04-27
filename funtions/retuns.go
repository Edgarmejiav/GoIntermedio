package main

import "fmt"

func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}
func printeNames(values ...string) {

	for _, name := range values {
		fmt.Println(name)
	}

}

func getValues(x int) (double int, triple int, quad int) {

	double = x * 2
	triple = x * 3
	quad = x * 4
	return

}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
	printeNames("1, 2, 3, 4", "edgar")

	fmt.Println(getValues(2))

}
