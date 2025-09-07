package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"io"
	"net/http"
)

func getSeasonAnime(r *gin.Engine) {
	r.GET("/anime/now", func(c *gin.Context) {
		resp, err := http.Get("https://api.jikan.moe/v4/seasons/now")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
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

		body, _ := io.ReadAll(resp.Body)
		c.Data(http.StatusOK, "application/json", body)
	})
}

func getAnimeByID(r *gin.Engine) {
	r.GET("/anime/:id", func(c *gin.Context) {
		id := c.Param("id")
		resp, err := http.Get("https://api.jikan.moe/v4/anime/" + id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		c.Data(http.StatusOK, "application/json", body)
	})
}

func getAnimeThemes(r *gin.Engine) {
	r.GET("/anime/:id/themes", func(c *gin.Context) {
		id := c.Param("id")
		resp, err := http.Get("https://api.jikan.moe/v4/anime/" + id + "/themes")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		c.Data(http.StatusOK, "application/json", body)
	})
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	getSeasonAnime(r)
	getTopAnime(r)
	getAnimeByID(r)
	getAnimeThemes(r)

	r.Run(":3000")
}