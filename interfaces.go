package figuras

import "fmt"

type Geometricas interface {
	area() float64
	perimetro() float64
}

func Medidas(g Geometricas) {
	fmt.Println("Medidas:", g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimetro:", g.perimetro())
}
