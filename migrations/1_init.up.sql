-- Таблица: category_of_difficulty
CREATE TABLE category_of_difficulty (
    id SERIAL PRIMARY KEY,
    title VARCHAR(20) UNIQUE NOT NULL
);

-- Таблица: sport_category
CREATE TABLE sport_category (
    id SERIAL PRIMARY KEY,
    title VARCHAR(10) UNIQUE NOT NULL
);

-- Таблица: position
CREATE TABLE position (
    id SERIAL PRIMARY KEY,
    title VARCHAR(30) UNIQUE NOT NULL,
    description_of VARCHAR(100)
);

-- Таблица: alpinists
CREATE TABLE alpinists (
    id SERIAL PRIMARY KEY,
    surname VARCHAR(100) NOT NULL,
    name_ VARCHAR(100) NOT NULL,
    address_ VARCHAR(255),
    phone VARCHAR(20) UNIQUE,
    sex VARCHAR(10) NOT NULL,
    id_sport_category INT,
    username VARCHAR(50),
    password_ VARCHAR(100),
    FOREIGN KEY (id_sport_category) REFERENCES sport_category(id) ON DELETE CASCADE
);

-- Таблица: equipment
CREATE TABLE equipment (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) UNIQUE NOT NULL,
    quantity_available INT NOT NULL CHECK (quantity_available >= 0),
    image_url VARCHAR(255),
    description TEXT
);

-- Таблица: mountain
CREATE TABLE mountain (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) UNIQUE,
    height INT NOT NULL,
    mountain_range VARCHAR(255)
);

-- Таблица: mountain_climbs
CREATE TABLE mountain_climbs (
    id SERIAL PRIMARY KEY,
    id_mountain INT,
    id_category INT,
    title VARCHAR(255),
    season VARCHAR(50),
    duration VARCHAR(50),
    distance VARCHAR(50),
    elevation VARCHAR(50),
    map_url VARCHAR(255),
    rating DECIMAL,
    description VARCHAR(255),
    start_date DATE,
    end_date DATE,
    total VARCHAR(10),
    places_left INT CHECK (places_left >= 0),
    photo_url VARCHAR(255),
    FOREIGN KEY (id_mountain) REFERENCES mountain(id) ON DELETE CASCADE,
    FOREIGN KEY (id_category) REFERENCES category_of_difficulty(id) ON DELETE CASCADE
);

-- Таблица: equipment_inventory
CREATE TABLE equipment_inventory (
    id SERIAL PRIMARY KEY,
    id_equipment INT,
    id_alpinist INT,
    FOREIGN KEY (id_alpinist) REFERENCES alpinists(id) ON DELETE CASCADE,
    FOREIGN KEY (id_equipment) REFERENCES equipment(id) ON DELETE CASCADE
);

-- Таблица: team
CREATE TABLE team (
    id SERIAL PRIMARY KEY,
    surname_name VARCHAR(100),
    experience TEXT,
    date_of_birth DATE,
    address_ VARCHAR(150),
    id_position INT,
    phone VARCHAR(20) UNIQUE,
    password_ VARCHAR(255),
    login_ VARCHAR(255),
    FOREIGN KEY (id_position) REFERENCES position(id) ON DELETE CASCADE
);

-- Таблица: team_leaders
CREATE TABLE team_leaders (
    id SERIAL PRIMARY KEY,
    id_mountain_climb INT,
    id_team_member INT,
    FOREIGN KEY (id_mountain_climb) REFERENCES mountain_climbs(id) ON DELETE CASCADE,
    FOREIGN KEY (id_team_member) REFERENCES team(id) ON DELETE CASCADE
);

CREATE TABLE climb_equipment (
    climb_id INT REFERENCES mountain_climbs(id) ON DELETE CASCADE,
    equipment_id INT REFERENCES equipment(id) ON DELETE CASCADE,
    PRIMARY KEY (climb_id, equipment_id)
);

CREATE TABLE climb_images (
    id SERIAL PRIMARY KEY,
    climb_id INT REFERENCES mountain_climbs(id) ON DELETE CASCADE,
    url TEXT
);

CREATE TABLE alpinist_equipment (
    id SERIAL PRIMARY KEY,
    equipment_id INT REFERENCES equipment(id) ON DELETE CASCADE,
    alpinist_id INT REFERENCES alpinists(id) ON DELETE CASCADE,
    date_of_issue DATE,
    date_of_return DATE,
    UNIQUE(alpinist_id, equipment_id)
);

CREATE TABLE alpinist_climb (
    id SERIAL PRIMARY KEY,
    alpinist_id INT REFERENCES alpinists(id) ON DELETE CASCADE,
    climb_id INT REFERENCES mountain_climbs(id) ON DELETE CASCADE,
    status VARCHAR(20) NOT NULL,
    UNIQUE(alpinist_id, climb_id)
);

