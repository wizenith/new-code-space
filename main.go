package main

import (
	"fmt"

	u "github.com/rjNemo/underscore"
	"github.com/rs/xid"
)

func main() {
	// fmt.Println("main")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// filter even numbers from the slice
	evens := u.Filter(numbers, func(n int) bool { return n%2 == 0 })
	// square every number in the slice
	squares := u.Map(evens, func(n int) int { return n * n })
	// reduce to the sum
	res := u.Reduce(squares, func(n, acc int) int { return n + acc }, 0)

	fmt.Println(res) // 120

	lab()
	lab2()
	lab3()
}

func lab() {

	nums := []int{2, 4, 5, 6, 8, 0}
	is_odd := func(n int) bool { return n%2 != 0 }
	n, err := u.Find(nums, is_odd)

	fmt.Println(n)   // 5
	fmt.Println(err) // nil

}

func lab2() {
	numbers := []string{"aaa", "bbb", "ccc", "ddd", "eee"}
	fmt.Println("numbers:", numbers)
	fmt.Println("lab2")
}

func lab3() {
	for i := 0; i <= 100; i++ {
		//will print 100 IDs
		guid := xid.New()
		println("guid:", guid.String())
	}
}
