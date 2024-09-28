package main

import (
	"github.com/gin-gonic/gin"
	"music-library/pkg/api"
	"music-library/pkg/store"
)

func main() {
	store.InitDB() // Ma'lumotlar bazasi bilan ulanishni yoqish
	router := gin.Default()

	router.GET("/songs", api.GetAllSongs)
	router.POST("/songs", api.CreateSong)
	router.GET("/songs/:id/lyrics", api.GetSongsLyrics)
	router.GET("/songs/:id", api.Getsong)
	router.PUT("/songs/:id", api.UpdateSong)
	router.DELETE("/songs/:id", api.DeleteSong)

	router.Run(":8080")
}
