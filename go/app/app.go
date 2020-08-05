package main

import (
    "fmt";
)

func helper(n int,c chan int ){
	x := (fibo(n-1) + fibo(n-2))
	c <- x
}

func fibo_chan(n int,c chan int ){
	if (n > 2){

		ci1 := make(chan int)
		ci2 := make(chan int)
		go fibo_chan(n-1, ci1)
		go fibo_chan(n-2, ci2)

		x, y := <- ci1, <- ci2
		c <- (x + y)
		return
	}
	c <- 1
	return
}

func fibo(n int) int{
	if (n > 2){
		return (fibo(n-1) + fibo(n-2))
	}
	return 1
}

func main() {
c1 := make(chan int)
c2 := make(chan int)
go helper(49, c1)
go helper(48, c2)
x := <- c1
y := <- c2
fmt.Println((x+y))
}