package rectangle

import (
	"fmt"
)

func SliceMaps(){
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for index, val := range numbers {
		sum += val
		fmt.Print("[", index, ",", val, "] ")
	}

	fmt.Println("\nSum is :: ", sum)
	kvs := map[int]string{1: "apple", 2: "banana"}
	for k, v := range kvs {
		fmt.Println(k, " -> ", v)
	}

	str := "Hello, World!"
	ptr := &str 
	for index, c := range str {
		fmt.Print("[", index, ",", string(c), "] ")
	}
	fmt.Println("\nvalue stored at ptr address: ", *ptr)
	fmt.Println("pointer address: ", ptr)
}
