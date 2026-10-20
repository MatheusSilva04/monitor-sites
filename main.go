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

	agenda := indexar(sites)
	fmt.Println(len(agenda))

	buscas := []string{"www.matheus.teste2.go", "www.naoexiste.go"}

	for _, busca := range buscas {
		site, ok := agenda[busca]
		if ok {
			fmt.Println("Site encontrado:", descrever(site))
		} else {
			fmt.Println("Site nao encontrado:", busca)
		}
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

func indexar(sites []Site) map[string]Site {
	agenda := make(map[string]Site)

	for _, site := range sites {
		err := validar(site)
		if err != nil {
			continue
		}
		agenda[site.Url] = site
	}
	return agenda
}
