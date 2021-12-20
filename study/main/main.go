package main

import (
	"fmt"
	"os"
)

func main() {
	//go原生支持Unicode,所以可以处理所有国家的语言
	fmt.Println("hello, 世界")
	//对于变量,在没有显示初始化的情况是,数字是 0,string 是 ""
	var ans string
	sep := " "
	for k, arg := range os.Args {
		if k == 0 {
			//第一个元素是文件的名字
			continue
		}
		ans += sep + arg
	}
	fmt.Println(ans)
}

//1. 运行程序的时候使用 go run file.go 指令
//2. 当需要多次运行的时候,可以先试用 go build file.go 指令,将其编译为二进制文件,然后通过 ./file 进行执行
//3. go不需要再语句后面使用 ; 进行结尾,事实上,跟在特定符号后的换行符会被转化为分号,所以使用换行时要注意
//4. gofmt file.go 可以对文件进行标准的格式化
//5. goimport 指令可以按需求管理倒入声明的插入和移除
