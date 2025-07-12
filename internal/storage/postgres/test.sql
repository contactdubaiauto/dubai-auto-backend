-- drop table if exists vehicles;

-- CREATE TABLE vehicles (
--     Make TEXT,
--     Model TEXT,
--     Annual_Petroleum_Consumption_FT1 NUMERIC,
--     Annual_Petroleum_Consumption_FT2 NUMERIC,
--     Time_to_charge_120V NUMERIC,
--     Time_to_charge_240V NUMERIC,
--     City_MPG_FT1 INTEGER,
--     Unrounded_City_MPG_FT1 NUMERIC,
--     City_MPG_FT2 INTEGER,
--     Unrounded_City_MPG_FT2 NUMERIC,
--     City_gasoline_consumption NUMERIC,
--     City_electricity_consumption NUMERIC,
--     EPA_city_utility_factor NUMERIC,
--     CO2_FT1 NUMERIC,
--     CO2_FT2 NUMERIC,
--     CO2_Tailpipe_FT2 NUMERIC,
--     CO2_Tailpipe_FT1 NUMERIC,
--     Combined_MPG_FT1 INTEGER,
--     Unrounded_Combined_MPG_FT1 NUMERIC,
--     Combined_MPG_FT2 INTEGER,
--     Unrounded_Combined_MPG_FT2 NUMERIC,
--     Combined_electricity_consumption NUMERIC,
--     Combined_gasoline_consumption NUMERIC,
--     EPA_combined_utility_factor NUMERIC,
--     Cylinders INTEGER,
--     Engine_displacement NUMERIC,
--     Drive TEXT,
--     EPA_model_type_index INTEGER,
--     Engine_descriptor TEXT,
--     EPA_Fuel_Economy_Score INTEGER,
--     Annual_Fuel_Cost_FT1 INTEGER,
--     Annual_Fuel_Cost_FT2 INTEGER,
--     Fuel_Type TEXT,
--     Fuel_Type1 TEXT,
--     GHG_Score INTEGER,
--     GHG_Score_Alt_Fuel INTEGER,
--     Highway_MPG_FT1 INTEGER,
--     Unrounded_Highway_MPG_FT1 NUMERIC,
--     Highway_MPG_FT2 INTEGER,
--     Unrounded_Highway_MPG_FT2 NUMERIC,
--     Highway_gasoline_consumption NUMERIC,
--     Highway_electricity_consumption NUMERIC,
--     EPA_highway_utility_factor NUMERIC,
--     Hatchback_luggage_volume NUMERIC,
--     Hatchback_passenger_volume NUMERIC,
--     ID INTEGER,
--     Two_door_luggage_volume NUMERIC,
--     Four_door_luggage_volume NUMERIC,
--     MPG_Data TEXT,
--     PHEV_Blended BOOLEAN,
--     Two_door_passenger_volume NUMERIC,
--     Four_door_passenger_volume NUMERIC,
--     Range_FT1 NUMERIC,
--     Range_City_FT1 NUMERIC,
--     Range_City_FT2 NUMERIC,
--     Range_Highway_FT1 NUMERIC,
--     Range_Highway_FT2 NUMERIC,
--     Transmission TEXT,
--     Unadjusted_City_MPG_FT1 NUMERIC,
--     Unadjusted_City_MPG_FT2 NUMERIC,
--     Unadjusted_Highway_MPG_FT1 NUMERIC,
--     Unadjusted_Highway_MPG_FT2 NUMERIC,
--     Vehicle_Size_Class TEXT,
--     Year INTEGER,
--     You_Save_Spend INTEGER,
--     Guzzler TEXT,
--     Transmission_descriptor TEXT,
--     T_Charger TEXT,
--     S_Charger TEXT,
--     ATV_Type TEXT,
--     Fuel_Type2 TEXT,
--     EPA_Range_FT2 TEXT,
--     Electric_motor TEXT,
--     MFR_Code TEXT,
--     c240Dscr TEXT,
--     charge240b TEXT,
--     C240B_Dscr TEXT,
--     Created_On DATE,
--     Modified_On DATE,
--     Start_Stop BOOLEAN,
--     PHEV_City NUMERIC,
--     PHEV_Highway NUMERIC,
--     PHEV_Combined integer,
--     baseModel TEXT
-- );
-- COPY vehicles FROM '/tmp/vehicles.csv' DELIMITER ';' CSV HEADER;


-- --  new tables

-- -- 1. Engine Table
-- CREATE TABLE engine (
--     engine_id SERIAL PRIMARY KEY,
--     cylinders INTEGER,
--     displacement NUMERIC,
--     engine_descriptor TEXT
-- );

-- -- 2. Fuel Table
-- CREATE TABLE fuel (
--     fuel_id SERIAL PRIMARY KEY,
--     fuel_type TEXT,
--     fuel_type1 TEXT,
--     fuel_type2 TEXT
-- );

-- -- 3. Transmission Table
-- CREATE TABLE transmission (
--     transmission_id SERIAL PRIMARY KEY,
--     transmission TEXT,
--     transmission_descriptor TEXT,
--     t_charger TEXT,
--     s_charger TEXT
-- );

-- -- 4. Drive Table
-- CREATE TABLE drive (
--     drive_id SERIAL PRIMARY KEY,
--     drive TEXT
-- );

-- -- 5. Size Class Table
-- CREATE TABLE vehicle_size_class (
--     size_class_id SERIAL PRIMARY KEY,
--     vehicle_size_class TEXT
-- );

-- -- 6. ATV Type Table
-- CREATE TABLE atv_type (
--     atv_type_id SERIAL PRIMARY KEY,
--     atv_type TEXT
-- );

-- -- 7. Electric Motor Table
-- CREATE TABLE electric_motor (
--     electric_motor_id SERIAL PRIMARY KEY,
--     electric_motor TEXT
-- );


-- -- ENGINE
-- INSERT INTO engine (cylinders, displacement, engine_descriptor)
-- SELECT DISTINCT cylinders, engine_displacement, engine_descriptor
-- FROM vehicles
-- WHERE cylinders IS NOT NULL OR engine_displacement IS NOT NULL OR engine_descriptor IS NOT NULL;

-- -- FUEL
-- INSERT INTO fuel (fuel_type, fuel_type1, fuel_type2)
-- SELECT DISTINCT fuel_type, fuel_type1, fuel_type2
-- FROM vehicles
-- WHERE fuel_type IS NOT NULL OR fuel_type1 IS NOT NULL OR fuel_type2 IS NOT NULL;

-- -- TRANSMISSION
-- INSERT INTO transmission (transmission, transmission_descriptor, t_charger, s_charger)
-- SELECT DISTINCT transmission, transmission_descriptor, t_charger, s_charger
-- FROM vehicles
-- WHERE transmission IS NOT NULL;

-- -- DRIVE
-- INSERT INTO drive (drive)
-- SELECT DISTINCT drive
-- FROM vehicles
-- WHERE drive IS NOT NULL;

-- -- SIZE CLASS
-- INSERT INTO vehicle_size_class (vehicle_size_class)
-- SELECT DISTINCT vehicle_size_class
-- FROM vehicles
-- WHERE vehicle_size_class IS NOT NULL;

-- -- ATV TYPE
-- INSERT INTO atv_type (atv_type)
-- SELECT DISTINCT atv_type
-- FROM vehicles
-- WHERE atv_type IS NOT NULL;

-- -- ELECTRIC MOTOR
-- INSERT INTO electric_motor (electric_motor)
-- SELECT DISTINCT electric_motor
-- FROM vehicles
-- WHERE electric_motor IS NOT NULL;

-- ALTER TABLE vehicles ADD COLUMN engine_id INTEGER;
-- ALTER TABLE vehicles ADD COLUMN fuel_id INTEGER;
-- ALTER TABLE vehicles ADD COLUMN transmission_id INTEGER;
-- ALTER TABLE vehicles ADD COLUMN drive_id INTEGER;
-- ALTER TABLE vehicles ADD COLUMN size_class_id INTEGER;
-- ALTER TABLE vehicles ADD COLUMN atv_type_id INTEGER;
-- ALTER TABLE vehicles ADD COLUMN electric_motor_id INTEGER;

-- -- ENGINE
-- UPDATE vehicles v
-- SET engine_id = e.engine_id
-- FROM engine e
-- WHERE v.cylinders = e.cylinders
--   AND v.engine_displacement = e.displacement
--   AND v.engine_descriptor = e.engine_descriptor;

-- -- FUEL
-- UPDATE vehicles v
-- SET fuel_id = f.fuel_id
-- FROM fuel f
-- WHERE v.fuel_type = f.fuel_type
--   AND v.fuel_type1 = f.fuel_type1
--   AND v.fuel_type2 = f.fuel_type2;

-- -- TRANSMISSION
-- UPDATE vehicles v
-- SET transmission_id = t.transmission_id
-- FROM transmission t
-- WHERE v.transmission = t.transmission
--   AND v.transmission_descriptor = t.transmission_descriptor
--   AND v.t_charger = t.t_charger
--   AND v.s_charger = t.s_charger;

-- -- DRIVE
-- UPDATE vehicles v
-- SET drive_id = d.drive_id
-- FROM drive d
-- WHERE v.drive = d.drive;

-- -- SIZE CLASS
-- UPDATE vehicles v
-- SET size_class_id = s.size_class_id
-- FROM vehicle_size_class s
-- WHERE v.vehicle_size_class = s.vehicle_size_class;

-- -- ATV TYPE
-- UPDATE vehicles v
-- SET atv_type_id = a.atv_type_id
-- FROM atv_type a
-- WHERE v.atv_type = a.atv_type;

-- -- ELECTRIC MOTOR
-- UPDATE vehicles v
-- SET electric_motor_id = e.electric_motor_id
-- FROM electric_motor e
-- WHERE v.electric_motor = e.electric_motor;

-- ALTER TABLE vehicles ADD FOREIGN KEY (engine_id) REFERENCES engine(engine_id);
-- ALTER TABLE vehicles ADD FOREIGN KEY (fuel_id) REFERENCES fuel(fuel_id);
-- ALTER TABLE vehicles ADD FOREIGN KEY (transmission_id) REFERENCES transmission(transmission_id);
-- ALTER TABLE vehicles ADD FOREIGN KEY (drive_id) REFERENCES drive(drive_id);
-- ALTER TABLE vehicles ADD FOREIGN KEY (size_class_id) REFERENCES vehicle_size_class(size_class_id);
-- ALTER TABLE vehicles ADD FOREIGN KEY (atv_type_id) REFERENCES atv_type(atv_type_id);
-- ALTER TABLE vehicles ADD FOREIGN KEY (electric_motor_id) REFERENCES electric_motor(electric_motor_id);


-- ALTER TABLE vehicles
--     DROP COLUMN electric_motor,
--     DROP COLUMN atv_type,
--     DROP COLUMN vehicle_size_class,
--     DROP COLUMN drive,
--     DROP COLUMN fuel_type,
--     DROP COLUMN fuel_type1,
--     DROP COLUMN fuel_type2,
--     DROP COLUMN transmission,
--     DROP COLUMN transmission_descriptor,
--     DROP COLUMN t_charger,
--     DROP COLUMN s_charger,
--     DROP COLUMN cylinders,
--     DROP COLUMN engine_displacement,
--     DROP COLUMN engine_descriptor;


-- --  the vehicle table after script
-- CREATE TABLE vehicles (
--     Make TEXT,
--     Model TEXT,
--     Annual_Petroleum_Consumption_FT1 NUMERIC,
--     Annual_Petroleum_Consumption_FT2 NUMERIC,
--     Time_to_charge_120V NUMERIC,
--     Time_to_charge_240V NUMERIC,
--     City_MPG_FT1 INTEGER,
--     Unrounded_City_MPG_FT1 NUMERIC,
--     City_MPG_FT2 INTEGER,
--     Unrounded_City_MPG_FT2 NUMERIC,
--     City_gasoline_consumption NUMERIC,
--     City_electricity_consumption NUMERIC,
--     EPA_city_utility_factor NUMERIC,
--     CO2_FT1 NUMERIC,
--     CO2_FT2 NUMERIC,
--     CO2_Tailpipe_FT2 NUMERIC,
--     CO2_Tailpipe_FT1 NUMERIC,
--     Combined_MPG_FT1 INTEGER,
--     Unrounded_Combined_MPG_FT1 NUMERIC,
--     Combined_MPG_FT2 INTEGER,
--     Unrounded_Combined_MPG_FT2 NUMERIC,
--     Combined_electricity_consumption NUMERIC,
--     Combined_gasoline_consumption NUMERIC,
--     EPA_combined_utility_factor NUMERIC,
--     EPA_model_type_index INTEGER,
--     EPA_Fuel_Economy_Score INTEGER,
--     Annual_Fuel_Cost_FT1 INTEGER,
--     Annual_Fuel_Cost_FT2 INTEGER,
--     GHG_Score INTEGER,
--     GHG_Score_Alt_Fuel INTEGER,
--     Highway_MPG_FT1 INTEGER,
--     Unrounded_Highway_MPG_FT1 NUMERIC,
--     Highway_MPG_FT2 INTEGER,
--     Unrounded_Highway_MPG_FT2 NUMERIC,
--     Highway_gasoline_consumption NUMERIC,
--     Highway_electricity_consumption NUMERIC,
--     EPA_highway_utility_factor NUMERIC,
--     Hatchback_luggage_volume NUMERIC,
--     Hatchback_passenger_volume NUMERIC,
--     ID INTEGER,
--     Two_door_luggage_volume NUMERIC,
--     Four_door_luggage_volume NUMERIC,
--     MPG_Data TEXT,
--     PHEV_Blended BOOLEAN,
--     Two_door_passenger_volume NUMERIC,
--     Four_door_passenger_volume NUMERIC,
--     Range_FT1 NUMERIC,
--     Range_City_FT1 NUMERIC,
--     Range_City_FT2 NUMERIC,
--     Range_Highway_FT1 NUMERIC,
--     Range_Highway_FT2 NUMERIC,
--     Unadjusted_City_MPG_FT1 NUMERIC,
--     Unadjusted_City_MPG_FT2 NUMERIC,
--     Unadjusted_Highway_MPG_FT1 NUMERIC,
--     Unadjusted_Highway_MPG_FT2 NUMERIC,
--     Year INTEGER,
--     You_Save_Spend INTEGER,
--     Guzzler TEXT,
--     EPA_Range_FT2 TEXT,
--     MFR_Code TEXT,x d
--     c240Dscr TEXT,
--     charge240b TEXT,
--     C240B_Dscr TEXT,
--     Created_On DATE,
--     Modified_On DATE,
--     Start_Stop BOOLEAN,
--     PHEV_City NUMERIC,
--     PHEV_Highway NUMERIC,
--     PHEV_Combined integer,
--     baseModel TEXT
-- );


select 
    vs.id,
    bs.name as brand,
    rs.name as region,
    cs.name as city,
    ms.name as model,
    ts.name as transmission,
    es.value as engine,
    ds.name as drive,
    bts.name as body_type,
    fts.name as fuel_type,
    vs.year,
    vs.price,
    vs.mileage,
    vs.vin_code,
    vs.exchange,
    vs.credit,
    vs.new,
    vs.color,
    vs.credit_price,
    vs.status,
    vs.created_at,
    vs.updated_at,
    images
from vehicles vs
left join brands bs on vs.brand_id = bs.id
left join regions rs on vs.region_id = rs.id
left join cities cs on vs.city_id = cs.id
left join models ms on vs.model_id = ms.id
left join transmissions ts on vs.transmission_id = ts.id
left join engines es on vs.engine_id = es.id
left join drives ds on vs.drive_id = ds.id
left join body_types bts on vs.body_type_id = bts.id
left join fuel_types fts on vs.fuel_type_id = fts.id
left join lateral (
    select 
        json_agg(image) as images
    from images 
    where vehicle_id = vs.id
) images on true;


with fts as (
    select
        gft.generation_id,
        json_agg(
            json_build_object(
                'id', ft.id,
                'name', ft.name
            )
        ) as fuel_types
    from generation_fuel_types gft
    left join fuel_types ft on gft.fuel_type_id = ft.id
    group by gft.generation_id
), bts as (
    select
        gbts.generation_id,
        json_agg(
            json_build_object(
                'id', bts.id,
                'name', bts.name,
                'image', bts.image
            )
        ) as body_types
    from generation_body_types gbts 
    left join body_types bts on gbts.body_type_id = bts.id
    group by gbts.generation_id
), dts as (
    select
        gds.generation_id,
        json_agg(
            json_build_object(
                'id', ds.id,
                'name', ds.name
            )
        ) as drivetrains
    from generation_drivetrains gds
    left join drivetrains ds on gds.drivetrain_id = ds.id
    group by gds.generation_id
), ts as (
    select
        gts.generation_id,
        json_agg(
            json_build_object(
                'id', t.id,
                'name', t.name
            )
        ) as transmissions
    from generation_transmissions gts
    left join transmissions t on gts.transmission_id = t.id
    group by gts.generation_id
)

SELECT 
    gs.id, 
    gs.name, 
    gs.image, 
    gs.start_year, 
    gs.end_year,
    fts.fuel_types,
    bts.body_types,
    dts.drivetrains,
    ts.transmissions
FROM generations gs
left join fts on gs.id = fts.generation_id
left join bts on gs.id = bts.generation_id
left join dts on gs.id = dts.generation_id
left join ts on gs.id = ts.generation_id
WHERE gs.model_id = 1;







SELECT 
    gs.id, 
    gs.name, 
    gs.image, 
    gs.start_year, 
    gs.end_year,
    ARRAY_AGG(DISTINCT bt.name) AS body_types,
    ARRAY_AGG(DISTINCT t.name) AS transmissions,
    ARRAY_AGG(DISTINCT ft.name) AS fuel_types,
    ARRAY_AGG(DISTINCT d.name) AS drivetrains
FROM generations gs
LEFT JOIN generation_body_types gbt ON gs.id = gbt.generation_id
LEFT JOIN body_types bt ON gbt.body_type_id = bt.id
LEFT JOIN generation_transmissions gt ON gs.id = gt.generation_id
LEFT JOIN transmissions t ON gt.transmission_id = t.id
LEFT JOIN generation_fuel_types gft ON gs.id = gft.generation_id
LEFT JOIN fuel_types ft ON gft.fuel_type_id = ft.id
LEFT JOIN generation_drivetrains gd ON gs.id = gd.generation_id
LEFT JOIN drivetrains d ON gd.drivetrain_id = d.id
WHERE gs.model_id = 1
GROUP BY gs.id, gs.name, gs.image, gs.start_year, gs.end_year;





		select 
			vs.id,
			bs.name as brand,
			rs.name as region,
			cs.name as city,
			cls.name as color,
			icls.name as interior_color,
			ms.name as model,
			ts.name as transmission,
			es.value as engine,
			ds.name as drive,
			bts.name as body_type,
			fts.name as fuel_type,
			vs.year,
			vs.price,
			vs.mileage_km,
			vs.vin_code,
			vs.exchange,
			vs.credit,
			vs.new,
			vs.credit_price,
			vs.status,
			vs.created_at,
			vs.updated_at,
			images,
			vs.phone_numbers
		from vehicles vs
		left join colors icls on icls.id = vs.interior_color_id
		left join colors cls on vs.color_id = cls.id
		left join brands bs on vs.brand_id = bs.id
		left join regions rs on vs.region_id = rs.id
		left join cities cs on vs.city_id = cs.id
		left join models ms on vs.model_id = ms.id
		left join transmissions ts on vs.transmission_id = ts.id
		left join engines es on vs.engine_id = es.id
		left join drivetrains ds on vs.drivetrain_id = ds.id
		left join body_types bts on vs.body_type_id = bts.id
		left join fuel_types fts on vs.fuel_type_id = fts.id
		left join lateral (
			select 
				json_agg(image) as images
			from images 
			where vehicle_id = vs.id
		) images on true
		where vs.user_id = $1


        update vehicles 
			set status = 2, 
                user_id = 1
		where id = 1;


with popular as (
    SELECT 
        json_agg(
            json_build_object(
                'id', id, 
                'name', name, 
                'car_count', car_count 
            )
        ) as popular_models
    FROM models 
    models WHERE brand_id = 1 and popular = true
), all_models as (
    SELECT 
        json_agg(
            json_build_object(
                'id', id, 
                'name', name, 
                'car_count', car_count 
            )
        ) as all_models
    FROM models 
    models WHERE brand_id = 1 
)
select 
    pp.popular_models,
    ab.all_models
from popular as pp
left join all_models as ab on true;



select 
    c.id, 
    c.name,
    json_agg(
        json_build_object(
            'id', r.id,
            'name', r.name
        )
    ) as regions
from cities c
left join regions r on r.city_id = c.id
group by c.id, c.name;

select 
    bts.id,
    bts.name
from generations gs
left join body_types bts on bts.id = gs.body_type_id
where gs.start_year <= 2020 and gs.end_year >= 2020;


with gms as (
    select 
        json_agg(
            json_build_object(
                'engine_id', es.id, 
                'engine', es.value, 
                'fuel_type_id', fts.id, 
                'fuel_type', fts.name, 
                'drivetrain_id', ds.id, 
                'drivetrain', ds.name, 
                'transmission_id', ts.id, 
                'transmission', ts.name
            )
        ) as modifications,
        gms.generation_id
    from generation_modifications gms
    left join engines es on es.id = gms.engine_id
    left join fuel_types fts on fts.id = gms.fuel_type_id
    left join drivetrains ds on ds.id = gms.drivetrain_id
    left join transmissions ts on ts.id = gms.transmission_id
    where gms.generation_id = any (
        select 
            id 
        from generations 
        where 
            model_id = 1 and start_year <= 2020 and end_year >= 2020
            and body_type_id = 1
    )
    group by gms.generation_id 
)
select
    gs.id,
    gs.name,
    gs.image,
    gs.start_year,
    gs.end_year,
    gms.modifications
from gms
left join generations gs on gs.id = gms.generation_id;

