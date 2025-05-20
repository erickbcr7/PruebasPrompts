package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
	
)

type Stack struct{ items []int }
func (s *Stack) Push(item int) { s.items = append(s.items, item) }
func (s *Stack) Pop() (int, bool) {
    if len(s.items) == 0 { return 0, false }
    val := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return val, true
}
func (s *Stack) IsEmpty() bool { return len(s.items) == 0 }

func main() {
    s := Stack{}
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Println("\nOpciones: 1) Push 2) Pop 3) Ver pila 4) Salir")
        scanner.Scan()
        op := scanner.Text()
        switch op {
        case "1":
            fmt.Print("Valor a apilar: ")
            scanner.Scan()
            v, _ := strconv.Atoi(scanner.Text())
            s.Push(v)
            fmt.Println("Valor apilado.")
        case "2":
            val, ok := s.Pop()
            if ok { fmt.Println("Valor desapilado:", val) } else { fmt.Println("Pila vacía.") }
        case "3":
            fmt.Println("Contenido pila:", s.items)
        case "4":
            return
        }
    }
}
