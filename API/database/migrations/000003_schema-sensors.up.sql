-- Modificar funciones para usar la zona horaria
CREATE OR REPLACE FUNCTION get_air_quality_avg(
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
        sensors_readings sr
    INNER JOIN sensors s 
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


CREATE OR REPLACE FUNCTION get_temperature_last24(
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
            sensors_readings sr
        INNER JOIN sensors s 
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


CREATE OR REPLACE FUNCTION get_humidity_last24(
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
            sensors_readings sr
        INNER JOIN sensors s 
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

CREATE OR REPLACE FUNCTION get_air_quality_last24(
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
            sensors_readings sr
        INNER JOIN sensors s 
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
