package api

import (
	"log"
	"music-library/pkg/model"
	"music-library/pkg/store"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllSongs(c *gin.Context) {
	var songs []model.Song
	db := store.GetDB()
	db.Find(&songs)
	if err := db.First(&songs).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "database is clear"})
		return
	}
	c.JSON(http.StatusOK, songs)
}

// CreateSong - yangi qo'shiq qo'shish uchun handler
func CreateSong(c *gin.Context) {
	var song model.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := store.GetDB()
	result := db.Create(&song)
	if result.Error != nil {
		log.Printf("Error adding song: %s", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ma'lumot qo'shishda xato"})
		return
	}

	c.JSON(http.StatusCreated, song)
}

func GetSongsLyrics(c *gin.Context) {
	id := c.Param("id")
	var song model.Song

	db := store.GetDB()
	if err := db.First(&song, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Music not found"})
		return
	}

	lyrics := strings.Split(song.Text, "\n\n")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "3"))
	total := len(lyrics)
	start := (page - 1) * pageSize

	if start > total {
		start = total
	}

	end := start + pageSize
	if end > total {
		end = total
	}

	c.JSON(http.StatusOK, gin.H{
		"total":    total,
		"page":     page,
		"pegeSize": pageSize,
		"lyrics":   lyrics[start:end],
	})

}

func Getsong(c *gin.Context) {
	id := c.Param("id")
	var song model.Song
	db := store.GetDB()

	if err := db.First(&song, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}
	c.JSON(http.StatusOK, song)

}

// put
func UpdateSong(c *gin.Context) {
	id := c.Param("id")
	var song model.Song
	var input model.Song

	db := store.GetDB()

	if err := db.First(&song, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Music not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&song).Updates(input)
	c.JSON(http.StatusOK, song)
}

//Delete

func DeleteSong(c *gin.Context) {
	id := c.Param("id")

	db := store.GetDB()

	if err := db.Delete(&model.Song{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Qo'shiqni uchirishda hato"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Qo'shiq uchirildi"})
}

func AllDeleteSongs(c *gin.Context) {
	db := store.GetDB()
	var song model.Song

	if err := db.First(&song).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "database is clear"})
		return
	}

	db.Where("1=1").Delete(&model.Song{})

	c.JSON(http.StatusOK, gin.H{"message": "All song deleted"})
}
