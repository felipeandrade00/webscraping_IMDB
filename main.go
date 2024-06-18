package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Instanciando o coletor, a partir do módulo Colly
	c := colly.NewCollector()

	// Lista pra receber os dados
	var filmes []string

	// Encontrar o que queremos extrair
	c.OnHTML(".ipc-metadata-list-summary-item__c", func(e *colly.HTMLElement) {
		filme := e.ChildText(".ipc-title__text")
		filmes = append(filmes, filme)
	})

	// Aqui vamos inserir o link da página que queremos extrair os dados
	c.Visit("https://www.imdb.com/chart/top/?ref_=nv_mv_250")

	// Imprime os nomes dos times coletados
	for _, filme := range filmes {
		fmt.Println("Filme:", filme)
	}
}
