package main

import "fmt"

type stack struct {
	i    int
	data [10]int
	x    map[string]int
}

func (s *stack) init() {
	s.x = make(map[string]int)
}
func (s *stack) push(k int) {
	//s.x = make(map[string]int)
	s.data[s.i] = k
	s.i++
}
func (s *stack) set(str string, v int) {
	s.x[str] = v
}

func (s *stack) pop() int {
	s.i--
	ret := s.data[s.i]
	s.data[s.i] = 0
	return ret
}
func main() {
	fmt.Println("Hello World")
	var s stack
	s.init()
	s.push(25)
	s.push(14)
	s.set("name", 20)
	s.set("place", 30)
	fmt.Printf("stack %v\n", s)
}
