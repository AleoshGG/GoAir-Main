DO $$ BEGIN
    CREATE TYPE sensor_type_enum AS ENUM ('air_quality', 'temperature', 'humidity');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

DO $$ BEGIN
    CREATE TYPE status_application_enum AS ENUM ('requested', 'pending', 'complete');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE TABLE IF NOT EXISTS users (
    id_user     SERIAL PRIMARY KEY,
    first_name  VARCHAR(50) NOT NULL,
    last_name   VARCHAR(50) NOT NULL,
    email       VARCHAR(50) NOT NULL UNIQUE,
    password    VARCHAR(200) NOT NULL
);

CREATE TABLE IF NOT EXISTS places (
    id_place    SERIAL PRIMARY KEY,
    id_user     INT,
    name        VARCHAR(45) NOT NULL,
    create_at   TIMESTAMP DEFAULT now(), 
    FOREIGN KEY (id_user) REFERENCES users(id_user) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS sensors (
    id_sensor           VARCHAR(100),
    id_place            INT,
    sensor_type         sensor_type_enum NOT NULL,
    model               VARCHAR(10) NOT NULL,
    installation_date   TIMESTAMP DEFAULT now(),
    PRIMARY KEY (id_sensor, sensor_type),
    FOREIGN KEY (id_place) REFERENCES places(id_place) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS sensors_readings (
    id_sensor   VARCHAR(100),
    sensor_type sensor_type_enum NOT NULL,
    create_at   TIMESTAMP DEFAULT now(),
    value       DECIMAL(10,4) NOT NULL,
    PRIMARY KEY (id_sensor, create_at),
    FOREIGN KEY (id_sensor, sensor_type) REFERENCES sensors(id_sensor, sensor_type) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS devices (
    id_device   VARCHAR(100),
    id_place    INT NOT NULL,
    PRIMARY KEY (id_device),
    FOREIGN KEY (id_place) REFERENCES places(id_place) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS applications (
    id_application SERIAL PRIMARY KEY,
    status_application status_application_enum NOT NULL,
    id_user     INT,
    FOREIGN KEY (id_user) REFERENCES users(id_user) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS admin (
    password VARCHAR(200) NOT NULL,
    email    VARCHAR(50),
    PRIMARY KEY (email)
);

