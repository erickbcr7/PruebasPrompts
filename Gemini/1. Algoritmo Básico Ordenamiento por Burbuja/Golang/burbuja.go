package main

import (
	"fmt"
	"strconv"
	"strings"
)

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	var entrada string
	fmt.Print("Ingrese una lista de números enteros separados por espacios: ")
	fmt.Scanln(&entrada)
	numerosStr := strings.Split(entrada, " ")
	numeros := make([]int, len(numerosStr))
	for i, s := range numerosStr {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Error al convertir:", err)
			return
		}
		numeros[i] = num
	}

	listaOrdenada := bubbleSort(numeros)
	fmt.Println("Lista ordenada:", listaOrdenada)
}