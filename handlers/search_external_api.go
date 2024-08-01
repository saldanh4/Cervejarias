package handlers

import (
	"cervejarias-app/controllers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetExternalApi(c *gin.Context) {
	var breweriesList []controllers.Brewery
	url := "https://api.openbrewerydb.org/v1/breweries?by_country=" + typedBrewery.Country

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erro na chamada da API: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Erro ao ler a resposta da chamada: %v", err)
	}

	err = json.Unmarshal(body, &breweriesList)
	if err != nil {
		fmt.Printf("Erro ao fazer Unmarshal da resposta da chamada: %v", err)
	}
	c.IndentedJSON(http.StatusOK, breweriesList)
}
