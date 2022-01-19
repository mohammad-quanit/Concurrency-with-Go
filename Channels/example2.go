package channels

import "fmt"

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
		fmt.Println(s.Data...)
	}
}

func Producer() {

}

func Consumer() {

}
