package main

import "fmt"

func main() {
	fmt.Println("Monitor de Sites Iniciado")

	sites := []Site{
		{Url: "www.matheus.teste1.go", Intervalo: 10},
		{Url: "www.matheus.teste2.go", Intervalo: 20},
		{Url: "www.matheus.teste3.go", Intervalo: 30},
	}

	for _, site := range sites {
		fmt.Println(descrever(site))
	}
}

func descrever(site Site) string {
	return fmt.Sprintf("%s sera checado a cada %d segundos", site.Url, site.Intervalo)
}

type Site struct {
	Url       string
	Intervalo int
}
