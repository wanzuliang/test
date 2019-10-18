package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var list [] int
	var list2 [] int
	
	list_add(&list, 10000, 0, 9999)
	list_add(&list2, 10000, 0, 9999)

	sort_insert(&list)
	sort_bubble(&list2)

	// fmt.Println(list)

	// for _, value := range *list {
	// 	if true || value == 20 || value == 30{
	// 		fmt.Print(value, " ")
	// 	}
	// }
	for i := 0; i < 100; i++ {
		fmt.Print(list[i] , " ")
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	for i := 0; i < 100; i++ {
		fmt.Print(list2[i] , " ")
	}
}

func list_add(list *[]int, len, min, max int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		*list = append(*list, rand.Int() % (max-min+1)+min)
	}
}

func sort_insert(list *[]int) {
	len := len(*list)
	for i := 1; i < len; i++ {
		for j := 0; j < i; j++ {
			if (*list)[i] < (*list)[j] {
				temp := (*list)[i]
				for k := i; k > j; k-- {
					(*list)[k] = (*list)[k-1]
				}
				(*list)[j] = temp
			}
		}
	}

	// for i := 0; i < len; i++ {
	// 	fmt.Print( (*list)[i] ," AAA ")
	// }
}

func sort_bubble(list *[]int) {
	len := len(*list)
	for i := 0; i < len; i++ {
		min := (*list)[i]
		num := i
		
		for j := i+1; j < len; j++ {
			if (*list)[j] < min {
				min = (*list)[j]
				num = j
			}
		}
		(*list)[num] = (*list)[i]
		(*list)[i] = min
	}

	// for i := 0; i < len; i++ {
	// 	fmt.Print( (*list)[i] ," BBB ")
	// }
}