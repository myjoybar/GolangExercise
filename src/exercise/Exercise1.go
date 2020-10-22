package exercise
import (
	"fmt"       // 导入 fmt 包，打印字符串是需要用到
)
var (
	a int
	b string
	c []float32
	d func() bool
	e struct {
		x int
	}
)
func print()  {
	x1:=100
	a1,s:=1, "abc" //简短格式
	fmt.Println(a)
	fmt.Println(x1)
	fmt.Println(a1)
	fmt.Println(s)
}
