CREATE TABLE IF NOT EXISTS animal (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    province VARCHAR(3),
);

CREATE TABLE IF NOT EXISTS sightings (
    id SERIAL PRIMARY KEY,
    animal_id INT NOT NULL,
    latitude NUMERIC(9, 6) NOT NULL,
    longitude NUMERIC(9, 6) NOT NULL,
    sighting_time TIMESTAMP NOT NULL,
    FOREIGN KEY (animal_id) REFERENCES animals (id) ON DELETE CASCADE
);