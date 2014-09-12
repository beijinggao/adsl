package test

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func (self *Student) SInfo() {
	fmt.Println(self.Name, self.Age)
}

type GradStudent struct {
	Name string
	Age  int
}

func (self *GradStudent) SInfo() {
	fmt.Println(self.Name, self.Age)
}
