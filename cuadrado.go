package figuras

type Cuadrado struct {
	Lado float64
}

func (cu *Cuadrado) area() float64 {
	return cu.Lado * cu.Lado
}

func (cu *Cuadrado) perimetro() float64 {
	return 4 * cu.Lado
}
