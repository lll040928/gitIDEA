package main

import (
	"fmt"
	"sort"
)

func main() {
	//a := make([]string, 5, 10)
	//fmt.Println(a)
	//for i := 0; i < 10; i++ {
	//	a = append(a, fmt.Sprintf("%v", i))
	//}
	//fmt.Println(a)
	//冒泡排序
	var a = [...]int{3, 7, 9, 1}
	sort.Ints(a[:])
	fmt.Println(a)
}
