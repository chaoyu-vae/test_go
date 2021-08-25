package main

import (
	"container/list"
	"fmt"
	"sync"
)

func pipei_20(s string) bool {
	var stack = list.New()
	var size = len(s)
	for index := 0; index < size; index++ {
		value := string(s[index])
		if value == "[" || value == "{" || value == "(" {
			stack.PushBack(value)
		} else if value == "]" && stack.Back().Value == "[" {
			stack.Remove(stack.Back())
		} else if value == "}" && stack.Back().Value == "{" {
			stack.Remove(stack.Back())
		} else if value == ")" && stack.Back().Value == "(" {
			stack.Remove(stack.Back())
		} else {
			stack.PushBack(value)
		}
	}
	if stack.Len() == 0 {
		return true
	} else {
		return false
	}
}

func maopaopaixu(arr []int) []int {
	var size = len(arr)
	for i := 1; i < size; i++ {
		for j := 0; j < size-i; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}

		}
	}
	return arr
}

//合并俩个有序数组
func merge(num1 []int, m int, num2 []int, n int) []int {
	last_inex := len(num1) - 1
	var i = m - 1
	var j = n - 1
	for i >= 0 && j >= 0 {
		if num1[i] >= num2[j] {
			num1[last_inex] = num1[i]
			fmt.Printf("num1 is %v", num1)
			i--
			last_inex--
			fmt.Printf("%d %d\n", i, last_inex)
		} else {
			num1[last_inex] = num2[j]
			fmt.Printf("num1 is %v", num1)
			j--
			last_inex--
			fmt.Printf("%d %d\n", j, last_inex)
		}
	}
	for ; j >= 0; j-- {
		num1[last_inex] = num2[j]
		last_inex--
	}
	return num1
}

func sync_test() {
	var a string
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		a = "hello web"
		wg.Done()
	}()
	wg.Wait()
	fmt.Print(a)

}
func main() {
	var num1 = [...]int{0}
	var num2 = [...]int{1}
	var m = 0
	var n = 1
	var res_num = merge(num1[:], m, num2[:], n)
	fmt.Printf("result is %v", res_num)
	sync_test()
}
