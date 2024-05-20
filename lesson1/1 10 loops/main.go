package main

import "fmt"

func main() {
	for i := 1; i < 11; i++ {
		fmt.Println(i * i)
	}

}

package main

import "fmt"

func main() {
	var a, b int
	var summ int
	fmt.Scan(&a, &b)

	for i := a; i <= b; i++ {
		summ += i
	}

	fmt.Println(summ)
}


package main

import "fmt"

func main() {
	var n, num, sum int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		if num >= 10 && num <= 99 && num%8 == 0 {
			sum += num
		}
	}

	fmt.Println(sum)
}


package main

import "fmt"

func main() {
	var n int
	var count int = 0
	var res = -1
	for fmt.Scan(&n); n != 0; fmt.Scan(&n){
		if (res < n) {
			res = n
			count = 1
		} else if (res == n) {
			count += 1
		}
	}
	fmt.Println(count)
}


package main

import "fmt"

func main() {
	var n, c, d int
	fmt.Scan(&n, &c, &d)

	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			return
		}
	}
}

package main

import "fmt"

func main () {
	var n int

	for fmt.Scan(&n); n != 0; fmt.Scan(&n){
		if (n < 10) {
			continue
		} else if (n > 100) {
			break
		} else {
			fmt.Println(n)
		}
	}
}

package main

import "fmt"

func main() {
	var x, p, y int
	fmt.Scan(&x, &p, &y)

	years := 0
	for x < y {
		x += (x * p) / 100
		years++
	}

	fmt.Println(years)
}

package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)

	str1 := fmt.Sprintf("%d", num1)
	str2 := fmt.Sprintf("%d", num2)

	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				fmt.Printf("%c ", str1[i])
				break
			}
		}
	}
}