package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pila struct {
	items []string
}

func (p *Pila) Push(item string) {
	p.items = append(p.items, item)
}

func (p *Pila) Pop() string {
	if !p.IsEmpty() {
		lastIndex := len(p.items) - 1
		item := p.items[lastIndex]
		p.items = p.items[:lastIndex]
		return item
	}
	return "La pila está vacía"
}

func (p *Pila) IsEmpty() bool {
	return len(p.items) == 0
}

func main() {
	pila := Pila{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Ingrese una acción (push <elemento>, pop, verificar, salir): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		accion := strings.Split(input, " ")
		comando := strings.ToLower(accion[0])

		switch comando {
		case "push":
			if len(accion) > 1 {
				elemento := accion[1]
				pila.Push(elemento)
				fmt.Printf("Se agregó '%s' a la pila.\n", elemento)
			} else {
				fmt.Println("Por favor, especifique el elemento a agregar.")
			}
		case "pop":
			elemento := pila.Pop()
			fmt.Printf("Se extrajo '%s' de la pila.\n", elemento)
		case "verificar":
			if pila.IsEmpty() {
				fmt.Println("La pila está vacía.")
			} else {
				fmt.Println("La pila no está vacía.")
			}
		case "salir":
			return
		default:
			fmt.Println("Comando no válido.")
		}
	}
}