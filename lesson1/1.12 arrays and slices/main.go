package __12_arrays_and_slices


// task 1
package main
import "fmt"

func main() {
	var workArray [10]uint8

	for i := 0; i < 10; i++ {
		var num uint8
		fmt.Scan(&num)
		workArray[i] = num
	}

	for i := 0; i < 3; i++ {
		var index1, index2 int
		fmt.Scan(&index1, &index2)
		workArray[index1], workArray[index2] = workArray[index2], workArray[index1]
	}

	for _, num := range workArray {
		fmt.Printf("%d ", num)
	}
}


// task 2

package main
import "fmt"

func main() {
	var N int

	fmt.Scan(&N)

	if (N >= 4) {
		var sl []int = make([]int, N , N)
		for i:=0; i< len(sl); i++ {
			fmt.Scan(&sl[i])
		}

		fmt.Println(sl[3])
	}
}


// task 3

package main
import "fmt"

func main()  {
	array := [5]int{}
	var a int
	for i:=0; i < 5; i++{
		fmt.Scan(&a)
		array[i] = a
	}
	// здесь ваш код
	var max int = array[0]

	for i:=0; i< len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}

	fmt.Println(max)
}

//task 4

package main

import "fmt"

func main () {
	var N int

	fmt.Scan(&N)

	if (1<= N && N <= 100) {
		var ls = make([]int, N,N)

		for i := 0; i<len(ls); i++ {
			fmt.Scan(&ls[i])
		}

		for i :=0; i<len(ls); i += 2 {
			fmt.Printf("%v ",ls[i])
		}
	}
}

// task 5

package main
import "fmt"

func main() {
	var  N int

	fmt.Scan(&N)

	var slice = make([]int, N, N)

	var countPositive int = 0

	if (1<= N && N <= 100) {
		for i:= 0; i < N; i++ {
			fmt.Scan(&slice[i])
		}

		for i:=0; i < len(slice); i++ {
			if (slice[i] > 0) {
				countPositive += 1;
			}
		}

		fmt.Println(countPositive)
	}
}