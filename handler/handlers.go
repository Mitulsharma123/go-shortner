package handler

import (
	"net/http"

	shortener "go-shortner/shortner"
	"go-shortner/store"

	"github.com/gin-gonic/gin"
)

// Request model definition
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

// handler function
func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// shortner.GenerateShortLink will generate hash
	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)
	// storre the mapping of output
	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})

}

//redirection handler
func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	//getting the short url from path parameter
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	// call to store the retrival url
	c.Redirect(302, initialUrl)
	//apply httpd redirection function
}
