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


select 
    us.id,
    us.email,
    us.phone,
    ps.driving_experience,
    ps.notification,
    ps.username,
    ps.google,
    ps.birthday,
    ps.about_me
from users us
left join profiles as ps on ps.user_id = us.id
where us.id = 1;



select DISTINCT ON (bts.id)
    bts.id,
    bts.name,
    bts.image
from generation_modifications gms
left join body_types bts on bts.id = gms.body_type_id
where gms.generation_id in (
    20637504,
    21460328,
    7754683
);

select DISTINCT ON (gms.generation_id)
    gms.generation_id
from generation_modifications;



select DISTINCT ON (gs.model_id)
    gs.model_id
from generations;


delete from models where id not in (
     3232,
     3233,
     3234
);



SELECT 
    MIN(start_year) AS start_year,
    MAX(end_year) AS end_year
FROM 
    generations
WHERE 
    model_id = 3226;


SELECT 
    array_agg(y ORDER BY y) AS years
FROM (
    SELECT generate_series(start_year, end_year) AS y
    FROM generations
    WHERE model_id = 3233 AND wheel = true
) AS years_series;

ALTER TABLE profiles
ADD CONSTRAINT profiles_city_id_fk
FOREIGN KEY (city_id)
REFERENCES cities(id)

select 
    vs.id,
    json_build_object(
        'id', bs.id,
        'name', bs.name,
        'logo', bs.logo,
        'model_count', bs.model_count
    ) as brand,
    json_build_object(
        'id', rs.id,
        'name', rs.name
    ) as region,
    json_build_object(
        'id', cs.id,
        'name', cs.name
    ) as city,
    json_build_object(
        'id', ms.id,
        'name', ms.name
    ) as model,
    json_build_object(
        'id', mfs.id,
        'engine', es.value,
        'fuel_type', fts.name,
        'drivetrain', ds.name,
        'transmission', ts.name
    ) as modification,
    json_build_object(
        'id', ms.id,
        'name', ms.name,
        'image', cls.image
    ) as color,
    vs.year,
    vs.price,
    vs.odometer,
    vs.vin_code,
    vs.exchange,
    vs.credit,
    vs.new,
    vs.status,
    vs.created_at,
    images,
    vs.phone_numbers,
    vs.view_count,
    CASE
        WHEN vs.user_id = 3 THEN TRUE
        ELSE FALSE
    END AS my_car
from vehicles vs
left join colors cls on vs.color_id = cls.id
left join generation_modifications mfs on mfs.id = vs.modification_id
left join engines es on es.id = mfs.engine_id
left join transmissions ts on es.id = mfs.transmission_id
left join drivetrains ds on es.id = mfs.drivetrain_id
left join fuel_types fts on es.id = mfs.fuel_type_id
left join brands bs on vs.brand_id = bs.id
left join regions rs on vs.region_id = rs.id
left join cities cs on vs.city_id = cs.id
left join models ms on vs.model_id = ms.id
left join lateral (
    select 
        json_agg(image) as images
    from images 
    where vehicle_id = vs.id
) images on true
where vs.id = 2;



alter table profiles drop column avater;
ALTER TABLE profiles ADD COLUMN avatar varchar(200);
ALTER TABLE profiles
ADD CONSTRAINT profiles_profiles_fk
FOREIGN KEY (profiles)
REFERENCES cities(id)


DELETE FROM profiles
WHERE id NOT IN (
    SELECT MIN(id)
    FROM profiles
    GROUP BY user_id
);

ALTER TABLE profiles
ADD CONSTRAINT unique_user_id UNIQUE (user_id);


insert into images (vehicle_id, image) values (64, '/images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (63, '/images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (62, '/images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (61, '/images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');
insert into images (vehicle_id, image) values (60, '/images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');




ALTER TABLE images
ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT NOW();


ALTER TABLE videos
ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT NOW();





category                   Moto,    skuter,         motovezdehody,          snegohody
type,                      1                        1                       1
drivetrain,                1                        1                       
transmission,              1        1               1                       
cylinder_count,            1                        1                       1
cylinder_arrangement,      1                        1                       1
equipment,                 1        1               1                       1
-- ////////////////////////////////////////////////////////////////////////////////////////////
brand,                     1        1               1                       1
model,                     1        1               1                       1
engine,                    1        1               1                       1
power,                     1        1               1                       1
year,                      1        1               1                       1
fuel_type,                 1        1               1                       1
number_of_cycles,          1        1               1                       1
color,                     1        1               1                       1
odometer,                  1        1               1                       1
crash,                     1        1               1                       1
not_cleared,               1        1               1                       1
owners,                    1        1               1                       1
date_of_purchase,          1        1               1                       1
warranty_date,             1        1               1                       1
ptc,                       1        1               1                       1
vin_code,                  1        1               1                       1
certificate,               1        1               1                       1
description,               1        1               1                       1
city,                      1        1               1                       1
can_look_coordinate,       1        1               1                       1
phone_number,              1        1               1                       1
refuse_dealers_calls,      1        1               1                       1
only_chat,                 1        1               1                       1
protect_spam,              1        1               1                       1
verified_buyers,           1        1               1                       1
contact_person,            1        1               1                       1
email,                     1        1               1                       1
price,                     1        1               1                       1
price_type,                1        1               1                       1
exchange,                  1        1               1                       1



category            lyekie kommerceskie,    gruzoviki,      cedelnye tyagachi,      avtobusy,
loading_kg,             1                   1                                       
type,                                                                               1
body_type,              1                   1                                       
cabin_type,                                 1               1                       
drivetrain,             1                                                           
wheel_arrangement,                          1               1                       1
wheel_height,                                               1                       
chassis_suspension,                         1                                       
chassis_cabin,                              1               1                       
euro_exhaust,                               1                                       
seats,                  1                                                           1
equipment,              1                   1               1                       1
-- ////////////////////////////////////////////////////////////////////////////////////////////
brand,                  1                   1               1                       1
model,                  1                   1               1                       1
transmission,           1                   1               1                       1
year,                   1                   1               1                       1
fuel_type,              1                   1               1                       1
engine,                 1                   1               1                       1
power,                  1                   1               1                       1
wheel,                  1                   1               1                       1
color,                  1                   1               1                       1
odometer,               1                   1               1                       1
crash,                  1                   1               1                       1
not_cleared,            1                   1               1                       1
owners,                 1                   1               1                       1
date_of_purchase,       1                   1               1                       1
warranty_date,          1                   1               1                       1
ptc,                    1                   1               1                       1
vin_code,               1                   1               1                       1
certificate,            1                   1               1                       1
description,            1                   1               1                       1
city,                   1                   1               1                       1
can_look_coordinate,    1                   1               1                       1
phone_number,           1                   1               1                       1
refuse_dealers_calls,   1                   1               1                       1
only_chat,              1                   1               1                       1
protect_spam,           1                   1               1                       1
verified_buyers,        1                   1               1                       1
contact_person,         1                   1               1                       1
email,                  1                   1               1                       1
price,                  1                   1               1                       1
price_type,             1                   1               1                       1
exchange,               1                   1               1                       1




-- Parameters will be handled with LEFT JOIN LATERAL below
select 
    mcs.id,
    json_build_object(
        'id', pf.user_id,
        'username', pf.username,
        'avatar', pf.avatar
    ) as owner,
    mcs.engine,
    mcs.power,
    mcs.year,
    mcs.number_of_cycles,
    mcs.odometer,
    mcs.crash,
    mcs.not_cleared,
    mcs.owners,
    mcs.date_of_purchase,
    mcs.warranty_date,
    mcs.ptc,
    mcs.vin_code,
    mcs.certificate,
    mcs.description,
    mcs.can_look_coordinate,
    mcs.phone_number,
    mcs.refuse_dealers_calls,
    mcs.only_chat,
    mcs.protect_spam,
    mcs.verified_buyers,
    mcs.contact_person,
    mcs.email,
    mcs.price,
    mcs.price_type,
    mcs.status,
    mcs.updated_at,
    mcs.created_at,
    mocs.name as moto_category,
    mbs.name as moto_brand,
    mms.name as moto_model,
    fts.name as fuel_type,
    cs.name as city,
    cls.name as color,
    CASE
        WHEN mcs.user_id = 1 THEN TRUE
        ELSE FALSE
    END AS my_car,
    ps.parameters,
    images.images,
    videos.videos
from motorcycles mcs
left join profiles pf on pf.user_id = mcs.user_id
left join moto_categories mocs on mocs.id = mcs.moto_category_id
left join moto_brands mbs on mbs.id = mcs.moto_brand_id
left join moto_models mms on mms.id = mcs.moto_model_id
left join fuel_types fts on fts.id = mcs.fuel_type_id
left join cities cs on cs.id = mcs.city_id
left join colors cls on cls.id = mcs.color_id
LEFT JOIN LATERAL (
    SELECT json_agg(
        json_build_object(
            'parameter_id', mcp.moto_parameter_id,
            'parameter_value_id', mpv.id,
            'parameter', mp.name,
            'parameter_value', mpv.name
        )
    ) AS parameters
    FROM motorcycle_parameters mcp
    LEFT JOIN moto_parameters mp ON mp.id = mcp.moto_parameter_id
    LEFT JOIN moto_parameter_values mpv ON mpv.id = mcp.moto_parameter_value_id
    WHERE mcp.motorcycle_id = mcs.id
) ps ON true
LEFT JOIN LATERAL (
    SELECT json_agg(img.image) AS images
    FROM (
        SELECT image
        FROM moto_images
        WHERE moto_id = mcs.id
        ORDER BY created_at DESC
    ) img
) images ON true
LEFT JOIN LATERAL (
    SELECT json_agg(v.video) AS videos
    FROM (
        SELECT video
        FROM moto_videos
        WHERE moto_id = mcs.id
        ORDER BY created_at DESC
    ) v
) videos ON true
where mcs.id = 1;



brand_id, model_id, year, odometer, city_id, modification_id




query:


my data:
 price  
--------
  23000
  32000
  16000
  26000
  25000
  28000
  25000
  26000
 399999
(9 rows)






-- insert into brands (name, logo, car_count, popular) values ('Ford', '/images/logo/ford.png', 38, false);
-- insert into brands (name, logo, car_count, popular) values ('Chevrolet', '/images/logo/chevrolet.png', 5, false);
-- insert into brands (name, logo, car_count, popular) values ('Nissan', '/images/logo/nissan.png', 22, true);
-- insert into brands (name, logo, car_count, popular) values ('Hyundai', '/images/logo/hyundai.png', 74, false);
-- insert into brands (name, logo, car_count, popular) values ('Kia', '/images/logo/kia.png', 33, true);




-- -- ford
-- insert into models (name, brand_id, popular, car_count) values ('F-150', 3, true, 56);
-- insert into models (name, brand_id, popular, car_count) values ('Mustang', 3, true, 29);
-- insert into models (name, brand_id, popular, car_count) values ('Explorer', 3, false, 664);
-- insert into models (name, brand_id, popular, car_count) values ('Bronco', 3, false, 34);

-- -- chevrolet
-- insert into models (name, brand_id, popular, car_count) values ('Camaro', 4, true, 23);
-- insert into models (name, brand_id, popular, car_count) values ('Corvette', 4, true, 645);
-- insert into models (name, brand_id, popular, car_count) values ('Equinox', 4, false, 23);
-- insert into models (name, brand_id, popular, car_count) values ('Silverado', 4, false, 55);

-- -- nissan
-- insert into models (name, brand_id, popular, car_count) values ('Altima', 5, true, 23);
-- insert into models (name, brand_id, popular, car_count) values ('Pathfinder', 5, true, 56);
-- insert into models (name, brand_id, popular, car_count) values ('Rogue', 5, false, 22);
-- insert into models (name, brand_id, popular, car_count) values ('Sentra', 5, false, 53);

-- -- hyundai
-- insert into models (name, brand_id, popular, car_count) values ('Elantra', 6, true, 54);
-- insert into models (name, brand_id, popular, car_count) values ('Sonata', 6, false, 12);
-- insert into models (name, brand_id, popular, car_count) values ('Kona', 6, false, 45);

-- -- kia
-- insert into models (name, brand_id, popular, car_count) values ('K5', 7, true, 50);
-- insert into models (name, brand_id, popular, car_count) values ('K7', 7, false, 98);


-- insert into vehicles (
--     user_id, modification_id, brand_id, region_id, city_id, model_id, ownership_type_id, owners, view_count, year, popular, description, exchange, credit, wheel, crash, odometer, vin_code, phone_numbers, price, new, color_id, trade_in, status
-- ) values
-- (1, 1, 310, 1, 1, 3232, 1, 1, 10, 2020, 0, 'Mock car 1', false, true, true, false, 10000, 'VIN00001', ARRAY['01234567890'], 20000, false, 1, 1, 3),
-- (2, 9, 310, 2, 2, 3233, 2, 2, 20, 2019, 1, 'Mock car 2', true, false, false, true, 20000, 'VIN00002', ARRAY['0111222222'], 25000, true, 2, 2, 3),
-- (1, 13, 310, 3, 3, 3234, 1, 1, 15, 2018, 0, 'Mock car 3', false, false, true, false, 15000, 'VIN00003', ARRAY['01234567890'], 18000, false, 3, 3, 3),
-- (2, 16, 310, 4, 1, 3232, 2, 2, 5, 2021, 1, 'Mock car 4', true, true, true, false, 5000, 'VIN00004', ARRAY['0111222222'], 30000, true, 4, 1, 3),
-- (1, 19, 310, 5, 2, 3233, 1, 1, 8, 2022, 0, 'Mock car 5', false, false, false, true, 8000, 'VIN00005', ARRAY['01234567890'], 22000, false, 5, 2, 3),
-- (2, 20, 310, 6, 3, 3234, 2, 2, 12, 2023, 1, 'Mock car 6', true, true, true, false, 12000, 'VIN00006', ARRAY['0111222222'], 27000, true, 6, 3, 3),
-- (1, 25, 310, 7, 1, 3232, 1, 1, 18, 2020, 0, 'Mock car 7', false, false, false, false, 18000, 'VIN00007', ARRAY['01234567890'], 21000, false, 7, 1, 3),
-- (2, 27, 310, 8, 2, 3233, 2, 2, 22, 2019, 1, 'Mock car 8', true, true, true, true, 22000, 'VIN00008', ARRAY['0111222222'], 26000, true, 1, 2, 3),
-- (1, 35, 310, 9, 3, 3234, 1, 1, 25, 2018, 0, 'Mock car 9', false, false, true, false, 25000, 'VIN00009', ARRAY['01234567890'], 19000, false, 2, 3, 3),
-- (2, 36, 310, 10, 1, 3232, 2, 2, 30, 2021, 1, 'Mock car 10', true, false, false, false, 30000, 'VIN00010', ARRAY['0111222222'], 32000, true, 3, 1, 3),
-- (1, 37, 310, 11, 2, 3233, 1, 1, 35, 2022, 0, 'Mock car 11', false, true, true, true, 35000, 'VIN00011', ARRAY['01234567890'], 23000, false, 4, 2, 3),
-- (2, 39, 310, 12, 3, 3234, 2, 2, 40, 2023, 1, 'Mock car 12', true, false, false, false, 40000, 'VIN00012', ARRAY['0111222222'], 28000, true, 5, 3, 3),
-- (1, 41, 310, 13, 1, 3232, 1, 1, 45, 2020, 0, 'Mock car 13', false, true, true, false, 45000, 'VIN00013', ARRAY['01234567890'], 24000, false, 6, 1, 3),
-- (2, 84, 310, 1, 2, 3233, 2, 2, 50, 2019, 1, 'Mock car 14', true, false, false, true, 50000, 'VIN00014', ARRAY['0111222222'], 29000, true, 7, 2, 3),
-- (1, 89, 310, 2, 3, 3234, 1, 1, 55, 2018, 0, 'Mock car 15', false, true, true, false, 55000, 'VIN00015', ARRAY['01234567890'], 20000, false, 1, 3, 3),
-- (2, 93, 310, 3, 1, 3232, 2, 2, 60, 2021, 1, 'Mock car 16', true, false, false, false, 60000, 'VIN00016', ARRAY['0111222222'], 33000, true, 2, 1, 3),
-- (1, 96, 310, 4, 2, 3233, 1, 1, 65, 2022, 0, 'Mock car 17', false, true, true, true, 65000, 'VIN00017', ARRAY['01234567890'], 25000, false, 3, 2, 3),
-- (2, 97, 310, 5, 3, 3234, 2, 2, 70, 2023, 1, 'Mock car 18', true, false, false, false, 70000, 'VIN00018', ARRAY['0111222222'], 30000, true, 4, 3, 3),
-- (1, 98, 310, 6, 1, 3232, 1, 1, 75, 2020, 0, 'Mock car 19', false, true, true, false, 75000, 'VIN00019', ARRAY['01234567890'], 26000, false, 5, 1, 3),
-- (2, 99, 310, 7, 2, 3233, 2, 2, 80, 2019, 1, 'Mock car 20', true, false, false, true, 80000, 'VIN00020', ARRAY['0111222222'], 31000, true, 6, 2, 3),
-- (1, 107, 310, 8, 3, 3234, 1, 1, 85, 2018, 0, 'Mock car 21', false, true, true, false, 85000, 'VIN00021', ARRAY['01234567890'], 21000, false, 7, 3, 3),
-- (2, 108, 310, 9, 1, 3232, 2, 2, 90, 2021, 1, 'Mock car 22', true, false, false, false, 90000, 'VIN00022', ARRAY['0111222222'], 34000, true, 1, 1, 3),
-- (1, 109, 310, 10, 2, 3233, 1, 1, 95, 2022, 0, 'Mock car 23', false, true, true, true, 95000, 'VIN00023', ARRAY['01234567890'], 27000, false, 2, 2, 3),
-- (2, 111, 310, 11, 3, 3234, 2, 2, 100, 2023, 1, 'Mock car 24', true, false, false, false, 100000, 'VIN00024', ARRAY['0111222222'], 35000, true, 3, 3, 3),
-- (1, 113, 310, 12, 1, 3232, 1, 1, 105, 2020, 0, 'Mock car 25', false, true, true, false, 105000, 'VIN00025', ARRAY['01234567890'], 28000, false, 4, 1, 3),
-- (2, 156, 310, 13, 2, 3233, 2, 2, 110, 2019, 1, 'Mock car 26', true, false, false, true, 110000, 'VIN00026', ARRAY['0111222222'], 36000, true, 5, 2, 3),
-- (1, 160, 310, 1, 3, 3234, 1, 1, 115, 2018, 0, 'Mock car 27', false, true, true, false, 115000, 'VIN00027', ARRAY['01234567890'], 22000, false, 6, 3, 3),
-- (2, 165, 310, 2, 1, 3232, 2, 2, 120, 2021, 1, 'Mock car 28', true, false, false, false, 120000, 'VIN00028', ARRAY['0111222222'], 37000, true, 7, 1, 3),
-- (1, 167, 310, 3, 2, 3233, 1, 1, 125, 2022, 0, 'Mock car 29', false, true, true, true, 125000, 'VIN00029', ARRAY['01234567890'], 29000, false, 1, 2, 3),
-- (2, 168, 310, 4, 3, 3234, 2, 2, 130, 2023, 1, 'Mock car 30', true, false, false, false, 130000, 'VIN00030', ARRAY['0111222222'], 38000, true, 2, 3, 1);





-- insert into images (vehicle_id, image) values (1, '/images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (1, '/images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (1, '/images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (1, '/images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (2, '/images/cars/2/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (2, '/images/cars/2/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (2, '/images/cars/2/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (2, '/images/cars/2/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (3, '/images/cars/3/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (3, '/images/cars/3/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (3, '/images/cars/3/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (3, '/images/cars/3/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (4, '/images/cars/4/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (4, '/images/cars/4/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (4, '/images/cars/4/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (4, '/images/cars/4/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (5, '/images/cars/5/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (5, '/images/cars/5/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');    
-- insert into images (vehicle_id, image) values (5, '/images/cars/5/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (5, '/images/cars/5/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (6, '/images/cars/6/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (6, '/images/cars/6/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (6, '/images/cars/6/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (6, '/images/cars/6/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (7, '/images/cars/7/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (7, '/images/cars/7/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (7, '/images/cars/7/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (7, '/images/cars/7/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (8, '/images/cars/8/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (8, '/images/cars/8/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (8, '/images/cars/8/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (8, '/images/cars/8/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (9, '/images/cars/9/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (9, '/images/cars/9/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (9, '/images/cars/9/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (9, '/images/cars/9/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (10, '/images/cars/10/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (10, '/images/cars/10/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (10, '/images/cars/10/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (10, '/images/cars/10/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (11, '/images/cars/11/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (11, '/images/cars/11/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (11, '/images/cars/11/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (11, '/images/cars/11/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (12, '/images/cars/12/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (12, '/images/cars/12/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (12, '/images/cars/12/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (12, '/images/cars/12/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (13, '/images/cars/13/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (13, '/images/cars/13/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (13, '/images/cars/13/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (13, '/images/cars/13/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (14, '/images/cars/14/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (14, '/images/cars/14/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (14, '/images/cars/14/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (14, '/images/cars/14/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (15, '/images/cars/15/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (15, '/images/cars/15/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (15, '/images/cars/15/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (15, '/images/cars/15/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (16, '/images/cars/16/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (16, '/images/cars/16/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (16, '/images/cars/16/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (16, '/images/cars/16/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (17, '/images/cars/17/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (17, '/images/cars/17/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (17, '/images/cars/17/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (17, '/images/cars/17/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (18, '/images/cars/18/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (18, '/images/cars/18/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (18, '/images/cars/18/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (18, '/images/cars/18/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (19, '/images/cars/19/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (19, '/images/cars/19/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (19, '/images/cars/19/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (19, '/images/cars/19/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (20, '/images/cars/20/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (20, '/images/cars/20/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (20, '/images/cars/20/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (20, '/images/cars/20/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (21, '/images/cars/21/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (21, '/images/cars/21/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (21, '/images/cars/21/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (21, '/images/cars/21/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (22, '/images/cars/22/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (22, '/images/cars/22/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (22, '/images/cars/22/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (22, '/images/cars/22/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (23, '/images/cars/23/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (23, '/images/cars/23/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (23, '/images/cars/23/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (23, '/images/cars/23/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (24, '/images/cars/24/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (24, '/images/cars/24/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (24, '/images/cars/24/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (24, '/images/cars/24/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (25, '/images/cars/25/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (25, '/images/cars/25/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (25, '/images/cars/25/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (25, '/images/cars/25/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (26, '/images/cars/26/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (26, '/images/cars/26/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (26, '/images/cars/26/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (26, '/images/cars/26/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (27, '/images/cars/27/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (27, '/images/cars/27/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (27, '/images/cars/27/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (27, '/images/cars/27/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (28, '/images/cars/28/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (28, '/images/cars/28/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (28, '/images/cars/28/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (28, '/images/cars/28/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (29, '/images/cars/29/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (29, '/images/cars/29/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (29, '/images/cars/29/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (29, '/images/cars/29/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- insert into images (vehicle_id, image) values (30, '/images/cars/30/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
-- insert into images (vehicle_id, image) values (30, '/images/cars/30/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
-- insert into images (vehicle_id, image) values (30, '/images/cars/30/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
-- insert into images (vehicle_id, image) values (30, '/images/cars/30/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');



-- update brands set model_count = 3;
