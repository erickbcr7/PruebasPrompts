package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	valor string
}

var instance *singleton
var once sync.Once

func GetInstance(valorInicial string) *singleton {
	once.Do(func() {
		instance = &singleton{valor: valorInicial}
	})
	return instance
}

func (s *singleton) GetValor() string {
	return s.valor
}

func (s *singleton) SetValor(nuevoValor string) {
	s.valor = nuevoValor
}

func main() {
	instancia1 := GetInstance("Primer valor")
	fmt.Println("Valor de instancia1:", instancia1.GetValor())

	instancia2 := GetInstance("Segundo valor")
	fmt.Println("Valor de instancia2:", instancia2.GetValor())
	fmt.Println("¿Instancia1 es la misma que instancia2?", instancia1 == instancia2)
	fmt.Println("Valor de instancia1 después de obtener instancia2:", instancia1.GetValor())

	instancia2.SetValor("Valor modificado")
	fmt.Println("Valor de instancia1 después de modificar instancia2:", instancia1.GetValor())
}