package __5_Variables_and_arithmetic_operations

//first task
package main

import "fmt"

func main(){
	var a int
	fmt.Scan(&a)
	// здесь ваш код
	fmt.Println(a * 2 + 100)
}

//second task
package main
import "fmt"
func main(){

	var a, b int
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	fmt.Scan(&b) // считаем переменную 'b' с консоли

	a = a * a
	b = b * b
	c := a + b
	fmt.Println(c)
}

//third task
package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)
	fmt.Println(a * a)
}

//fourth task
package main

import "fmt"

func main () {
	var a int

	fmt.Scan(&a)
	fmt.Println(a % 10)
}

//fifth task
package main

import "fmt"

func main () {
	var num uint32

	fmt.Scan(&num)

	fmt.Println(num % 100 / 10)
}

//sixth task
package main

import "fmt"

func main () {
	var deg uint16

	fmt.Scan(&deg)

	h := deg * 2 / 60 // использовал magic numbers формула расчета
	m := deg * 2 % 60 // использовал magic numbers формула расчета
	fmt.Println("It is", h ,"hours", m, "minutes.")
}
