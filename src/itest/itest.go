package itest

import (
	"fmt"
)

type IStudent interface {
	SInfo()
}

type ICall interface {
	SInfo()
}

func Test(s IStudent) {
	fmt.Println("======In test=====")
	s.SInfo()
}

func Test1(s ICall) {
	fmt.Println("=======In test1========")
	s.SInfo()
}
