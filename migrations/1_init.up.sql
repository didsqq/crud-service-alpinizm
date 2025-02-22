CREATE DATABASE alpinizm;
GO

USE alpinizm;
GO

CREATE TABLE category_of_difficulty
(
 ID_category INT,
 Title NVARCHAR(20) UNIQUE NOT NULL,
 CONSTRAINT PK_category_of_difficulty PRIMARY KEY(ID_category)
)

CREATE TABLE sport_category
(
 ID_sport_category INT,
 Title NVARCHAR(10) UNIQUE NOT NULL,
 CONSTRAINT PK_sport_category PRIMARY KEY(ID_sport_category)
)

CREATE TABLE position
(
ID_position INT,
Title NVARCHAR(30) UNIQUE NOT NULL,
Description_of NVARCHAR(40),
CONSTRAINT PK_position PRIMARY KEY(ID_position)
)

CREATE TABLE alpinists
(
 ID_alpinist INT,
 Surname NVARCHAR(20) NOT NULL,
 Name_ NVARCHAR(10) NOT NULL,
 Address_ NVARCHAR(150),
 Phone VARCHAR(15) UNIQUE,
 Sex NVARCHAR(1) NOT NULL,
 ID_sport_category INT,
 CONSTRAINT PK_alpinists PRIMARY KEY(ID_alpinist),
 CONSTRAINT FK_alipinists_category FOREIGN KEY(ID_sport_category) REFERENCES sport_category(ID_sport_category) ON DELETE CASCADE
)

CREATE TABLE equipment
(
 ID_equipment INT,
 Title NVARCHAR(20) UNIQUE NOT NULL,
 Quantity_available INT NOT NULL CHECK(Quantity_available > 0),
 CONSTRAINT PK_equipment PRIMARY KEY(ID_equipment)
)

CREATE TABLE mountain
(
 ID_mountain INT ,
 Title NVARCHAR(40) UNIQUE,
 Height INT NOT NULL,
 Mountain_range NVARCHAR(60)
 CONSTRAINT PK_mountain1 PRIMARY KEY(ID_mountain)
)

CREATE TABLE groups
(
 ID_groups INT,
 Number_of_participants INT CHECK(Number_of_participants > 1),
 CONSTRAINT PK_groups PRIMARY KEY(ID_groups)
)

CREATE TABLE mountain_climbs
(
 ID_mountain_climbs INT,
 ID_groups INT,
 ID_mountain INT ,
 ID_category INT,
 Start_date_ DATE,
 End_date_ DATE,
 Total NVARCHAR(10),
 CONSTRAINT PK_mountain_climbs PRIMARY KEY(ID_mountain_climbs),
 CONSTRAINT FK_climbs_groups FOREIGN KEY(ID_groups) REFERENCES groups(ID_groups) ON DELETE CASCADE,
 CONSTRAINT FK_climbs_mountain FOREIGN KEY(ID_mountain) REFERENCES mountain(ID_mountain) ON DELETE CASCADE,
 CONSTRAINT FK_climbs_category FOREIGN KEY(ID_category) REFERENCES category_of_difficulty(ID_category) ON DELETE CASCADE
)

CREATE TABLE equipment_inventory
(
ID_entry INT,
ID_groups  INT,
ID_equipment INT,
Quantity_taken INT,
CONSTRAINT PK_equipment_inventory PRIMARY KEY(ID_entry),
CONSTRAINT FK_equipment_group FOREIGN KEY(ID_groups) REFERENCES groups(ID_groups) ON DELETE CASCADE,
CONSTRAINT FK_equimp_invent FOREIGN KEY(ID_equipment) REFERENCES equipment(ID_equipment) ON DELETE CASCADE
)

CREATE TABLE team 
(
ID_team_member INT,
Surname_name NVARCHAR(100),
Date_of_birth DATE,
Address_ NVARCHAR(150),
ID_position INT,
Phone INT UNIQUE,
Password_ VARCHAR(255),
Login_ VARCHAR(255),
CONSTRAINT PK_team PRIMARY KEY(ID_team_member),
CONSTRAINT FK_position_team FOREIGN KEY(ID_position) REFERENCES position(ID_position) ON DELETE CASCADE
)


CREATE TABLE team_leaders
(
ID_entry INT,
ID_groups INT,
ID_team_member INT,
CONSTRAINT FK_leaders_groups FOREIGN KEY(ID_groups) REFERENCES groups(ID_groups) ON DELETE CASCADE,
CONSTRAINT FK_leaders_member_group FOREIGN KEY(ID_team_member) REFERENCES team(ID_team_member) ON DELETE CASCADE,
CONSTRAINT PK_team_leaders PRIMARY KEY(ID_entry)
)

CREATE TABLE climbers_in_groups
(
ID_entry INT,
ID_alpinist INT ,
ID_groups INT,
CONSTRAINT FK_alpinist_group FOREIGN KEY(ID_alpinist) REFERENCES alpinists(ID_alpinist) ON DELETE CASCADE,
CONSTRAINT FK_group_alpinist FOREIGN KEY(ID_groups)  REFERENCES groups(ID_groups) ON DELETE CASCADE,
CONSTRAINT PK_climbers_in_groups PRIMARY KEY(ID_entry)
)