package main

import (
	"fmt"
	"github.com/boombuler/led"
	"image/color"
	"time"
        "net/http"
        "github.com/gin-gonic/gin"
)

// TODO: query system for blink device only once
// TODO: add color parameter to each call to blink
// TODO: make each call to blink async

var RED color.RGBA = color.RGBA{0xFF, 0x00, 0x00, 0xFF}
var GREEN color.RGBA = color.RGBA{0x00, 0xFF, 0x00, 0xFF}
var BLUE color.RGBA = color.RGBA{0x00, 0x00, 0xFF, 0xFF}

func qAlbums() {
  for devInfo := range led.Devices() {
    dev, err := devInfo.Open()
    if err != nil {
      fmt.Println(err)
      continue
    }

    defer dev.Close()

    dev.SetColor(RED)
    time.Sleep(2 * time.Second)
  }
}

func qAlbumByID() {
  for devInfo := range led.Devices() {
    dev, err := devInfo.Open()
    if err != nil {
      fmt.Println(err)
      continue
    }

    defer dev.Close()

    dev.SetColor(GREEN)
    time.Sleep(2 * time.Second)
  }
}

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

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    qAlbums()
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            qAlbumByID()
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
