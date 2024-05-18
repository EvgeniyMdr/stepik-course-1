package __9_conditionals

package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	switch {
	case num > 0: {
		fmt.Println("Число положительное")
	}
	case num < 0:{
		fmt.Println("Число отрицательное")
	}
	case num == 0: {
		fmt.Println("Ноль")
	}
	}
}


package main

import "fmt"

func main () {
	var num int

	fmt.Scan(&num)

	i0 := num % 10
	i1 := num % 100 / 10
	i2 := num / 100

	if (i0 != i1 && i0 != i2 && i1 != i2) {
		fmt.Println("YES")
	} else {
		fmt.Print("NO")
	}
}


package main

import "fmt"

func main () {
	var num int

	fmt.Scan(&num)

	var firstDigit int
	switch {
	case num < 10:
		firstDigit = num
	case num < 100:
		firstDigit = num / 10
	case num < 1000:
		firstDigit = num / 100
	case num < 10000:
		firstDigit = num / 1000
	default:
		firstDigit = num / 10000
	}

	fmt.Println(firstDigit)
}


package main

import "fmt"

func main () {
	var inNum int

	fmt.Scan(&inNum)

	var (
		i0 int = inNum / 100000
		i1 int = inNum % 100000 / 10000
		i2 int = inNum % 10000 / 1000
		i3 int = inNum % 1000 / 100
		i4 int = inNum % 100 / 10
		i5 int = inNum % 10
	)

	if (i0 + i1 + i2 == i3 + i4 + i5) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

package main

import "fmt"

func main () {
	var year int
	fmt.Scan(&year)

	if (year % 400 == 0 || (year % 4 == 0 && year % 100 != 0)) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}


