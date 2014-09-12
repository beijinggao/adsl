package adsl

import (
	"arrangement"
)

var comb []byte = make([]byte, 36)
var buf = make([]byte, 4)

//C(m n) 长度
func length(m int, n int) int {
	d := m - n + 1
	facT := fac(m, n, d)
	facN := fac1(n)
	return facT / facN

}

func fac(m int, n int, d int) int {
	if m == d {
		return m
	} else {
		return fac(m-1, n, d) * m
	}
}

func fac1(n int) int {
	if n == 1 {
		return n
	} else {
		return fac1(n-1) * n
	}
}

func Combination(src []byte, m int, n int, d int) {
	//a:=length(m,n)
	for i := m; i >= n; i-- {
		// 选择当前的头元素
		comb[n] = src[i-1]
		if n > 1 {
			//进入下一次更小的组合问题
			Combination(src, i-1, n-1, d)
		} else {
			//满了需要的组合数，输出
			for j := d; j > 0; j-- {
				//fmt.Print(string(comb[j]))
				buf[j-1] = comb[j]
			}
			//fmt.Print("  ")
			arrangement.Foo(buf)
			//writefile(buf)
			//writefile([]byte("\r\n"))
		}
	}
}
