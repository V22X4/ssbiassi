-- scripts/init.sql
CREATE TABLE rooms (
    room_id VARCHAR(50) PRIMARY KEY,
    rate_per_night DECIMAL(10, 2),
    max_guests INT,
    available_dates JSONB
);

-- Sample data
INSERT INTO rooms (room_id, rate_per_night, max_guests, available_dates) VALUES 
('ROOM001', 150.00, 4, 
    '{"2024-01-01": true, "2024-01-02": true, "2024-01-03": false, "2024-02-15": true, "2024-03-20": true}'
),
('ROOM002', 200.00, 6, 
    '{"2024-`12-05": true, "2025-01-06": true, "2025-02-10": false, "2025-03-15": true, "2025-04-22": true}'
);