package main

import (
	"fmt"
	"math"
)

// Funciones para Cuadrado
func areaCuadrado(lado float64) float64 {
	return lado * lado
}
func perimetroCuadrado(lado float64) float64 {
	return 4 * lado
}

// Funciones para Rectángulo
func areaRectangulo(base, altura float64) float64 {
	return base * altura
}
func perimetroRectangulo(base, altura float64) float64 {
	return 2 * (base + altura)
}

// Funciones para Triángulo
func areaTriangulo(base, altura float64) float64 {
	return (base * altura) / 2
}
func perimetroTriangulo(lado1, lado2, lado3 float64) float64 {
	return lado1 + lado2 + lado3
}

// Funciones para Rombo
func areaRombo(diagonalMayor, diagonalMenor float64) float64 {
	return (diagonalMayor * diagonalMenor) / 2
}
func perimetroRombo(lado float64) float64 {
	return 4 * lado
}

// Funciones para Círculo
func areaCirculo(radio float64) float64 {
	return math.Pi * radio * radio
}
func perimetroCirculo(radio float64) float64 {
	return 2 * math.Pi * radio
}

// Funciones para Pentágono
func areaPentagono(lado float64) float64 {
	apotema := lado / (2 * math.Tan(math.Pi/5))
	perimetro := 5 * lado
	return (perimetro * apotema) / 2
}
func perimetroPentagono(lado float64) float64 {
	return 5 * lado
}

// Funciones para Hexágono
func areaHexagono(lado float64) float64 {
	apotema := (lado * math.Sqrt(3)) / 2
	perimetro := 6 * lado
	return (perimetro * apotema) / 2
}
func perimetroHexagono(lado float64) float64 {
	return 6 * lado
}

// Funciones para Trapecio
func areaTrapecio(baseMayor, baseMenor, altura float64) float64 {
	return ((baseMayor + baseMenor) * altura) / 2
}
func perimetroTrapecio(lado1, lado2, baseMayor, baseMenor float64) float64 {
	return lado1 + lado2 + baseMayor + baseMenor
}

// Funciones para Paralelogramo
func areaParalelogramo(base, altura float64) float64 {
	return base * altura
}
func perimetroParalelogramo(base, lado float64) float64 {
	return 2 * (base + lado)
}

func main() {
	// Valores por defecto
	ladoCuadrado := 5.0
	baseRectangulo, alturaRectangulo := 8.0, 6.0
	baseTriangulo, alturaTriangulo, lado1Triangulo, lado2Triangulo, lado3Triangulo := 6.0, 3.0, 5.0, 5.0, 6.0
	diagonalMayorRombo, diagonalMenorRombo, ladoRombo := 10.0, 8.0, 7.0
	radioCirculo := 4.0
	ladoPentagono := 6.0
	ladoHexagono := 5.0
	baseMayorTrapecio, baseMenorTrapecio, alturaTrapecio, lado1Trapecio, lado2Trapecio := 10.0, 6.0, 4.0, 5.0, 7.0
	baseParalelogramo, alturaParalelogramo, ladoParalelogramo := 8.0, 4.0, 6.0

	// Pruebas
	fmt.Printf("Cuadrado - Área: %.2f, Perímetro: %.2f\n", areaCuadrado(ladoCuadrado), perimetroCuadrado(ladoCuadrado))
	fmt.Printf("Rectángulo - Área: %.2f, Perímetro: %.2f\n", areaRectangulo(baseRectangulo, alturaRectangulo), perimetroRectangulo(baseRectangulo, alturaRectangulo))
	fmt.Printf("Triángulo - Área: %.2f, Perímetro: %.2f\n", areaTriangulo(baseTriangulo, alturaTriangulo), perimetroTriangulo(lado1Triangulo, lado2Triangulo, lado3Triangulo))
	fmt.Printf("Rombo - Área: %.2f, Perímetro: %.2f\n", areaRombo(diagonalMayorRombo, diagonalMenorRombo), perimetroRombo(ladoRombo))
	fmt.Printf("Círculo - Área: %.2f, Perímetro: %.2f\n", areaCirculo(radioCirculo), perimetroCirculo(radioCirculo))
	fmt.Printf("Pentágono - Área: %.2f, Perímetro: %.2f\n", areaPentagono(ladoPentagono), perimetroPentagono(ladoPentagono))
	fmt.Printf("Hexágono - Área: %.2f, Perímetro: %.2f\n", areaHexagono(ladoHexagono), perimetroHexagono(ladoHexagono))
	fmt.Printf("Trapecio - Área: %.2f, Perímetro: %.2f\n", areaTrapecio(baseMayorTrapecio, baseMenorTrapecio, alturaTrapecio), perimetroTrapecio(lado1Trapecio, lado2Trapecio, baseMayorTrapecio, baseMenorTrapecio))
	fmt.Printf("Paralelogramo - Área: %.2f, Perímetro: %.2f\n", areaParalelogramo(baseParalelogramo, alturaParalelogramo), perimetroParalelogramo(baseParalelogramo, ladoParalelogramo))
}
