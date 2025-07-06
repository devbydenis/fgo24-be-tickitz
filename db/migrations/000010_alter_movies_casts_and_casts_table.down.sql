ALTER TABLE movies_casts
DROP COLUMN character_name;

ALTER TABLE casts
ADD COLUMN character_name VARCHAR(255);