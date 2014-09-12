// Arrangement
package arrangement

import (
	"fmt"
	"os"
	"regexp"
)

func swap(str []byte, a int, b int) {
	c := str[a]
	str[a] = str[b]
	str[b] = c
}

//k 表示当前选取到第几个数，m表示共有多少个数
func allRange(str []byte, k int, m int) {
	if k == m {
		regchar := regexp.MustCompile("[\\D]")
		regnum := regexp.MustCompile("[\\d]")
		if regchar.Find(str) != nil && regnum.Find(str) != nil {
			fmt.Println(string(str))
			writefile(str)
			writefile([]byte("\r\n"))
		}

	} else {
		for i := k; i <= m; i++ {
			swap(str, k, i)
			allRange(str, k+1, m)
			swap(str, k, i)
		}
	}

}

func createfile() (fout *os.File) {
	file := "allpasswd.txt"
	_, err := os.Stat(file)
	if err != nil && os.IsNotExist(err) {
		_, err1 := os.Create(file)
		if err1 != nil {
			fmt.Println(file, err1)
			return
		}
	}
	fd, _ := os.OpenFile(file, os.O_APPEND, 0777)

	return fd
}

func writefile(by []byte) {
	fout := createfile()
	fout.Write(by)
	defer fout.Close()
}

func Foo(str []byte) {
	allRange(str, 0, len(str)-1)
}
