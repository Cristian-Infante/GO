package main

import (
	"fmt"
)

var otra int

func main() {
	var elemento int
	var elemento2 int
	elemento2 = 20
	elemento3 := 20
	otra2 := 20
	otra = elemento3 + otra2

	fmt.Println(elemento)
	fmt.Println(elemento2)
	fmt.Println(otra)

	var b, c, d int = 3, 4, 5
	fmt.Println(b, c, d)

	var e, f = true, "hola"
	fmt.Println(e, f)

	g, h := 23.4, false
	fmt.Println(g, h)

	var (
		a int
		i int    = 20
		j string = "adios"
	)
	fmt.Println(a, i, j)

	const Pi = 3.1416
	fmt.Println(Pi)

	const (
		k int    = 1
		l int    = 20
		m string = "adios"
	)
	fmt.Println(k, l, m)
	fmt.Printf("%v %T %v\n", k, l, m)

	var booleano bool
	var flotante float32
	var cadena string
	fmt.Println(booleano, flotante, cadena)

	// Comentario de una linea
	/*
		Comentario multiples lineas
	*/
}
