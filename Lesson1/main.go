package main

import "fmt"

type AmericanVelocity float64
type EuropeanVelocity float64

func main() {

	//Вывод строки, работа с конвертацией строк
	/*fmt.Println("Hello,World!", time.Now())

	var myInt int = 35
	myString := "104"

	fmt.Println(strconv.Itoa(myInt))
	fmt.Println(strconv.Atoi(myString))*/

	//Объявление кастомных типов данных, базовая работа с функциями
	var AmV AmericanVelocity
	AmV = convA()
	fmt.Println("130 м/с:", AmV, "миль/ч")

	var EuV EuropeanVelocity
	EuV = convE()
	fmt.Println("120.4 м/с:", EuV, "км/ч")

}

func convA() (V AmericanVelocity) {
	V = (130 * 3600) / 1609
	return
}

func convE() (V EuropeanVelocity) {
	V = (120.4 * 3600) / 1000
	return
}
