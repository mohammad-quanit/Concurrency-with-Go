package examples

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // sends sum to c
}

func Example2() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	go sum(s, c)

	x := <-c + <-c + <-c // receive from c
	fmt.Println(x, <-c, <-c, <-c)
}
