package main

import "fmt"

func main() {
	/*var a []string
	var b []int

	var c = []bool{false, true}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)*/
	a := [5]int{55, 56, 57, 58, 59}
	b := a[1:2]

	fmt.Println(b)
	fmt.Printf("%T\n", b)
	c := b[0:len(b)]
	fmt.Printf("%T\n", c)
	fmt.Println(c)
	//make 函数构造切片
	d := make([]int, 5, 10)
	fmt.Println(d)
	//通过len()函数获取切片的长度，通过cap()获取切片的容量
	fmt.Println(len(d))
	fmt.Println(cap(d))
	//与nil值切片比较
	var m []int     //声明类型
	var E = []int{} //声明并且初始化
	f := make([]int, 0)
	fmt.Println(m, len(m), cap(m))
	fmt.Println(E, len(E), cap(E))
	fmt.Println(f, len(f), cap(f))
	if m == nil {
		fmt.Println("a是一个nil")
	}
	if E == nil {
		fmt.Println("E是一个nil")
	}
	if f == nil {
		fmt.Println("f是一个nil")
	}
}
