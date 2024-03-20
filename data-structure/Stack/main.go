package main

import (
	"fmt"
)

type Stack struct{
	data []int
	top int
}

func createStack(capacity int) *Stack{
	return &Stack{
		data : make([]int, capacity)
		top  : -1  
	}
}
/**
	- check is the stack is full
		- if the stack data array length == top 
	- If not, 
		- increase top 1 
		- append the element in data array at top position 
*/
func (s *Stack) push(value int){
	if (len(s.data) -1) == s.top{
		fmt.Println('your stack is full')
	}
	s.top++
	append(s.data , value)
}

/**
	check if the top == -1
*/
func (s *Stack) isEmpty()bool {
	if s.top == -1 {
		return true 
	}
	return false 
}
/**
	- Check if the data is empty
	- If not, 
		- take the top element value 
		- remove the top element 
		- decrease the top by 1
*/
func (s *Stack) pop() int{
	if s.isEmpty() {
		fmt.Println('your stack is empty')
	}
	element := s.data[s.top]
	s.data = s.data[:s.top]
	s.top--
	return element 
}
func (s *stack) top(){
	fmt.Println(s.top)
}

func (s *stack) Print(){
	for i, value :=range s.data {
		fmt.Printf(" %d : %d " , i , value)
	}
}
func main(){
	stack := createStack(3)
	stack.push(1)

	s.Print()


}