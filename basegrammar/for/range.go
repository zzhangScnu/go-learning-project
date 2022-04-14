package _for

import (
	"fmt"
	"time"
)

type example struct {
	num  int
	name string
}

func RangeStruct() {
	arr := []example{
		{num: 1, name: "a"}, // 字段多的时候这样写比较好
		{2, "b"},
	}
	// 输出索引
	for i := range arr {
		fmt.Println(i)
	}
	// 输出值
	for _, v := range arr {
		fmt.Println(v)
	}
	// i和v是特定的中间变量，在整个遍历过程中重复使用。range的过程中会将arr的值进行拷贝
	// 一般情况下，只获取i，然后根据下标索引arr[i]会减少赋值，提高性能
	// 但是map是随机遍历，如果获取i再去hash获取map[i]，性能不一定更好
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func RangeStructConcurrently() {
	arr := []example{
		{num: 1, name: "a"},
		{2, "b"},
	}
	go func() {
		arr[1].name = "modified"
		fmt.Println("modified success")
		arr = append(arr, example{3, "c"})
		fmt.Println("add success")
	}()
	// 循环次数一开始就定好了，所以中途新增的元素是不会被遍历到的（因为append在末尾）
	// 但是并发修改元素是会实时遍历到的，因为还是基于原切片的遍历，而不是原切片的拷贝
	for i, v := range arr {
		fmt.Println(i, v)
		time.Sleep(time.Second * 5)
	}
	fmt.Println("newest arr is: ")
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func RangeString() {
	for i, c := range "collection" {
		// c是ASCII码，直接输出会显示数字
		fmt.Printf("%d, %c \n", i, c)
	}
}

func RangeMap() {
	m := map[string]string{
		"a": "1",
		"b": "2",
	}
	// 随机顺序遍历
	for i, v := range m {
		fmt.Println(i, v)
	}
}
