ALTER TABLE casts
DROP COLUMN character_name;

ALTER TABLE movies_casts
ADD COLUMN character_name VARCHAR(255);

