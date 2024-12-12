// models/room.go
package models

import (
    "time"
)

type Room struct {
    RoomID         string            `json:"room_id"`
    RatePerNight   float64           `json:"rate_per_night"`
    MaxGuests      int               `json:"max_guests"`
    AvailableDates map[string]bool   `json:"available_dates"`
}

func (r *Room) CalculateOccupancyPercentage() map[string]float64 {
    occupancy := make(map[string]float64)
    now := time.Now()

    for i := 0; i < 5; i++ {
        month := now.AddDate(0, i, 0)
        monthKey := month.Format("2024-01")
        
        totalDays := daysInMonth(month)
        unavailableDays := 0

        for day := 1; day <= totalDays; day++ {
            dateStr := month.AddDate(0, 0, day-1).Format("2006-01-02")
            if available, exists := r.AvailableDates[dateStr]; exists && !available {
                unavailableDays++
            }
        }

        occupancy[monthKey] = float64(unavailableDays) / float64(totalDays) * 100
    }

    return occupancy
}

func (r *Room) GetNightRates(days int) map[string]float64 {
    rates := make([]float64, 0)
    now := time.Now()

    for i := 0; i < days; i++ {
        date := now.AddDate(0, 0, i)
        dateStr := date.Format("2024-01-01") //yyy-mm-dd
        if available, exists := r.AvailableDates[dateStr]; exists && available {
            rates = append(rates, r.RatePerNight)
        }
    }

    result := map[string]float64{
        "average_rate": 0,
        "highest_rate": 0,
        "lowest_rate":  r.RatePerNight,
    }

    if len(rates) > 0 {
        total := 0.0
        for _, rate := range rates {
            total += rate
            if rate > result["highest_rate"] {
                result["highest_rate"] = rate
            }
            if rate < result["lowest_rate"] {
                result["lowest_rate"] = rate
            }
        }
        result["average_rate"] = total / float64(len(rates))
    }

    return result
}

func daysInMonth(date time.Time) int {
    return time.Date(date.Year(), date.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
}