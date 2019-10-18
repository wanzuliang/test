package main

import (
	"fmt"
	"math/rand"
	"time"
	// "reflect"
)

func main() {

	var list [] int

	list_add(&list, 10*1000, 0, 10)
	// list_show(&list)


	// T := sort_insert(&list)
	T := sort_bubble(&list)
	// T := sort_select(&list)
	// T := sort_mege(&list)
	
	fmt.Println(T)

	list_show(&list)
}
func list_show(list *[]int) {
	var L = *list
	fmt.Println(L)
	// for k, value := range *list  {
	// 	fmt.Println(k, ":" , value)
	// }
}

func list_add(list *[]int, len, min, max int) {
	//添加数组
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		*list = append(*list, rand.Int() % (max-min+1)+min)
	}
}

func sort_insert(list *[]int) time.Duration{
	//直接 插入
	var L = *list
	len := len(L)

	start := time.Now()
	for i := 1; i < len; i++ {
		for j := 0; j < i; j++ {
			if L[i] < L[j] {
				temp := L[i]
				for k := i; k > j; k-- {
					L[k] = L[k-1]
				}
				L[j] = temp
			}
		}
	}
	end := time.Now()
	return end.Sub(start)
}

func sort_bubble(list *[]int) time.Duration{
	//交换 冒泡
	var L = *list
	len := len(L)

	start := time.Now()
	for i := 0; i < len; i++ {
		for j := len-1; i < j; j-- {
			if L[j] < L[j-1] {
				L[j], L[j-1] = L[j-1], L[j]
			}
		}
	}
	end := time.Now()
	return end.Sub(start)
}

func sort_what(list *[]int) time.Duration{
	//
	var L = *list
	len := len(L)

	start := time.Now()
	for i := 0; i < len; i++ {
		min := L[i]
		num := i
		
		for j := i+1; j < len; j++ {
			if L[j] < min {
				min = L[j]
				num = j
			}
		}
		L[num] = L[i]
		L[i] = min
	}
	end := time.Now()
	return end.Sub(start)
}

func sort_select(list *[]int) time.Duration{
	//选择 简单
	L := *list
	len := len(L)

	start := time.Now()
	for i := 1; i < len; i++ {
		for j := 0; j < i; j++ {
			if L[i] < L[j] {
				L[i], L[j] = L[j], L[i]
			}
		}
	}
	end := time.Now()
	return end.Sub(start)
}

func sort_mege(list *[]int) time.Duration{
	//希尔
	L := *list
	len := len(L)

	start := time.Now()
	for step := len/2; 0 < step ; step /= 2 {
		for i := step; i < len; i++ {
			for j := i-step; j>=0 && L[j+step]<L[j]; j -= step {
				L[j], L[j+step] = L[j+step], L[j]
			}
		} 
	}
	end := time.Now()
	return end.Sub(start)
}