package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"                  // Swagger файлы
	ginSwagger "github.com/swaggo/gin-swagger" // Подключение роутера Gin-Swagger
	// Для генерации документации
	_ "music-library/cmd/server/docs"
	"music-library/pkg/api"
	"music-library/pkg/store"
)

// @title Music Library API
// @version 1.0
// @description API для музыкальной библиотеки.
// @termsOfService http://swagger.io/terms/

// @contact.name Поддержка API
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /
func main() {
	store.InitDB() // Подключение к базе данных
	router := gin.Default()

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/songs", api.GetAllSongs)               // Получить все песни
	router.POST("/songs", api.CreateSong)               // Добавить новую песню
	router.GET("/songs/:id/lyrics", api.GetSongsLyrics) // Получить текст песни по id
	router.GET("/songs/:id", api.Getsong)               // Получить песню по id
	router.PUT("/songs/:id", api.UpdateSong)            // Обновить данные песни
	router.DELETE("/songs/:id", api.DeleteSong)         // Удалить песню по id
	router.DELETE("/songs/all", api.AllDeleteSongs)     // Удалить все песни

	// Запуск сервера на порту 8080
	router.Run(":8080")
}
