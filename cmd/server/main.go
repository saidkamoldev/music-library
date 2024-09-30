package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"                  // Swagger files
	ginSwagger "github.com/swaggo/gin-swagger" // Gin-Swagger routerini ulash
	// Hujjatlar yaratish uchun
	_ "music-library/cmd/server/docs"
	"music-library/pkg/api"
	"music-library/pkg/store"
)

// @title Music Library API
// @version 1.0
// @description Qo'shiqlar kutubxonasi uchun API hujjatlari.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /
func main() {
	store.InitDB() // Ma'lumotlar bazasi bilan ulanishni yoqish
	router := gin.Default()

	// Swagger uchun endpoint qo'shish
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/songs", api.GetAllSongs)
	router.POST("/songs", api.CreateSong)
	router.GET("/songs/:id/lyrics", api.GetSongsLyrics)
	router.GET("/songs/:id", api.Getsong)
	router.PUT("/songs/:id", api.UpdateSong)
	router.DELETE("/songs/:id", api.DeleteSong)
	router.DELETE("/songs/all", api.AllDeleteSongs)
	router.Run(":8080")
}
