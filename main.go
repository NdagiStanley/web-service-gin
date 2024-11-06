package main
// A standalone program (as opposed to a library) is always in package main
import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// Steps
// 1. Design API endpoints
// /albums - GET, POST | /albums/:id - GET
// 2. Data
// 3. Logic to prepare a response
// 4. Code to map the request path to your logic

// 2. Data
// struct declaration
// album represents data about a record album.
// Struct tags (json:"artist") specify what a field’s name should be when the struct’s contents are serialized into JSON
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

// 3. Logic to prepare a response
// Initialize a Gin router using Default
// Note: getAlbums NOT getAlbums() | function name not function result
// Run function to attach the router to an http.Server and start the server.
func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)

    router.Run("localhost:8080")
}

// 4. Code to map the request path to your logic
// getAlbums responds with the list of all albums as JSON.
// gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more
// gin.Context != Go's built-in context package

// Context.IndentedJSON is called to serialize the struct into JSON and add it to the response

// Alternatively you can use Context.JSON to send more compact JSON
// In practice, the indented form is much easier to work with when debugging and the size difference is usually small.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}
