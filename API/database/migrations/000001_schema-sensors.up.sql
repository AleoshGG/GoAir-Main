DO $$ BEGIN
    CREATE TYPE sensor_type_enum AS ENUM ('air_quality', 'temperature', 'humidity');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;


CREATE TABLE IF NOT EXISTS places (
    id_place    SERIAL PRIMARY KEY,
    name        VARCHAR(45) NOT NULL,
    create_at   TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS sensors (
    id_sensor           SERIAL PRIMARY KEY,
    id_place            INT,
    sensor_type         sensor_type_enum NOT NULL,
    model               VARCHAR(50) NOT NULL,
    installation_date   DATE,
    FOREIGN KEY (id_place) REFERENCES places(id_place) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS sensors_readings (
    id_sensor   INT,
    create_at   TIMESTAMP DEFAULT now(),
    value       DECIMAL(10,4) NOT NULL,
    PRIMARY KEY (id_sensor, create_at),
    FOREIGN KEY (id_sensor) REFERENCES sensors(id_sensor) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS users (
    id_user     SERIAL PRIMARY KEY,
    first_name  VARCHAR(50) NOT NULL,
    last_name   VARCHAR(50) NOT NULL,
    email       VARCHAR(50) NOT NULL UNIQUE,
    password    VARCHAR(200) NOT NULL
);

CREATE TABLE IF NOT EXISTS users_places (
    id_place INT,
    id_user INT,
    PRIMARY KEY (id_place, id_user),
    FOREIGN KEY (id_place) REFERENCES places(id_place) ON DELETE CASCADE,
    FOREIGN KEY (id_user) REFERENCES users(id_user) ON DELETE CASCADE
);


