package main

//import (
//	"fmt"
//	"os"
//)
//
//func main() {
//	s := ""
//	//如果没有_,会报错，range产生一对值，要用一对值接受，否则对应不上
//	//go不允许无用的临时变量的存在，编译会报错，保证代码的简洁，提供了_（空标识符）的方式忽略参数
//	for _, arg := range os.Args[1:] {
//		s += arg
//	}
//	fmt.Println(s)
//}
