package stack

import (
	"fmt"
)

func ExampleStack() {
	s := Stack[string]{}
	s.Push("foo")
	s.Push("bar")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	//output:
	//bar true
	//foo true
	//  false
}
