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
	fmt.Print("Escolha a quantidade de elementos que participarão da lista: ")
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

	// 5. Crie um Array bidimensional de inteiros com 3 linhas e 2 colunas. Solicite ao usuário que informe os valores de cada elemento da matriz. Em seguida, imprima a matriz resultante.

	matriz := [3][2]int{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			var valor int
			fmt.Printf("Digite o valor da posição [%d][%d]: ", i, j)
			fmt.Scan(&valor)
			matriz[i][j] = valor

			fmt.Println("Matriz resultante:")
			for i := 0; i < 3; i++ {
				for j := 0; j < 2; j++ {
					fmt.Printf("%d ", matriz[i][j])
				}
				fmt.Println()
			}
		}
	}

	// 6. Crie um Array de inteiros com 10 elementos. Em seguida, solicite ao usuário que informe um valor e verifique se esse valor está presente no Array. Imprima o resultado da busca.

	var x int

	fmt.Print("Escolha qual elemento deseja verificar se tem na lista: ")
	fmt.Scan(&x)
	fmt.Printf("Um dos elementos que será verificado é: %d.\n", x)

	lista := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Scan(&lista)

	for i := 1; i < 2; i++ {

		if x == lista[0] {
			fmt.Printf("O elemento %d está contido na lista.\n", x)
		} else if x == lista[1] {
			fmt.Printf("O elemento %d está contido na lista.\n", x)
		} else if x == lista[2] {
			fmt.Printf("O elemento %d está contido na lista.\n", x)
		} else if x == lista[3] {
			fmt.Printf("O elemento %d está contido na lista.\n", x)
		} else if x == lista[4] {
			fmt.Printf("O elemento %d está contido na lista.\n", x)
		} else if x == lista[5] {
			fmt.Printf("O elemento %d está contido na lista.\n", x)
		} else if x == lista[6] {
			fmt.Printf("O elemento %d está contido na lista.\n", x)
		} else if x == lista[7] {
			fmt.Printf("O elemento %d está contido na lista.\n", x)
		} else if x == lista[8] {
			fmt.Printf("O elemento %d está contido na lista.\n", x)
		} else if x == lista[9] {
			fmt.Printf("O elemento %d está contido na lista.\n", x)
		} else {
			fmt.Printf("O elemento %d não está contido na lista.\n", x)
		}
	}

	// Obs: Pode ser feito também dessa maneira: Ao invés de else if, colocar OU (||).

	var x int

	fmt.Print("Escolha qual elemento deseja verificar se tem na lista: ")
	fmt.Scan(&x)
	fmt.Printf("Um dos elementos que será verificado é: %d.\n", x)

	lista := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Scan(&lista)

	for i := 1; i < 2; i++ {

		if x == lista[0] || x == lista[1] || x == lista[2] || x == lista[3] || x == lista[4] || x == lista[5] || x == lista[6] || x == lista[7] || x == lista[8] || x == lista[9] {
			fmt.Printf("O elemento %d está contido na lista.\n", x)
		} else {
			fmt.Printf("O elemento %d não está contido na lista.\n", x)
		}
	}

	// 7. Crie um Slice de inteiros com o tamanho 5. Em seguida, solicite ao usuário que informe um número inteiro. Adicione esse número ao Slice apenas se ele não estiver presente. Imprima o Slice resultante.

	slice := []int{1, 2, 3, 4, 5}
	fmt.Scan(&slice)

	var x int
	fmt.Print("Usuário, digite um valor qualquer: ")
	fmt.Scan(&x)
	fmt.Printf("O valor que você escolheu foi: %d.\n", x)

	for i := 1; i < 2; i++ {

		if x != slice[0] && x != slice[1] && x != slice[2] && x != slice[3] && x != slice[4] {
			slice = append(slice, x)
			fmt.Printf("O número %d está sendo armazenado na lista, aguarde...\n", x)
		} else {
			fmt.Printf("O número %d já está contido na lista.\n", x)
		}

		fmt.Println(slice)

	}

	// 8. Crie um Slice de strings com tamanho 8 e solicite ao usuário que informe um valor. Remova todas as ocorrências desse valor no Slice e imprima o resultado.

	slice := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	fmt.Scan(&slice)

	var x string
	fmt.Print("Usuário, digite um caractere que você deseja remover da lista: ")
	fmt.Scan(&x)
	fmt.Printf("O caractere que você deseja remover da lista é: %s.\n", x)

	for i := 1; i < 2; i++ {

		if x == slice[0] {
			slice = slice[1:]
		} else if x == slice[1] {
			slice = append(slice[:1], slice[2:]...)
		} else if x == slice[2] {
			slice = append(slice[:2], slice[3:]...)
		} else if x == slice[3] {
			slice = append(slice[:3], slice[4:]...)
		} else if x == slice[4] {
			slice = slice[:len(slice)-1]
		} else {
			fmt.Println("Digite um caractere que esteja na lista, por favor.")
		}

		fmt.Println(slice)

	}

	// 9. Crie um Array de floats com 6 elementos. Solicite ao usuário que informe um número que será adicionado em todas as posições do Array. Imprima o Array resultante.

	array := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}
	fmt.Scan(&array)

	var x float64
	fmt.Print("Usuário, digite um número que substituirá todos os elementos da lista: ")
	fmt.Scan(&x)
	fmt.Printf("O número escolhido foi: %.1f.\n", x)

	array[0] = x
	array[1] = x
	array[2] = x
	array[3] = x
	array[4] = x
	array[5] = x

	fmt.Print(array)

	// 10. Crie um Slice de inteiros com tamanho 10 e imprima o valor mínimo e o valor máximo armazenados no Slice.

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Scan(&slice)

	var maior int
	var menor int

	maior = slice[0]
	for i := 1; i < 10; i++ {
		if slice[i] > maior {
			maior = slice[i]
		}
	}

	menor = slice[0]
	for i := 1; i < 10; i++ {
		if slice[i] < menor {
			menor = slice[i]
		}
	}

	fmt.Printf("Slice: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", slice[i])
	}
	fmt.Printf("\n")

	fmt.Printf("Maior valor: %d.\n", maior)
	fmt.Printf("Menor valor: %d.", menor)

	// 11. Crie um Array bidimensional de inteiros com 2 linhas e 3 colunas. Em seguida, solicite ao usuário que informe um índice de linha e outro de coluna e imprima o valor armazenado nessa posição da matriz.

	matriz := [2][3]int{}

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			var valor int
			fmt.Printf("Digite o valor da posição [%d][%d]: ", i, j)
			fmt.Scan(&valor)
			matriz[i][j] = valor

			fmt.Println("Matriz resultante:")
			for i := 0; i < 2; i++ {
				for j := 0; j < 3; j++ {
					fmt.Printf("%d ", matriz[i][j])
				}
				fmt.Println()
			}
		}
	}

	// 12. Crie um Array de inteiros com 5 elementos. Em seguida, crie um novo Slice que contenha apenas os elementos do Array que são múltiplos de 3. Imprima o novo Slice.

	array := []int{1, 2, 3, 4, 5}
	slice := []int{}

	for _, i := range array {
		if i%3 == 0 {
			slice = append(slice, i)
		}
	}

	fmt.Println(slice)

	// 13. Crie um Array de inteiros com 7 elementos. Solicite ao usuário que informe um número que será adicionado ao primeiro e ao último elemento do Array. Imprima o Array resultante.

	array := []int{1, 2, 3, 4, 5, 6, 7}
	var x int
	var y int

	fmt.Print("Usuário, digite um número que queira adicionar ao início da lista, por favor: ")
	fmt.Scan(&x)
	fmt.Print("Usuário, digite um número que queira adicionar ao final da lista, por favor: ")
	fmt.Scan(&y)

	array = append([]int{x}, array...)
	array = append(array, y)
	fmt.Print(array)

	// 14. Crie um Slice de inteiros com tamanho 8 e solicite ao usuário que informe dois índices de elementos que deverão ser trocados de posição. Imprima o Slice resultante.

	slice := make([]int, 8)

	for i := 0; i < len(slice); i++ {
		fmt.Print("Digite o valor do elemento: ")
		fmt.Scan(&slice[i])
	}

	var i, j int
	fmt.Print("Digite o elemento que deseja trocar: ")
	fmt.Scan(&i)
	fmt.Print("Digite o elemento que quer usar para fazer a troca com o de cima: ")
	fmt.Scan(&j)

	slice[i], slice[j] = slice[j], slice[i]

	fmt.Println(slice)

	// 15. Crie um Array de floats com 10 elementos. Crie um novo Slice que contenha apenas os elementos do Array que são maiores que 5. Imprima o novo Slice.

	array := make([]float64, 10)
	slice := []float64{}

	for i := 0; i < len(array); i++ {
		fmt.Print("Digite o valor do elemento: ")
		fmt.Scan(&array[i])
	}

	for _, valor := range array {
		if valor > 5 {
			slice = append(slice, valor)
		}
	}

	fmt.Println(slice)

	// 16. Crie um Array de inteiros com 10 elementos. Crie um novo Slice que contenha apenas os elementos pares do Array original. Imprima o novo Slice.

	array := make([]int, 10)
	slice := []int{}

	for i := 0; i < len(array); i++ {
		fmt.Print("Digite o valor de um elemento, por favor: ")
		fmt.Scan(&array[i])
	}

	for _, valor := range array {
		if valor%2 == 0 {
			slice = append(slice, valor)
		}
	}

	fmt.Print(slice)

	//. 17 Crie um Array de inteiros com 10 elementos. Calcule e imprima a soma dos elementos nas posições pares do Array.

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for i := 0; i < len(arr); i += 2 {
		sum += arr[i]
	}
	fmt.Println("A soma dos elementos nas posições pares do Array é:", sum)

	// 18. Escreva um programa que leia um número inteiro positivo n e gere um array com os n primeiros números primos.

	package main

	import "fmt"

	func isPrimo(n int) bool {
		if n < 2 {
		return false
	}

		for i := 2; i*i <= n; i++ {
		if n%i == 0 {
		return false
	}
	}

		return true
	}

	func main() {
		var n int

		fmt.Print("Digite um número inteiro positivo: ")
		fmt.Scan(&n)

		primos := make([]int, 0, n)

		i := 2
		for len(primos) < n {
			if isPrimo(i) {
				primos = append(primos, i)
			}
			i++
		}

		fmt.Println(primos)
	}

	// Obs: Foi necessário a abertura de outra função, logo, para verificar o código, vai ter que copiar e colar.

	// 19. Faça um programa que leia dois arrays de inteiros de tamanho n e gere um terceiro array que seja a soma dos dois primeiros.

	var t int
	fmt.Print("Digite o tamanho dos arrays: ")
	fmt.Scan(&t)

	array1 := make([]int, t)
	array2 := make([]int, t)

	fmt.Println("Digite os valores do primeiro array:")
	for i := 0; i < t; i++ {
		fmt.Printf("Digite o valor %d: ", i + 1)
		fmt.Scan(&array1[i])
	}

	fmt.Println("Digite os valores do segundo array:")
	for i := 0; i < t; i++ {
		fmt.Printf("Digite o valor %d: ", i + 1)
		fmt.Scan(&array2[i])
	}

	array3 := make([]int, t)
	for i := 0; i < t; i++ {
		array3[i] = array1[i] + array2[i]
	}

	fmt.Println("A soma dos dois arrays é:", array3)

	// 20. Crie um programa que leia um array de inteiros e verifique se ele está ordenado em ordem crescente.

		var t int
		fmt.Print("Informe o tamanho do array: ")
		fmt.Scan(&t)

		arr := make([]int, t)
		fmt.Printf("Digite %d números inteiros separados por espaço: ", t)
		for i := 0; i < t; i++ {
			fmt.Scan(&arr[i])
		}

		crescente := true
		for i := 1; i < t; i++ {
			if arr[i] < arr[i - 1] {
				crescente = false
				break
			}
		}

		if crescente {
			fmt.Println("O array está em ordem crescente.")
		} else {
			fmt.Println("O array não está em ordem crescente.")
		}

		// Professor, está entregue, mas, sendo sincero com o senhor, fiquei com dúvida em vários, alguns eu pedi no ChatGPT e pedi para ele me explicar, outros exercícios eu consegui desenrolar e deram certo. Mas fui aprendendo, utilizando e pesquisando. Inclusive acho que esse é o intúito dos exercícios.

}
