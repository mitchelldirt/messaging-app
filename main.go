package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ/examples/integration-gin/gintemplrenderer"
	"github.com/gin-gonic/gin"
	"github.com/mitchelldirt/messaging-app/pages"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var layout = pages.Layout

func main() {
	router := gin.Default()
	
	ginHtmlRenderer := router.HTMLRender
	router.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}


	// Disable trusted proxy warning.
	router.SetTrustedProxies(nil)

	// load layout


	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Home", layout(pages.Home("Mitchell")))
	})

	router.GET("/blogs", func(c *gin.Context) {
	
		
		var typeOfBlog = c.Query("type")

		//output typeOfBlog to console
		fmt.Println(typeOfBlog)	

		if (typeOfBlog != "activism" && typeOfBlog != "programming") {
			c.HTML(http.StatusOK, "Blog", layout(pages.NotFound()))	
		} else {
			c.HTML(http.StatusOK, "Blog", layout(pages.Blogs(typeOfBlog)))
		}

		
	})

	router.GET("/projects", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Projects", layout(pages.Projects()))
	})

	router.GET("/skills", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Skills", layout(pages.Skills()))
	})

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}