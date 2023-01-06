package main

import (
	"flag"
	"fmt"
)

// 1. run in cli
// $ go run main.go --mode=fast

var mode = flag.String("mode", "", "process mode")

// var mode = flag.String("mode", "", "process mode -help")

// var mode = flag.String("mode", "", "-help")

// 1. flag 底层知道怎么解析命令行，
// 2. 并且将值赋给 mode*string 指针
// 3. 通过自己注册的这个 mode 指针获取到最终的值
func main() {
	// 解析命令行参数
	flag.Parse()

	// 输出命令行参数
	fmt.Println(*mode)
}
