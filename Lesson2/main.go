package main

import "fmt"

const P float64 = 35

const pi float64 = 3.1415

func main() {
	/*//Создание указателя на тип
	var A *int
	var B int = 43

	//Присовение указателю на тип значения интового переменной, вывод на экран разыменования
	A = &B
	fmt.Println(*A)

	//присовение простой переменной B нового значения через указатель
	*A = 56
	fmt.Println(B)*/

	//вычисление радиуса окружности
	var R float64
	R = radiusCircle()
	fmt.Printf("Радиус окружности: %.2f\n", R)

	//вычисление площади окружности с помощью указателя на R и его разыменовывание
	RadCirc := &R
	CircArea := pi * (*RadCirc * *RadCirc)
	fmt.Printf("Площадь окружности: %.2f\n", CircArea)
}

func radiusCircle() (rad float64) {
	rad = P / (pi * 2)
	return
}
