package main

//import (
//	"fmt"
//	"io/ioutil"
//	"os"
//	"strings"
//)
//
//func main() {
//	counts := make(map[string]int)
//	for _, arg := range os.Args[1:] {
//		data, err := ioutil.ReadFile(arg)
//		if err != nil {
//			fmt.Fprintf(os.Stderr, "dup3 error:%v\n", err)
//			continue
//		}
//		for _, line := range strings.Split(string(data), "\n") {
//			counts[line]++
//		}
//		for key, n := range counts {
//			if n > 1 {
//				fmt.Printf("%s\t%d\n", key, n)
//			}
//		}
//	}
//}
