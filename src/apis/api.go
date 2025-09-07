package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"io"
	"strings"
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

func getAnimeCharacters(r *gin.Engine) {
	r.GET("/anime/:id/characters", func(c *gin.Context) {
		id := c.Param("id")
		resp, err := http.Get("https://api.jikan.moe/v4/anime/" + id + "/characters")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		c.Data(http.StatusOK, "application/json", body)
	})
}

func getAnimeBanner(r *gin.Engine) {
  r.GET("/anime/:id/banner", func(c *gin.Context) {
    id := c.Param("id")

    graphqlQuery := `{"query":"query ($idMal: Int) { Media(idMal: $idMal, type: ANIME) { bannerImage } }","variables":{"idMal":` + id + `}}`

    req, err := http.NewRequest("POST", "https://graphql.anilist.co", strings.NewReader(graphqlQuery))
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    req.Header.Set("Content-Type", "application/json")

    client := http.Client{}
    resp, err := client.Do(req)
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
	getAnimeCharacters(r)
	getAnimeBanner(r)

	r.Run(":3000")
}