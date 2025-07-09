CREATE TABLE movies(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    release_date DATE,
    duration INT,
    rating DECIMAL(3, 1), 
    status movie_status NOT NULL,
    language VARCHAR(50),
    backdrop_img VARCHAR(255),
    poster_img VARCHAR(255),
    popularity FLOAT, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);