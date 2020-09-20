// Go-Routine and Channels

package concurrency

import ("fmt"; "time")

func Ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func Pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func Printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func RoutineCall() {
	go f(0)
	var input string
	fmt.Scanln(&input)
}
