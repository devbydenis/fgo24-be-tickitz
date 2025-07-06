INSERT INTO users (username, email, password, role) VALUES
('johndoe', 'john@email.com', 'johndoe123', 'user'),
('janedoe', 'jane@email.com', 'janedoe123', 'user'),
('adminuser', 'admin@cinemax.com', 'admin123', 'admin'),
('mikebrown', 'mike@email.com', 'mikebrown123', 'user'),
('sarahjones', 'sarah@email.com', '$sarahjones123', 'user');

SELECT * FROM users;
SELECT * FROM users WHERE email = 'denis@gmail.com'

INSERT INTO genres (name) VALUES
('Action'),
('Comedy'),
('Drama'),
('Horror'),
('Sci-Fi'),
('Romance'),
('Thriller'),
('Adventure'),
('Fantasy'),
('Animation'),
('Documentary'),
('Mystery'),
('Crime'),
('Family'),
('Musical'),
('Western'),
('War'),
('History');

SELECT id FROM genres WHERE name = 'War';