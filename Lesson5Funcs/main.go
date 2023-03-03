package main

import (
	"fmt"
)

/*var a []string = []string{"1", "2", "3", "4"}
var x string = "3"*/

var slice []int = []int{1, 54, 34, 7, 5}

func main() {
	//Замыкание
	/*inc := increment()
	fmt.Println(inc())*/

	// Функция, где строка x сравнивается со значениями a путем форенджа слайса a и if
	/*searchX := contains(a, x)
	fmt.Println(searchX)*/

	//Поиск максимального значения в слайсе интов
	maxValue := getMax(slice)
	fmt.Println(maxValue)
}

func increment() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func contains(a []string, x string) bool {
	for _, v := range a {
		if v == x {
			fmt.Println(v)
		}
	}
	return true
}

func getMax(slice []int) int {
	max := slice[0]

	for _, v := range slice {
		if max < v {
			max = v
		}
	}
	return max
}
