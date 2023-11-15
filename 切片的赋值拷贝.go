package main

import "fmt"

func main() {
	//a := make([]int, 3)
	//b := a
	//b[0] = 100
	//fmt.Println(a)
	//fmt.Println(b)
	//切片的遍历
	a := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}
	fmt.Println()
	for _, value := range a {
		fmt.Println(value)
	}
	//apend()方法，扩容切片
	var b []int //此时并未申请内存
	for i := 0; i < 10; i++ {
		b = append(b, i)
		fmt.Printf("%d len:%d cap:%d ptr:%p\n", b, len(b), cap(b), b)
	}
}
