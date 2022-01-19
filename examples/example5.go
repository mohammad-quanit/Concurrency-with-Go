package examples

import "fmt"

func check(mychan chan string) {
	countries := []string{"Australia", "America", "Canada"}
	for i := 0; i < cap(mychan); i++ {
		mychan <- countries[i]
	}
	close(mychan)
}

// func fibonacci(n int, c chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		c <- x
// 		x, y = y, x+y
// 	}
// 	close(c)
// }

func Example5() {
	// ch := make(chan int, 10)
	// go fibonacci(cap(ch), ch)
	// for i := range ch {
	// 	fmt.Println(i)
	// }

	myChan := make(chan string, 3)
	go check(myChan)
	for i := range myChan {
		fmt.Println(i)
	}
}
