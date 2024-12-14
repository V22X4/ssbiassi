// handlers/room_handler.go
package handlers

import (
	"airbnb-room-api/models"
	"database/sql"
	"encoding/json"
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

    var availableDates []byte // This will hold the raw []byte from the DB

    // Scan the query results
    err := h.DB.QueryRow(query, roomID).Scan(
        &room.RoomID, 
        &room.RatePerNight, 
        &room.MaxGuests, 
        &availableDates, // Get raw byte data
    )

    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Unmarshal the available_dates from []byte into a map[string]bool
    var availableDatesMap map[string]bool
    if err := json.Unmarshal(availableDates, &availableDatesMap); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse available_dates"})
        return
    }

    // Assign the unmarshalled map to the Room struct
    room.AvailableDates = availableDatesMap

    // Prepare the response
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