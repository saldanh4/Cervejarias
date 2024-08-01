package handlers

import (
	"cervejarias-app/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

var typedBrewery controllers.SearchedInfo

// func SearchHandler(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == http.MethodPost {
// 		if err := r.ParseForm(); err != nil {
// 			fmt.Fprintf(w, "erro ao fazer parse do form: %v", err)
// 			return
// 		}
// 		typedBrewery.Country = r.PostForm.Get("country")
// 	}
// 	http.ServeFile(w, r, "static/search_form.html")
// 	//g.IndentedJSON(http.StatusOK, breweriesList)
// }

func SearchHandlerGin(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		c.String(http.StatusInternalServerError, "Erro ao fazer parse do form: %v", err)
		return
	}

	typedBrewery.Country = c.PostForm("country")
	c.HTML(http.StatusOK, "static/search_form.html", gin.H{
		"TypedBrewery": typedBrewery,
	})
}
