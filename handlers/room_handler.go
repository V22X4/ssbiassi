// handlers/room_handler.go
package handlers

import (
    "airbnb-room-api/models"
    "database/sql"
    "net/http"

    "github.com/gin-gonic/gin"
)

type RoomHandler struct {
    DB *sql.DB
}

func (h *RoomHandler) GetRoomDetails(c *gin.Context) {
    roomID := c.Param("room_id")

    var room models.Room
    query := `SELECT room_id, rate_per_night, max_guests, available_dates FROM rooms WHERE room_id = $1`
    err := h.DB.QueryRow(query, roomID).Scan(
        &room.RoomID, 
        &room.RatePerNight, 
        &room.MaxGuests, 
        &room.AvailableDates,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    response := gin.H{
        "room_id":             room.RoomID,
        "rate_per_night":      room.RatePerNight,
        "max_guests":          room.MaxGuests,
        "occupancy_percentage": room.CalculateOccupancyPercentage(),
        "night_rates":         room.GetNightRates(30),
    }

    c.JSON(http.StatusOK, response)
}

func NewRoomHandler(db *sql.DB) *RoomHandler {
    return &RoomHandler{DB: db}
}