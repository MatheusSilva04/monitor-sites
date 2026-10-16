package main

import "fmt"

func main() {
	url := "www.matheus.teste.go"
	fmt.Println("Monitor de Sites Iniciado")
	fmt.Printf("Monitorando: %s\n", url)
	mensagem := descrever(url, 20)
	fmt.Println(mensagem)
}

func descrever(url string, segundos int) string {
	return fmt.Sprintf("%s sera checado a cada %d segundos", url, segundos)
}
