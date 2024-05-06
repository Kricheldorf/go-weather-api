CREATE TABLE city_bookmarks (
                       id SERIAL PRIMARY KEY,
                       city VARCHAR(255) NOT NULL,
                       state VARCHAR(255) NOT NULL,
                       country VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
);