package main

import (
	"fmt"
	"myFirstModule/dimensions"
)

func main() {
	//новые характеристики длины\высоты\ширины
	nDimensions := newDimensions(194.5, 78.7, 169.6, "cm")

	//вывод на экран характеристик
	printDimensions(nDimensions, "inch")
	printDimensions(nDimensions, "cm")

	//создание переменных с экземплярами интерфейсов для каждого авто отдельно
	bmw := dimensions.NewAuto("BMW", "X6", newDimensions(204.5, 81.3, 162.6, "cm"), 300, 340)
	mercedes := dimensions.NewAuto("Mercedes-Benz", "CLS-350", newDimensions(204.5, 81.3, 162.6, "cm"), 290, 249)
	dodge := dimensions.NewAuto("Dodge", "Grand-Caravan", newDimensions(206.5, 120.7, 189.6, "cm"), 260, 210)

	//вывод на экран характеристик авто с указанием единици измерения длины\высоты\ширины
	printAuto(bmw, "inch")
	printAuto(mercedes, "inch")
	printAuto(dodge, "cm")
}

//создание длины\ширины\высоты с указанием единицы измерения. Возвращает интерфейс
func newDimensions(h, w, l float64, u dimensions.UnitType) dimensions.Dimensions {
	u1 := dimensions.NewUnit(h, u)
	u2 := dimensions.NewUnit(w, u)
	u3 := dimensions.NewUnit(l, u)
	d1 := dimensions.NewDimensions(u1, u2, u3, u)
	return d1
}

func printDimensions(d dimensions.Dimensions, u string) {
	fmt.Printf("IN %v dimensions: height: %v, length: %v, width: %v\n",
		u, d.Height().Get(dimensions.UnitType(u)), d.Length().Get(dimensions.UnitType(u)), d.Width().Get(dimensions.UnitType(u)))
}

func printAuto(a dimensions.Auto, u string) {
	fmt.Printf("brand: %v, model: %v, max speed: %v,  engine power: %v\n dimensions: height: %v, length: %v, width: %v  %v\n",
		a.Brand(), a.Model(), a.MaxSpeed(), a.EnginePower(), a.Dimensions().Height().Get(dimensions.UnitType(u)),
		a.Dimensions().Length().Get(dimensions.UnitType(u)), a.Dimensions().Width().Get(dimensions.UnitType(u)), u)
}
