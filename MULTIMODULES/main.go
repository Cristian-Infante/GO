package main

import (
	"fmt"
	"github.com/Cristian-Infante/MultiModules/Resta"
	"github.com/Cristian-Infante/MultiModules/Suma"
	"strconv"
	"strings"
)

/*
func partialSum(x int) func(int) string {
	fmt.Println(x)
	fmt.Println(x * x)
	x = x * 10
	return func(y int)
}
*/

// EsPalindromo - Verificar si una palabra es palindromo
func EsPalindromo(palabra string) bool {
	palabra = strings.ToLower(palabra)

	for i := 0; i < len(palabra)/2; i++ {
		if palabra[i] != palabra[len(palabra)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(strconv.Itoa(Suma.SumarEsto(20, 30)))
	fmt.Println(strconv.Itoa(Suma.SumarTres(10, 20, 30)))
	fmt.Println(strconv.Itoa(Resta.RestarEsto(20, 10)))

	/*
		today := time.Now()
		switch today.Day() {
		case 5:
			fmt.Println("Today is 5th. Clean your house")
			fallthrough // Hace que se ejecute el siguiente case
		case 10:
			fmt.Println("Today is 10th. Buy some wine")
			fallthrough
		case 14:
			fmt.Println("Hi")
			fallthrough
		case 15:
			fmt.Println("Today is 15th. Visit a Doctor")
		case 20, 26, 27:
			fmt.Println("Today is 20th. Buy some food")
			fallthrough
		case 31:
			fmt.Println("Party Tonight!!!!!!!!!")
			fallthrough
		default:
			fmt.Println("No information available for that day")
		}

	*/

	/*
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}

		k := 1
		for ; k <= 10; k++ {
			fmt.Println(k)
		}

		k = 1
		for k <= 10 {
			fmt.Println(k)
			k++
		}

		for k := 1; ; k++ {
			fmt.Println(k)
			if k == 10 {
				break
			}
		}

		k = 1
		for {
			fmt.Println(k)
			if k == 10 {
				break
			}
			k++
		}

		arr1 := [4]string{"hola", "adios", "bye", "hi"}
		for i, ele := range arr1 {
			fmt.Println("i: ", i, " ele: ", ele)
		}

		strDict := map[string]string{"Japon": "Tokyo", "China": "Beijing", "Canada": "Otawa"}
		for index, element := range strDict {
			fmt.Println("Index: ", index, " Element: ", element)
		}
	*/

	fmt.Println(EsPalindromo("ala"))
}
