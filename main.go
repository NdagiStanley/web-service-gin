package main

// A standalone program (as opposed to a library) is always in package main
import (
    "net/http"

    "github.com/gin-gonic/gin"

    "database/sql"
    "log"
    _ "github.com/mattn/go-sqlite3"
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

// Define global database variable
var db *sql.DB

type album struct {
    ID     int     `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// Sample data to seed
var seedAlbums = []album{
    {Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// 3. Logic to prepare a response
// Initialize a Gin router using Default
// Note: getAlbums NOT getAlbums() | function name not function result
// Run function to attach the router to an http.Server and start the server.
func main() {
    var err error
    db, err = sql.Open("sqlite3", "./foo.db")
    if err != nil {
        log.Fatalf("Failed to open database: %v", err)
    }
    defer db.Close()

    // Ensure the database connection is active
    err = db.Ping()
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    // Seed data into the database
    seedData()

    router := gin.Default()
    router.GET("/albums", getAlbums)
    // the colon preceding an item in the path signifies that the item is a path parameter
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)

    log.Println("Starting server on localhost:8080")
    router.Run("localhost:8080")
}

// seedData inserts sample data into the database
func seedData() {
    // Check if the table is empty before seeding
    var count int
    err := db.QueryRow("SELECT COUNT(*) FROM albums").Scan(&count)
    if err != nil {
        log.Fatalf("Failed to check albums count: %v", err)
    }

    if count == 0 {
        log.Println("Seeding initial data into the albums table...")
        stmt, err := db.Prepare("INSERT INTO albums (title, artist, price) VALUES (?, ?, ?)")
        if err != nil {
            log.Fatalf("Failed to prepare seed insert statement: %v", err)
        }
        defer stmt.Close()

        for _, a := range seedAlbums {
            _, err = stmt.Exec(a.Title, a.Artist, a.Price)
            if err != nil {
                log.Fatalf("Failed to insert album %v: %v", a, err)
            }
        }
        log.Println("Seeding completed.")
    } else {
        log.Println("Albums table already has data; skipping seeding.")
    }
}

// 4. Code to map the request path to your logic
// getAlbums responds with the list of all albums as JSON.
// gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more
// gin.Context != Go's built-in context package

// Context.IndentedJSON is called to serialize the struct into JSON and add it to the response

// Alternatively you can use Context.JSON to send more compact JSON
// In practice, the indented form is much easier to work with when debugging and the size difference is usually small.
func getAlbums(c *gin.Context) {
    // Retrieve all albums from the database
    rows, err := db.Query("SELECT id, title, artist, price FROM albums")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query albums"})
        return
    }
    defer rows.Close()

    var albums []album
    for rows.Next() {
        var a album
        if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan album data"})
            return
        }
        albums = append(albums, a)
    }

    c.IndentedJSON(http.StatusOK, albums)
}

// Go doesn’t enforce the order in which you declare functions.
// postAlbums adds an album from JSON received in the request body.
// Context.BindJSON used to bind the request body to newAlbum.
// Append the album struct initialized from the JSON to the albums slice
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Bind the received JSON to newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Prepare the SQL statement for inserting a new album.
    stmt, err := db.Prepare("INSERT INTO albums(title, artist, price) VALUES (?, ?, ?)")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare query"})
        return
    }
    defer stmt.Close()

    // Execute the statement with newAlbum's data.
    result, err := stmt.Exec(newAlbum.Title, newAlbum.Artist, newAlbum.Price)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert album"})
        return
    }

    // Retrieve the ID of the new record.
    id, err := result.LastInsertId()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve new album ID"})
        return
    }

    // Set the ID of newAlbum and return the created album.
    newAlbum.ID = int(id)
    c.JSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
// Context.Param retrieves the id path parameter from the URL.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    var a album
    err := db.QueryRow("SELECT id, title, artist, price FROM albums WHERE id = ?", id).
        Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
    if err == sql.ErrNoRows {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
        return
    } else if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to query album"})
        return
    }

    c.IndentedJSON(http.StatusOK, a)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
