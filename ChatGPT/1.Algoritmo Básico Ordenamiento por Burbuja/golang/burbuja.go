package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func BubbleSort(arr []int) []int {
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
    fmt.Println("Ingrese números separados por espacios:")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    partes := strings.Fields(scanner.Text())
    arr := make([]int, len(partes))
    for i, p := range partes {
        arr[i], _ = strconv.Atoi(p)
    }
    fmt.Println("Lista ordenada:", BubbleSort(arr))
}
