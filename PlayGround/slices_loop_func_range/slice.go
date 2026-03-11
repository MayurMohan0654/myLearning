package main

import (
	"fmt"
	"sort"
)

func getAddNMul(a int, b int) (sum int, pro int) {
	return a + b, a * b
}
func main() {
	var s []int = []int{}

	s = append(s, 14)
	s = append(s, 12)
	s = append(s, 15)
	s = append(s, 15)
	s = append(s, 15)

	fmt.Println(len(s))
	println(s)

	s[1] = 99

	fmt.Println(s)

	//fruites
	// var fruites []string = []string{"apple", "mango", "banana", "pizza", "burgar"}
	// var i int;
	// var v string;

	// for i, v  = range fruites{
	// 	fmt.Printf("ind: %v, fruit: %v \n", i, v)
	// }

	var fruites []string = []string{"apple", "mango", "banana", "pizza", "burgar"}

	for i, v := range fruites { // := short hand
		fmt.Printf("ind: %v, fruit: %v \n", i, v)
	}

	for i := range 5 {
		fmt.Println("ind:", i, "fruit:", fruites[i])
	}

	var num1 int = 12
	var num2 int = 5

	for range 10 {
		println("yoo")
	}

	a, p := getAddNMul(num1, num2)

	println("add:", a, "prod:", p)
	var mpp = map[string]int{"pizza": 50, "burger": 20, "coke": 30}

	// var keys []string = []string{};
	keys := []string{}

	for i := range mpp {
		keys = append(keys, i)
	}

	sort.Strings(keys)

	for _, v := range keys {
		println(v, mpp[v])

	}

}
