// main.go
package main

import (
    "airbnb-room-api/database"
    "airbnb-room-api/handlers"
    "log"

    "github.com/gin-gonic/gin"
)

func main() {
    db := database.InitDB()
    defer db.Close()

    router := gin.Default()
    roomHandler := handlers.NewRoomHandler(db)

    router.GET("/:room_id", roomHandler.GetRoomDetails)

    log.Println("Server starting on :8080")
    router.Run(":8080")
}