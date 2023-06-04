package examples

import (
	"fmt"
	"time"
)

type Stack struct {
	Data []interface{}
	Size int
}

func (s *Stack) Len() int {
	return len(s.Data)
}

func (s *Stack) Push(value interface{}) {
	s.Data = append(s.Data, value)
}

func (s *Stack) Pop() {
	if s.Len() > 0 {
		s.Data = s.Data[0 : s.Len()-1]
		fmt.Println(s.Data)
	}
}

func Example4() {
	var s Stack
	var number int
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	fmt.Println("Enter the number of producers you want: ")
	fmt.Scanf("%d", &number)

	fmt.Println("Now we Produce :")
	start := time.Now()

	go Producer(&s, c1, c2, number)
	go Consumer(&s, c1, c2, c3)

	<-c3
	elapsed := time.Since(start)
	fmt.Println("Time taken is: ", elapsed)
}

func Producer(s *Stack, c1, c2 chan string, number int) {
	for i := 1; i < number; i++ {
		s.Push(i)
		c1 <- "produced"
		<-c2
	}
	close(c1)
}

func Consumer(s *Stack, c1, c2, c3 chan string) {
	for range c1 {
		c2 <- "Consumed"
	}
	close(c2)
	c3 <- "Completed"
}
