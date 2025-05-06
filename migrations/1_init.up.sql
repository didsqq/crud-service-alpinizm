-- Таблица: category_of_difficulty
CREATE TABLE IF NOT EXISTS category_of_difficulty (
    id SERIAL PRIMARY KEY,
    title VARCHAR(20) UNIQUE NOT NULL
);

-- Таблица: sport_category
CREATE TABLE IF NOT EXISTS sport_category (
    id SERIAL PRIMARY KEY,
    title VARCHAR(10) UNIQUE NOT NULL
);

-- Таблица: position
CREATE TABLE IF NOT EXISTS position (
    id SERIAL PRIMARY KEY,
    title VARCHAR(30) UNIQUE NOT NULL,
    description_of VARCHAR(40)
);

-- Таблица: alpinists
CREATE TABLE IF NOT EXISTS alpinists (
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
CREATE TABLE IF NOT EXISTS equipment (
    id SERIAL PRIMARY KEY,
    title VARCHAR(20) UNIQUE NOT NULL,
    quantity_available INT NOT NULL CHECK (quantity_available > 0)
);

-- Таблица: mountain
CREATE TABLE IF NOT EXISTS mountain (
    id SERIAL PRIMARY KEY,
    title VARCHAR(40) UNIQUE,
    height INT NOT NULL,
    mountain_range VARCHAR(60)
);

-- Таблица: groups
CREATE TABLE IF NOT EXISTS groups (
    id SERIAL PRIMARY KEY,
    number_of_participants INT CHECK (number_of_participants > 1)
);

-- Таблица: mountain_climbs
CREATE TABLE IF NOT EXISTS mountain_climbs (
    id SERIAL PRIMARY KEY,
    id_groups INT,
    id_mountain INT,
    id_category INT,
    start_date DATE,
    end_date DATE,
    total VARCHAR(10),
    photo_url VARCHAR(255),
    FOREIGN KEY (id_groups) REFERENCES groups(id) ON DELETE CASCADE,
    FOREIGN KEY (id_mountain) REFERENCES mountain(id) ON DELETE CASCADE,
    FOREIGN KEY (id_category) REFERENCES category_of_difficulty(id) ON DELETE CASCADE
);

-- Таблица: equipment_inventory
CREATE TABLE IF NOT EXISTS equipment_inventory (
    id SERIAL PRIMARY KEY,
    id_groups INT,
    id_equipment INT,
    quantity_taken INT,
    FOREIGN KEY (id_groups) REFERENCES groups(id) ON DELETE CASCADE,
    FOREIGN KEY (id_equipment) REFERENCES equipment(id) ON DELETE CASCADE
);

-- Таблица: team
CREATE TABLE IF NOT EXISTS team (
    id SERIAL PRIMARY KEY,
    surname_name VARCHAR(100),
    date_of_birth DATE,
    address_ VARCHAR(150),
    id_position INT,
    phone VARCHAR(20) UNIQUE,
    password_ VARCHAR(255),
    login_ VARCHAR(255),
    FOREIGN KEY (id_position) REFERENCES position(id) ON DELETE CASCADE
);

-- Таблица: team_leaders
CREATE TABLE IF NOT EXISTS team_leaders (
    id SERIAL PRIMARY KEY,
    id_groups INT,
    id_team_member INT,
    FOREIGN KEY (id_groups) REFERENCES groups(id) ON DELETE CASCADE,
    FOREIGN KEY (id_team_member) REFERENCES team(id) ON DELETE CASCADE
);

-- Таблица: climbers_in_groups
CREATE TABLE IF NOT EXISTS climbers_in_groups (
    id SERIAL PRIMARY KEY,
    id_alpinist INT,
    id_groups INT,
    FOREIGN KEY (id_alpinist) REFERENCES alpinists(id) ON DELETE CASCADE,
    FOREIGN KEY (id_groups) REFERENCES groups(id) ON DELETE CASCADE
);

-- CREATE TABLE IF NOT EXISTS alpinist_equipment (

-- );