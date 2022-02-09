package main

//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//func main() {
//	counts := make(map[string]int)
//	files := os.Args[1:]
//	if len(files) == 0 {
//		count(os.Stdin, counts)
//	} else {
//		//循环读取文件
//		for _, arg := range files {
//			f, err := os.Open(arg)
//			if err != nil {
//				fmt.Fprintf(os.Stdout, "dup2 error:%v\n", err)
//				continue
//			}
//			count(f, counts)
//			f.Close()
//		}
//	}
//}
//
//func count(stdin *os.File, counts map[string]int) {
//	intput := bufio.NewScanner(stdin)
//	for intput.Scan() {
//		counts[intput.Text()]++
//	}
//	//输出统计数据
//	for _, n := range counts {
//		if n > 1 {
//			fmt.Printf("%s\n", stdin.Name())
//			break
//		}
//	}
//}
