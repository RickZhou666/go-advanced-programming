package chap0

import (
	"flag"
	"fmt"
	"testing"
)

// I. approaching 1
// http://c.biancheng.net/view/21.html
// 1. 指针地址
// 2. 指针类型
// 3. 指针取值

/*
	4.1 认识指针地址和指针类型
	 在变量名前面添加&操作符（前缀）来获取变量的内存地址（取地址操作）
	 ptr := &v    // v 的类型为 T
	 其中 v 代表被取地址的变量，变量 v 的地址使用变量 ptr 进行接收，
	 ptr 的类型为*T，称做 T 的指针类型，*代表指针

	 ptr := &v ==> ptr is *T

	 提示：变量、指针和地址三者的关系是，每个变量都拥有地址，指针的值就是地址。

*
*/
func TestPtrAndAddr4_1(t *testing.T) {
	var cat int = 1
	var str string = "banana"
	fmt.Println("=============4.1======================")
	fmt.Printf("%p %p", &cat, &str) // 取变量的地址
	fmt.Println("\n====================================")
	// 0x1400 0016 2f8
}

// 4.2 从指针获取指针指向的值
// 当使用&操作符对普通变量进行取地址操作并得到变量的指针后，可以对指针使用*操作符，也就是指针取值，代码如下。
func TestPtrAndAddr4_2_getValueFromPointer(t *testing.T) {
	fmt.Println("=============4.2=======================")
	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"

	// 对字符串取地址，ptr类型为*string
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

	fmt.Println("\n====================================")
}

// 4.3 使用指针修改值
// 通过指针不仅可以取值，也可以修改值。
// *操作符的根本意义就是操作指针指向的变量。
// 		当操作在右值时，就是取指向变量的值，
// 		当操作在左值时，就是将值设置给指向的变量。

func TestPtrAndAddr4_3_ModifyValueViaPointer(t *testing.T) {
	fmt.Println("\n============4.3========================")
	// 准备两个变量，赋值1，2
	x, y := 1, 2

	// 交换变量值
	swap43(&x, &y)

	// 输出变量值
	fmt.Println(x, y)
	fmt.Println("\n====================================")
}

func swap43(a, b *int) {
	// 取a指针的值，赋给临时变量t
	t := *a // t 此时是int类型

	// 取b指针的值，赋给a指针指向的变量
	*a = *b // 左侧：此时*a 指 "a指向的变量".
	//  	   右侧：*b 取指针b的值

	// 将a指针的值赋给b指针指向的变量
	*b = t // 将 t 的值赋给指针 b 指向的变量
}

// 4.4 如果交换的是指针值
func TestPtrAndAddr4_4_ModifyValueViaPointer(t *testing.T) {
	fmt.Println("\n============4.4========================")
	x, y := 1, 2
	fmt.Println("inside main function before swap: ", &x, &y)
	swap44(&x, &y)
	fmt.Println("inside main function after swap: ", x, y)
	fmt.Println("\n====================================")
}

// 上面代码中的 swap() 函数交换的是 a 和 b 的地址，在交换完毕后，a 和 b 的变量值确实被交换。但和 a、b 关联的两个变量并没有实际关联。
// 这就像写有两座房子的卡片放在桌上一字摊开，交换两座房子的卡片后并不会对两座房子有任何影响。
func swap44(a, b *int) {
	b, a = a, b

	fmt.Println("inside swap44 function: ", a, b)
}

// 4.5 使用指针变量获取命令行的输入信息

var mode = flag.String("mode", "", "process mode")

func TestPtrAndAddr4_5_PointerToGetDataFromCLI(t *testing.T) {
	// 解析命令行参数
	flag.Parse()

	// 输出命令行参数
	fmt.Println(*mode)
}

// 4.6 创建指针的另一种方法——new() 函数
//  1. new() 函数可以创建一个对应类型的指针，
//  2. 创建过程会分配内存，
//  3. 被创建的指针指向默认值。
func TestPtrAndAddr4_6_NewAPointer(t *testing.T) {
	fmt.Println("\n============4.6========================")
	str := new(string)
	fmt.Printf("before assignment, str value is %s: \n", *str)

	*str = "Go语言教程"
	fmt.Printf("after assignment, str value is %s: \n", *str)
	fmt.Println("\n============4.6========================")
}

// ================================================================================================================================================
// ================================================================================================================================================
// ================================================================================================================================================

// II. approaching 2
// https://go.dev/tour/moretypes/1
func TestPtrAndAddr3(t *testing.T) {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

// ================================================================================================================================================
// ================================================================================================================================================
// ================================================================================================================================================

// III. approaching 3
// https://studygolang.com/articles/31495
// &和*用于指针和值的转换
// *&a == a == 3，*姑且可视为取值
func TestPtrAndAddr2(t *testing.T) {
	var a int = 3

	fmt.Printf("a stands for value: %v \n", a)

	fmt.Printf("&a stands for memory address: %v \n", &a)

	fmt.Printf("*&a stands for value of memory address: %v \n", *&a)

}

// ================================================================================================================================================
// ================================================================================================================================================
// ================================================================================================================================================

// IV. approaching 4
// csdn reference: https://blog.csdn.net/yezhijing/article/details/127585233
// fmt doc: https://pkg.go.dev/fmt
func TestPtrAndAddr(t *testing.T) {

	// 1. 当使用 & 操作符对普通变量进行取地址操作时 ==> 可以得到变量的指针
	// & 取出地址， * 根据地址取出地址指向的值。
	var myAddr = "tree road 1025, 100"

	// 对字符串取地址，ptr类型为*string
	ptr := &myAddr

	// 打印ptr的类型
	fmt.Printf("ptr的类型是: %T\n", ptr)
	// 打印ptr的指针地址
	fmt.Printf("ptr的地址是: %p\n", ptr)

	//  对指针进行取值操作
	value := *ptr

	// 打印取值后的类型
	fmt.Printf("value的类型是: %T\n", value)
	// 指针取值后就是指向变量的值
	fmt.Printf("value的值是: %s\n", value)

}

//  2. * 操作符的根本意义就是操作指针指向的变量。
// 		当操作在右值时，就是取指向变量的值，
// 		当操作在左值时，就是将值设置给指向的变量。

// a, b 类型都为*int指针类型
func swap(a, b *int) {

	// 取a指针的值，并把赋值给临时变量t，t此时是int类型
	t := *a

	// 取b指针的值，赋给a指针指向的变量，此时*a的意思不是取 a指针的值，而是 a指向的变量
	*a = *b

	// 将t的值赋给指针b 指向的变量
	*b = t

}

func TestPtrAssignment(t *testing.T) {
	x, y := 10, 20

	//  交换变量值
	swap(&x, &y)

	fmt.Println(x, y)
}

// 3. 在函数中的应用，比如返回一个结构体

// 定义一个人的结构体
type Person struct {
	Age     int
	Sex     int
	Name    string
	Address string
}

// 新建一个人													返回一个*Person的指针类型
func NewPerson(age int, sex int, name string, address string) *Person {
	var person Person

	person = Person{
		Age:     age,
		Sex:     sex,
		Name:    name,
		Address: address,
	}

	//  返回地址即可（因为指针指向的就是地址）
	return &person
}
