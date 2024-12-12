# Airbnb Room Analytics API

## Prerequisites

- Docker
- Go 1.21+

## Installation

1. Clone the repository:
```bash
git clone https://github.com/v22x4/ssbiassi.git
cd ssbiassi
```

2. Build and Run the Application:
```bash
docker-compose up --build
```

## API Endpoint

### Get Room Details
- **URL**: `http://localhost:8080/:room_id`
- **Method**: GET
- **URL Params**: `room_id` (e.g., ROOM001, ROOM002)

### Example Response
```json
{
  "room_id": "ROOM001",
  "rate_per_night": 150.00,
  "max_guests": 4,
  "occupancy_percentage": {
    "2024-01": 33.33,
    "2024-02": 20.00,
    ...
  },
  "night_rates": {
    "average_rate": 150.00,
    "highest_rate": 150.00,
    "lowest_rate": 150.00
  }
}
```

## Technologies

- **Language**: Go
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **Containerization**: Docker
