package main
/**
1. main 包是Go语言程序的入口包，一个Go语言程序必须有且仅有一个 main 包。
如果一个程序没有 main 包，那么编译时将会出错，无法生成可执行文件。

2. 包名可以与其目录名不同
3. 一个目录下的同级文件属于同一个包
 */
import (
	"fmt" // 导入 fmt 包，打印字符串是需要用到
	"math"
	"strconv"
	"time"

	//"exercise"

)
/**
1. 它是Go语言程序的入口函数，也即程序启动后运行的第一个函数。
2. main 函数只能声明在 main 包中，不能声明在其他包中，
3. 并且，一个 main 包中也必须有且仅有一个 main 函数。
 */
func main() {   // 声明 main 主函数
	fmt.Println("Hello World!") // 打印 Hello World!


	a, _ := GetData() //匿名变量的特点是一个下画线“_”
	_, b := GetData()
	fmt.Println(a, b)
	printLocalVariable()
	printFloat();
	printComplex();
	printStr()
	printMultiLineStr();
	printPointer()
	printConstant()
	printEnum()
	
}

func GetData() (int, int) {
	return 100, 200
}

func printLocalVariable()  {
	//声明局部变量 a 和 b 并赋值
	var a int = 3
	var b int = 4
	//声明局部变量 c 并计算 a 和 b 的和
	c := a + b
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	d := sum(a, b)
	fmt.Printf("main() 函数中 d = %d\n", d)
}

func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)
	num := a + b
	return num
}

/**
浮点数
 */
func printFloat()  {
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)
}

/**
complex 函数用于构建复数
 */
func printComplex()  {
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x*y)                 // "(-5+10i)"
	fmt.Println(real(x*y))           // "-5"
	fmt.Println(imag(x*y))           // "10"
}


// 如果b为真，btoi返回1；如果为假，btoi返回0
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

/**
Go语言中字符串的内部实现使用 UTF-8 编码
 */
func printStr()  {
	s := "hel" + "lo,"
	s += "world!"
	fmt.Println(s) //输出 “hello, world!”
}
/**
定义多行字符串
多行字符串一般用于内嵌源码和内嵌数据等
 */

func printMultiLineStr()  {
	const str = `第一行
第二行
第三行
\r\n
`
	fmt.Println(str)
}

/**
指针
 */
func printPointer()  {
	var cat int = 1
	var str string = "banana"
	fmt.Printf("%p %p", &cat, &str)



	//====从指针获取指针指向的值====


	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"
	// 对字符串取地址, ptr类型为*string
	ptr := &house
	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)
	// 对指针进行取值操作
	value := *ptr
	// 取值后的类型
	fmt.Printf("value type: %T\n", value)
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)

	//====创建指针的另一种方法——new() 函数====

	str1 := new(string)
	*str1 = "Go语言教程"
	fmt.Println(*str1)
}

/**
常量
 */
func printConstant()  {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"
}

/**
枚举
 */
type Weapon int
const (
	Arrow Weapon = iota    // 开始生成枚举值, 默认为0
	Shuriken
	SniperRifle
	Rifle
	Blower
)
func printEnum()  {

	// 输出所有枚举值
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)
	// 使用枚举类型并赋初值
	var weapon Weapon = Blower
	fmt.Println(weapon)
}

func convert() {
	// int 转换为字符串：Itoa()
	fmt.Println("a" + strconv.Itoa(32)) // a32

	//string 转换为 int：Atoi()
	i, _ := strconv.Atoi("3")
	fmt.Println(3 + i) // 6
	// Atoi()转换失败
	i, err := strconv.Atoi("a")
	if err != nil {
		fmt.Println("converted failed")
	}

}