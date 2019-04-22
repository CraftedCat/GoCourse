package main

import (
	"fmt"
	"math"
)

func main() {
	var rub float64
	var kate1 float64
	var kate2 float64
	fmt.Println("Введите ваше кол-во рублей")
	fmt.Scanln(&rub)
	fmt.Printf("Кол-во долларов:%.2f \n", currency(rub))
	fmt.Println("Введите длины катетов в сантиметрах через пробел:")
	fmt.Scanln(&kate1, &kate2)
	fmt.Printf("Площадь :%.2f кв. см", ploshad(kate1, kate2))
	fmt.Printf("\nДлина гипотенузы: %.2f см, \nДлина периметра: %.2f см", gipotenuza(kate1, kate2), perimeter(kate1, kate2))
}

func currency(rub float64) float64 {
	const curs = 65.5
	var usd float64
	usd = rub / curs
	return usd
}

func ploshad(kate1 float64, kate2 float64) float64 {
	return kate1 * kate2 / 2
}

func perimeter(kate1 float64, kate2 float64) float64 {
	return gipotenuza(kate1, kate2) + kate1 + kate2
}

func gipotenuza(kate1, kate2 float64) float64 {
	return math.Sqrt(math.Pow(kate1, 2) + math.Pow(kate2, 2))
}
