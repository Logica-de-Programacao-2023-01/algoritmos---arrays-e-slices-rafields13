package main

import "fmt"

func main() {

	// 1. Crie um Array de inteiros com 3 elementos e imprima a soma dos valores armazenados no Array.

	lista := [3]int{1, 2, 3}

	var soma int = lista[0] + lista[1] + lista[2]
	fmt.Printf("A soma dos elementos da lista é: %d.", soma)

	// 2. Crie um Slice de inteiros com os valores 1, 2, 3, 4 e 5. Remova o terceiro elemento e imprima o Slice resultante.

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	numbers = append(numbers[:2], numbers[3:]...)
	fmt.Println(numbers)

	// 3. Crie um Array de floats com 4 elementos e calcule o produto dos valores armazenados no Array.

	list := []float64{1, 2, 3, 4}

	var mult float64 = list[0] * list[1] * list[2] * list[3]
	fmt.Printf("A multiplicação dos valores da lista é: %.2f.", mult)

	// 4. Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice e os valores dos elementos. Em seguida, imprima o Slice e a soma dos valores armazenados.

	var x int
	fmt.Print("Escolha a qauntidade de elementos que participarão da lista: ")
	fmt.Scan(&x)
	fmt.Printf("A quantidade de elementos que participarão da lista é: %d. \n", x)

	var y int

	var slice []int

	for i := 0; i < x; i++ {

		fmt.Print("Escolha quais elementos participarão a lista: ")
		fmt.Scan(&y)
		fmt.Printf("Um dos elementos que participará da lista é: %d.\n", y)

		slice = append(slice, y)

	}

	fmt.Println(slice)
	
}
