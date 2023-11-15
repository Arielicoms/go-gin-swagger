package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// gin-swagger middleware

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

type generalError struct {
	message string
	status  int
}

var albums = []album{
	{ID: "1", Title: "Familia", Artist: "Camilla Cabello", Year: 2020},
	{ID: "4", Title: "Uptown Funk", Artist: "Mark Ronson ft. Bruno Mars", Year: 2014},
	{ID: "3", Title: "Shape of You", Artist: "Ed Sheeran", Year: 2017},
	{ID: "2", Title: "Blinding Lights", Artist: "The Weeknd", Year: 2019},
}

var customError = generalError{
	message: "Aqui no joven",
	status:  400,
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusOK, nil)
	} else {
		albums = append(albums, newAlbum)
		c.IndentedJSON(http.StatusCreated, albums)
	}

}

func getAlbumID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Aqui no joven en album"})
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/album", postAlbum)
	router.GET("album/:id", getAlbumID)

	router.GET("/docs", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run("localhost:8080")
}
