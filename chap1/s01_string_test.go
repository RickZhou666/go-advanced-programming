package chap1

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestArrayStringSlice(t *testing.T) {
	// 字符串的切片操作
	fmt.Println("==========")
	s := "hello, world"
	hello := s[:5] // s[start:end)
	world := s[7:]
	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]
	fmt.Println("string 1: ", s, "\n",
		"string 2: ", hello, "\n",
		"string 3: ", world, "\n",
		"string 4: ", s1, "\n",
		"string 5: ", s2)

	// Printf format printout https://pkg.go.dev/fmt
	fmt.Printf("string 1: %v", s)
	fmt.Println()
	fmt.Printf("string 1: %#v", s)
	fmt.Println()
	fmt.Printf("string 1: %T", s)
	fmt.Println()
	fmt.Printf("string 1: %X", s)
	fmt.Println()

	// 字符串内置 Len() 函数
	fmt.Println("==========")
	fmt.Println("len(s): ", ((*reflect.StringHeader)(unsafe.Pointer(&s)).Len))
	fmt.Println("len(s1): ", ((*reflect.StringHeader)(unsafe.Pointer(&s1)).Len))
	fmt.Println("len(s2): ", ((*reflect.StringHeader)(unsafe.Pointer(&s2)).Len))

	// 打印转型为字节类型来查看字符底层对应的数据
	fmt.Println("==========")
	fmt.Printf("%#v\n", []byte("hello, 世界"))

	// 指定UTF8编码后的值
	fmt.Println("==========")
	fmt.Println("\xe4\xb8\x96")
	fmt.Println("\xe7\x95\x8c")

	// UTF8 错误编码不会向后扩展
	fmt.Println("==========")
	fmt.Println("\xe4\x00\x00\xe7\x95\x8cabc")

	// 遍历原始字节码
	fmt.Println("==========")
	for i, c := range []byte("世界abc") {
		fmt.Println(i, c)
	}

	// 遍历字符串的字节数组
	fmt.Println("==========")
	const s3 = "\xe4\x00\x00\xe7\x95\x8cabc"
	for i := 0; i < len(s3); i++ {
		fmt.Printf("%d %x\n", i, s[i])
	}

	// []rune
	fmt.Println("==========")
	fmt.Printf("%#v\n", []rune("世界"))
	fmt.Println()
	fmt.Printf("%#v\n", string([]rune{'世', '界'}))

}
