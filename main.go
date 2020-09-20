package main

import (
	"fmt"
	"math"
	rect "basics/rectangle"
	conc "basics/concurrency"
	"encoding/json"
)

func factorial(n uint) uint {
	if n ==0 {
		return 1
	}
	return n * factorial(n -1)		
}

// function definition
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total/float64(len(xs))
}

func main()  {
	// Integers floating points and String literals
	fmt.Println("result: ", 1+1.2)
	fmt.Println(len("Hello, World"))
	fmt.Println("Hello, World"[2])
	fmt.Println(`Hello, ` + `World
	this is new line
	another newline`)

	// Variables Constants
	var name string = "Jai narayan"
	fmt.Println(`my name is : ` + name)

	var x string = "hello"
	var y string = "hello"
	fmt.Println("x eq y",x==y)

	z := 5
	fmt.Println(z)

	var p = 4.56
	fmt.Println(p)

	const village string = "Payagpur"
	fmt.Println(village)

	var (
		a = 1
		b = "bee"
		c = 3
	)

	fmt.Println(b)
	fmt.Println(a+c)

	// double the number
	fmt.Print("enter a number: ")
	var number float32
	fmt.Scanf("%f", &number)
	output := number * 2
	fmt.Println(output)
	
	// Control Structure	
	if (false || true) {
		fmt.Println("true expression")
	}

	/** 
		this is for loop
	**/
	for i:=1; i <= 10; i++ {
		fmt.Println(i)
		i = i + 1
	}

	// Array Slices Maps
	var arr [5]int
	arr[4] = 100
	fmt.Println(arr)

	var slice1 = []int{1,2,3,4,5}
	slice2 := append(slice1,6,7)
	fmt.Println(slice1,slice2)

	m := make(map[string] int)
	m["key"] = 10
	fmt.Println(m)

	// function call
	xs := []float64{33,44,65,23,87,98}
	fmt.Println(math.Round(average(xs)*100)/100)

	// Closure	
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println("Closure Addition: ",add(1,2))

	// Recursion 	
	fmt.Println("factorial: ",factorial(3))

	// creating slice maps
	rect.SliceMaps()

	// creating a linked list
	rect.CreateList()
	

	// struct
	r := rect.Rectangle{0, 0, 10, 10}
	fmt.Println("rectangle area: ",r.Area())

	// strings
	byte_arr := []byte("test")
	fmt.Println("str to slice of byte: ",byte_arr)
	fmt.Println("slice of byte to str: ",string(byte_arr))

	
	// calc hashes
	conc.CalculateNonCryptoHash()
	conc.CalculateCryptoHash()

	// goroutine
	conc.RoutineCall()

	// channels
	var chn chan string = make(chan string)
	go conc.Pinger(chn)
	go conc.Printer(chn)
	go conc.Ponger(chn)
	var input string
	fmt.Scanln(&input)

	// json encoding and decoding
	type Employee struct{
		ID int
		firstName,lastName,JobTitle string
	}

	emp := Employee{
		ID: 01,
		firstName: "Jai",
		lastName: "Narayan",
		JobTitle: "gopher",
	}

	data, err := json.Marshal(emp)
	jsonStr := string(data)
	if err != nil {
		fmt.Println("err encoding value")
	} else {		
		fmt.Println("Go Type to JSON: ", jsonStr)		
	}
	
	var emp1 Employee
	btJSNStr := []byte(jsonStr)
	err1 := json.Unmarshal(btJSNStr, &emp1)
	// decoding json object
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}
	fmt.Println(emp1)
		
}
