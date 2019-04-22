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
	fmt.Println(currency(rub))
	fmt.Println("Введите длины катетов в сантиметрах через пробел:")
	fmt.Scanln(&kate1, &kate2)
	println("Площадь:", ploshad(kate1, kate2), "квадртаных см")
	println("Длина гипотенузы", gipotenuza(kate1, kate2),"см, Длина периметра", perimeter(kate1, kate2), "см")

}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num * output)) / output
}

func currency(rub float64) float64 {
	curs := 65.5
	var usd float64
	usd = rub / curs
	return toFixed(usd, 2)
}

func ploshad (kate1 float64, kate2 float64) float64 {
	return kate1*kate2/2
}

func perimeter (kate1 float64, kate2 float64) float64 {
	return  toFixed(gipotenuza(kate1, kate2) + kate1 + kate2, 2)
}

func gipotenuza (kate1, kate2 float64) float64{
	return toFixed(math.Sqrt(kate1*kate1 + kate2*kate2), 2)
}