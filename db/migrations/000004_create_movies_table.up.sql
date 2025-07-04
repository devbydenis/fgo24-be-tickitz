-- CREATE TYPE movie_status AS ENUM ('now playing', 'coming soon', 'ended');

CREATE TABLE movies(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    release_date DATE NOT NULL,
    duration INT NOT NULL,
    rating DECIMAL(3, 1) CHECK (rating >= 0 AND rating <= 10),
    status movie_status NOT NULL,
    language VARCHAR(50) NOT NULL,
    backdrop_img VARCHAR(255),
    poster_img VARCHAR(255),
    popularity FLOAT NOT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);