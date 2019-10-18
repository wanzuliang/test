package main

import (
	"fmt"
	"math/rand"
	"time"
	// "reflect"
)

func main() {

	var list [] int

	list_add(&list, 10000, 0, 10000)
	// list_show(&list)


	// T := sort_insert(&list)
	// T := sort_bubble(&list)
	T := sort_select(&list)
	
	fmt.Println(T)

	// list_show(&list)
}
func list_show(list *[]int) {
	var L = *list
	fmt.Println(L)
	// for k, value := range *list  {
	// 	fmt.Println(k, ":" , value)
	// }
}

func list_add(list *[]int, len, min, max int) {
	//
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		*list = append(*list, rand.Int() % (max-min+1)+min)
	}
}

func sort_insert(list *[]int) time.Duration{
	//
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
	//
	L := *list
	len := len(L)

	start := time.Now()
	for i := 1; i < len; i++ {
		for j := 0; j < i; j++ {
			if L[i] < L[j] {
				temp := L[j]
				L[j] = L[i]
				L[i] = temp
			}
		}
	}
	end := time.Now()
	return end.Sub(start)
}