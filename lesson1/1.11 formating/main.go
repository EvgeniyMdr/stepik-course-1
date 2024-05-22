package main

import "fmt"

func main() {
	var num float64
	const max = 10000

	fmt.Scan(&num)
	if num > max {
		fmt.Printf("%g", num)
	} else if num <= 0 {
		fmt.Printf("число %2.2f не подходит", num)
	} else {
		fmt.Printf("%.4f", num*num)
	}
}
