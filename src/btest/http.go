// http project http.go
package btest

import (
	"fmt"
	"itest"
	"net/http"
	"test"
)

func main4() {

	s1 := test.Student{"john", 1}

	s2 := test.GradStudent{"john grad", 30}

	itest.Test(&s1)
	itest.Test(&s2)
	itest.Test1(&s1)
	itest.Test1(&s2)

	num := 100
	exits := make([]chan bool, num)
	for i := 0; i < num; i++ {
		exits[i] = make(chan bool, 1)
		ec := exits[i]
		j := i
		go func(e chan bool) {
			fmt.Println("go routine", j)
			e <- true
		}(ec)
	}

	for i := 0; i < num; i++ {
		<-exits[i]
	}

	return
	//client := &http.Client{Transport: transport}
	resp, err := http.Get("http://www.baidu.com")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var text []byte = make([]byte, 10000)
	resp.Body.Read(text)
	fmt.Println(string(text))
}
