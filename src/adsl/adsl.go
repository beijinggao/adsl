// adsl project adsl.go
package adsl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func executeCmd(cmd string) string {
	output, err := exec.Command("cmd.exe", "/c", cmd).Output()
	check(err)
	return string(output)
}

func connAdsl(adslTitle string, adslName string, adslPass string) {
	adslCmd := "rasdial " + adslTitle + " " + adslName + " " + adslPass
	tempCmd := executeCmd(adslCmd)
	if strings.Index(tempCmd, "已连接") > 0 {
		fmt.Println("用户名: " + adslName + "  密码: " + adslPass)
	}

}

func cutAdsl(adslTitle string) {
	cutCmd := "rasdial " + adslTitle + " /disconnect"
	result := executeCmd(cutCmd)
	if strings.Index(result, "没有连接") != -1 {
		fmt.Println(adslTitle + " 连接不存在")
	} else {
		fmt.Println("连接已断开")
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
	userfile := "./allpasswd.txt"
	fl, err := os.Open(userfile)
	if err != nil {
		fmt.Println(userfile, err)
	}
	defer fl.Close()
	buf := bufio.NewReader(fl)
	line, _, err := buf.ReadLine()
	for ; err == nil; line, _, err = buf.ReadLine() {
		fmt.Println(string(line))
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
