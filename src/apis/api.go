package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"net/http"
	"io/ioutil"
	"fmt"
)

func getAnimeById(r *gin.Engine) {
	r.GET("/anime/:id", func(c *gin.Context) {
		id := c.Param("id")
		url := fmt.Sprintf("https://api.jikan.moe/v4/anime/%s", id)

		resp, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		c.Data(http.StatusOK, "application/json", body)
	})
}

func getSeasonAnime(r *gin.Engine) {
	r.GET("/anime/now", func(c *gin.Context) {
		resp, err := http.Get("https://api.jikan.moe/v4/seasons/now")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		c.Data(http.StatusOK, "application/json", body)
	})
}

func getTopAnime(r *gin.Engine) {
	r.GET("/top/anime", func(c *gin.Context) {
		resp, err := http.Get("https://api.jikan.moe/v4/top/anime")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		c.Data(http.StatusOK, "application/json", body)
	})
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	getAnimeById(r)
	getSeasonAnime(r)
	getTopAnime(r)

	r.Run(":3000")
}