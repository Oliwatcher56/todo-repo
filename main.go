package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func main() {

	result := sum(3, 5)

	fmt.Println("Сумма:", result)

}
