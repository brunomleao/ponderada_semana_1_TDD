package starter

import (
	"fmt"
	"math"
	"net/http"
)

// Esta função recebe um nome como argumento e retorna uma saudação personalizada.
func SayHello(name string) string {
	return fmt.Sprintf("Hello %v. Welcome!", name)
}

// Esta função verifica se um número é ímpar ou par.
func OddOrEven(num int) string {
	criteria := math.Mod(float64(num), 2)
	if criteria == 1 || criteria == -1 {
		return fmt.Sprintf("%v is an odd number", num)
	}
	return fmt.Sprintf("%v is an even number", num)
}

// Esta função simula uma verificação de saúde de uma página HTTP.
func Checkhealth(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "health check passed")
}
