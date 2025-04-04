-- Configurar la zona horaria por defecto para la sesión (opcional pero recomendado)
ALTER DATABASE prueba SET timezone TO 'America/Mexico_City';

-- Modificar tablas para usar TIMESTAMP WITH TIME ZONE
DO $$ BEGIN
    CREATE TYPE sensor_type_enum AS ENUM ('air_quality', 'temperature', 'humidity');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

DO $$ BEGIN
    CREATE TYPE status_application_enum AS ENUM ('requested', 'pending');
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
    create_at   TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP, -- Cambiado a TIMESTAMPTZ
    FOREIGN KEY (id_user) REFERENCES users(id_user) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS sensors (
    id_sensor           VARCHAR(100),
    id_place            INT,
    sensor_type         sensor_type_enum NOT NULL,
    model               VARCHAR(10) NOT NULL,
    installation_date   TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP, -- Cambiado a TIMESTAMPTZ
    PRIMARY KEY (id_sensor, sensor_type),
    FOREIGN KEY (id_place) REFERENCES places(id_place) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS sensors_readings (
    id_sensor   VARCHAR(100),
    sensor_type sensor_type_enum NOT NULL,
    create_at   TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP, -- Cambiado a TIMESTAMPTZ
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

-- Modificar funciones para usar la zona horaria
CREATE OR REPLACE FUNCTION goair.get_air_quality_avg(
    id_place_param INT
)
RETURNS TABLE (
    fecha DATE,
    promedio_calidad_aire DECIMAL(10,4)
)
AS $$
BEGIN
    RETURN QUERY
    SELECT
        (DATE_TRUNC('day', sr.create_at AT TIME ZONE 'America/Mexico_City'))::DATE AS fecha,
        AVG(sr.value) AS promedio_calidad_aire
    FROM
        goair.sensors_readings sr
    INNER JOIN goair.sensors s 
        ON sr.id_sensor = s.id_sensor 
        AND sr.sensor_type = s.sensor_type
    WHERE
        s.id_place = id_place_param
        AND sr.sensor_type = 'air_quality'
        AND sr.create_at >= (CURRENT_TIMESTAMP AT TIME ZONE 'America/Mexico_City' - INTERVAL '3 days')::TIMESTAMPTZ
    GROUP BY
        DATE_TRUNC('day', sr.create_at AT TIME ZONE 'America/Mexico_City')
    ORDER BY
        fecha DESC;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION goair.get_temperature_last24(
    id_place_param INT
)
RETURNS TABLE (
    hora TIMESTAMP,         -- Hora local sin zona horaria
    temperatura_promedio DECIMAL(10,4)
)
AS $$
DECLARE
    start_time TIMESTAMP := date_trunc('hour', CURRENT_TIMESTAMP AT TIME ZONE 'America/Mexico_City') - INTERVAL '24 hours';
    end_time   TIMESTAMP := date_trunc('hour', CURRENT_TIMESTAMP AT TIME ZONE 'America/Mexico_City');
BEGIN
    RETURN QUERY
    SELECT
        gs.hora,
        COALESCE(temp_data.avg_temp, 0) AS temperatura_promedio
    FROM
        generate_series(start_time, end_time, '1 hour') AS gs(hora)
    LEFT JOIN (
        SELECT
            date_trunc('hour', sr.create_at AT TIME ZONE 'America/Mexico_City') AS hora,
            AVG(sr.value) AS avg_temp
        FROM
            goair.sensors_readings sr
        INNER JOIN goair.sensors s 
            ON sr.id_sensor = s.id_sensor 
           AND sr.sensor_type = s.sensor_type
        WHERE
            s.id_place = id_place_param
            AND sr.sensor_type = 'temperature'
            AND sr.create_at >= start_time
        GROUP BY
            date_trunc('hour', sr.create_at AT TIME ZONE 'America/Mexico_City')
    ) AS temp_data ON gs.hora = temp_data.hora
    ORDER BY
        gs.hora;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION goair.get_humidity_last24(
    id_place_param INT
)
RETURNS TABLE (
    hora TIMESTAMP,         -- Hora local sin zona horaria
    humedad_promedio DECIMAL(10,4)
)
AS $$
DECLARE
    start_time TIMESTAMP := date_trunc('hour', CURRENT_TIMESTAMP AT TIME ZONE 'America/Mexico_City') - INTERVAL '24 hours';
    end_time   TIMESTAMP := date_trunc('hour', CURRENT_TIMESTAMP AT TIME ZONE 'America/Mexico_City');
BEGIN
    RETURN QUERY
    SELECT
        gs.hora,
        COALESCE(temp_data.avg_temp, 0) AS humedad_promedio
    FROM
        generate_series(start_time, end_time, '1 hour') AS gs(hora)
    LEFT JOIN (
        SELECT
            date_trunc('hour', sr.create_at AT TIME ZONE 'America/Mexico_City') AS hora,
            AVG(sr.value) AS avg_temp
        FROM
            goair.sensors_readings sr
        INNER JOIN goair.sensors s 
            ON sr.id_sensor = s.id_sensor 
           AND sr.sensor_type = s.sensor_type
        WHERE
            s.id_place = id_place_param
            AND sr.sensor_type = 'humidity'
            AND sr.create_at >= start_time
        GROUP BY
            date_trunc('hour', sr.create_at AT TIME ZONE 'America/Mexico_City')
    ) AS temp_data ON gs.hora = temp_data.hora
    ORDER BY
        gs.hora;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION goair.get_air_quality_last24(
    id_place_param INT
)
RETURNS TABLE (
    hora TIMESTAMP,         -- Hora local sin zona horaria
    calidad_promedio DECIMAL(10,4)
)
AS $$
DECLARE
    start_time TIMESTAMP := date_trunc('hour', CURRENT_TIMESTAMP AT TIME ZONE 'America/Mexico_City') - INTERVAL '24 hours';
    end_time   TIMESTAMP := date_trunc('hour', CURRENT_TIMESTAMP AT TIME ZONE 'America/Mexico_City');
BEGIN
    RETURN QUERY
    SELECT
        gs.hora,
        COALESCE(temp_data.avg_temp, 0) AS humedad_promedio
    FROM
        generate_series(start_time, end_time, '1 hour') AS gs(hora)
    LEFT JOIN (
        SELECT
            date_trunc('hour', sr.create_at AT TIME ZONE 'America/Mexico_City') AS hora,
            AVG(sr.value) AS avg_temp
        FROM
            goair.sensors_readings sr
        INNER JOIN goair.sensors s 
            ON sr.id_sensor = s.id_sensor 
           AND sr.sensor_type = s.sensor_type
        WHERE
            s.id_place = id_place_param
            AND sr.sensor_type = 'air_quality'
            AND sr.create_at >= start_time
        GROUP BY
            date_trunc('hour', sr.create_at AT TIME ZONE 'America/Mexico_City')
    ) AS temp_data ON gs.hora = temp_data.hora
    ORDER BY
        gs.hora;
END;
$$ LANGUAGE plpgsql;
