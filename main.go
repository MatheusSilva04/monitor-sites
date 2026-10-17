package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Monitor de Sites Iniciado")

	sites := []Site{
		{Url: "www.matheus.teste1.go", Intervalo: 10},
		{Url: "www.matheus.teste2.go", Intervalo: 20},
		{Url: "www.matheus.teste3.go", Intervalo: 30},
		{Url: "", Intervalo: 0},
	}

	for _, site := range sites {
		err := validar(site)
		if err != nil {
			fmt.Println("site invalido:", err)
			continue
		}
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

func validar(site Site) error {
	if site.Url == "" {
		return errors.New("url vazio")
	}
	if site.Intervalo <= 0 {
		return fmt.Errorf("intervalo invalido: %d", site.Intervalo)
	}
	return nil
}
