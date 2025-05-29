drop table if exists vehicles;

CREATE TABLE vehicles (
    Make TEXT,
    Model TEXT,
    Annual_Petroleum_Consumption_FT1 NUMERIC,
    Annual_Petroleum_Consumption_FT2 NUMERIC,
    Time_to_charge_120V NUMERIC,
    Time_to_charge_240V NUMERIC,
    City_MPG_FT1 INTEGER,
    Unrounded_City_MPG_FT1 NUMERIC,
    City_MPG_FT2 INTEGER,
    Unrounded_City_MPG_FT2 NUMERIC,
    City_gasoline_consumption NUMERIC,
    City_electricity_consumption NUMERIC,
    EPA_city_utility_factor NUMERIC,
    CO2_FT1 NUMERIC,
    CO2_FT2 NUMERIC,
    CO2_Tailpipe_FT2 NUMERIC,
    CO2_Tailpipe_FT1 NUMERIC,
    Combined_MPG_FT1 INTEGER,
    Unrounded_Combined_MPG_FT1 NUMERIC,
    Combined_MPG_FT2 INTEGER,
    Unrounded_Combined_MPG_FT2 NUMERIC,
    Combined_electricity_consumption NUMERIC,
    Combined_gasoline_consumption NUMERIC,
    EPA_combined_utility_factor NUMERIC,
    Cylinders INTEGER,
    Engine_displacement NUMERIC,
    Drive TEXT,
    EPA_model_type_index INTEGER,
    Engine_descriptor TEXT,
    EPA_Fuel_Economy_Score INTEGER,
    Annual_Fuel_Cost_FT1 INTEGER,
    Annual_Fuel_Cost_FT2 INTEGER,
    Fuel_Type TEXT,
    Fuel_Type1 TEXT,
    GHG_Score INTEGER,
    GHG_Score_Alt_Fuel INTEGER,
    Highway_MPG_FT1 INTEGER,
    Unrounded_Highway_MPG_FT1 NUMERIC,
    Highway_MPG_FT2 INTEGER,
    Unrounded_Highway_MPG_FT2 NUMERIC,
    Highway_gasoline_consumption NUMERIC,
    Highway_electricity_consumption NUMERIC,
    EPA_highway_utility_factor NUMERIC,
    Hatchback_luggage_volume NUMERIC,
    Hatchback_passenger_volume NUMERIC,
    ID INTEGER,
    Two_door_luggage_volume NUMERIC,
    Four_door_luggage_volume NUMERIC,
    MPG_Data TEXT,
    PHEV_Blended BOOLEAN,
    Two_door_passenger_volume NUMERIC,
    Four_door_passenger_volume NUMERIC,
    Range_FT1 NUMERIC,
    Range_City_FT1 NUMERIC,
    Range_City_FT2 NUMERIC,
    Range_Highway_FT1 NUMERIC,
    Range_Highway_FT2 NUMERIC,
    Transmission TEXT,
    Unadjusted_City_MPG_FT1 NUMERIC,
    Unadjusted_City_MPG_FT2 NUMERIC,
    Unadjusted_Highway_MPG_FT1 NUMERIC,
    Unadjusted_Highway_MPG_FT2 NUMERIC,
    Vehicle_Size_Class TEXT,
    Year INTEGER,
    You_Save_Spend INTEGER,
    Guzzler TEXT,
    Transmission_descriptor TEXT,
    T_Charger TEXT,
    S_Charger TEXT,
    ATV_Type TEXT,
    Fuel_Type2 TEXT,
    EPA_Range_FT2 TEXT,
    Electric_motor TEXT,
    MFR_Code TEXT,
    c240Dscr TEXT,
    charge240b TEXT,
    C240B_Dscr TEXT,
    Created_On DATE,
    Modified_On DATE,
    Start_Stop BOOLEAN,
    PHEV_City NUMERIC,
    PHEV_Highway NUMERIC,
    PHEV_Combined integer,
    baseModel TEXT
);
COPY vehicles FROM '/tmp/vehicles.csv' DELIMITER ';' CSV HEADER;


--  new tables

-- 1. Engine Table
CREATE TABLE engine (
    engine_id SERIAL PRIMARY KEY,
    cylinders INTEGER,
    displacement NUMERIC,
    engine_descriptor TEXT
);

-- 2. Fuel Table
CREATE TABLE fuel (
    fuel_id SERIAL PRIMARY KEY,
    fuel_type TEXT,
    fuel_type1 TEXT,
    fuel_type2 TEXT
);

-- 3. Transmission Table
CREATE TABLE transmission (
    transmission_id SERIAL PRIMARY KEY,
    transmission TEXT,
    transmission_descriptor TEXT,
    t_charger TEXT,
    s_charger TEXT
);

-- 4. Drive Table
CREATE TABLE drive (
    drive_id SERIAL PRIMARY KEY,
    drive TEXT
);

-- 5. Size Class Table
CREATE TABLE vehicle_size_class (
    size_class_id SERIAL PRIMARY KEY,
    vehicle_size_class TEXT
);

-- 6. ATV Type Table
CREATE TABLE atv_type (
    atv_type_id SERIAL PRIMARY KEY,
    atv_type TEXT
);

-- 7. Electric Motor Table
CREATE TABLE electric_motor (
    electric_motor_id SERIAL PRIMARY KEY,
    electric_motor TEXT
);


-- ENGINE
INSERT INTO engine (cylinders, displacement, engine_descriptor)
SELECT DISTINCT cylinders, engine_displacement, engine_descriptor
FROM vehicles
WHERE cylinders IS NOT NULL OR engine_displacement IS NOT NULL OR engine_descriptor IS NOT NULL;

-- FUEL
INSERT INTO fuel (fuel_type, fuel_type1, fuel_type2)
SELECT DISTINCT fuel_type, fuel_type1, fuel_type2
FROM vehicles
WHERE fuel_type IS NOT NULL OR fuel_type1 IS NOT NULL OR fuel_type2 IS NOT NULL;

-- TRANSMISSION
INSERT INTO transmission (transmission, transmission_descriptor, t_charger, s_charger)
SELECT DISTINCT transmission, transmission_descriptor, t_charger, s_charger
FROM vehicles
WHERE transmission IS NOT NULL;

-- DRIVE
INSERT INTO drive (drive)
SELECT DISTINCT drive
FROM vehicles
WHERE drive IS NOT NULL;

-- SIZE CLASS
INSERT INTO vehicle_size_class (vehicle_size_class)
SELECT DISTINCT vehicle_size_class
FROM vehicles
WHERE vehicle_size_class IS NOT NULL;

-- ATV TYPE
INSERT INTO atv_type (atv_type)
SELECT DISTINCT atv_type
FROM vehicles
WHERE atv_type IS NOT NULL;

-- ELECTRIC MOTOR
INSERT INTO electric_motor (electric_motor)
SELECT DISTINCT electric_motor
FROM vehicles
WHERE electric_motor IS NOT NULL;

ALTER TABLE vehicles ADD COLUMN engine_id INTEGER;
ALTER TABLE vehicles ADD COLUMN fuel_id INTEGER;
ALTER TABLE vehicles ADD COLUMN transmission_id INTEGER;
ALTER TABLE vehicles ADD COLUMN drive_id INTEGER;
ALTER TABLE vehicles ADD COLUMN size_class_id INTEGER;
ALTER TABLE vehicles ADD COLUMN atv_type_id INTEGER;
ALTER TABLE vehicles ADD COLUMN electric_motor_id INTEGER;

-- ENGINE
UPDATE vehicles v
SET engine_id = e.engine_id
FROM engine e
WHERE v.cylinders = e.cylinders
  AND v.engine_displacement = e.displacement
  AND v.engine_descriptor = e.engine_descriptor;

-- FUEL
UPDATE vehicles v
SET fuel_id = f.fuel_id
FROM fuel f
WHERE v.fuel_type = f.fuel_type
  AND v.fuel_type1 = f.fuel_type1
  AND v.fuel_type2 = f.fuel_type2;

-- TRANSMISSION
UPDATE vehicles v
SET transmission_id = t.transmission_id
FROM transmission t
WHERE v.transmission = t.transmission
  AND v.transmission_descriptor = t.transmission_descriptor
  AND v.t_charger = t.t_charger
  AND v.s_charger = t.s_charger;

-- DRIVE
UPDATE vehicles v
SET drive_id = d.drive_id
FROM drive d
WHERE v.drive = d.drive;

-- SIZE CLASS
UPDATE vehicles v
SET size_class_id = s.size_class_id
FROM vehicle_size_class s
WHERE v.vehicle_size_class = s.vehicle_size_class;

-- ATV TYPE
UPDATE vehicles v
SET atv_type_id = a.atv_type_id
FROM atv_type a
WHERE v.atv_type = a.atv_type;

-- ELECTRIC MOTOR
UPDATE vehicles v
SET electric_motor_id = e.electric_motor_id
FROM electric_motor e
WHERE v.electric_motor = e.electric_motor;

ALTER TABLE vehicles ADD FOREIGN KEY (engine_id) REFERENCES engine(engine_id);
ALTER TABLE vehicles ADD FOREIGN KEY (fuel_id) REFERENCES fuel(fuel_id);
ALTER TABLE vehicles ADD FOREIGN KEY (transmission_id) REFERENCES transmission(transmission_id);
ALTER TABLE vehicles ADD FOREIGN KEY (drive_id) REFERENCES drive(drive_id);
ALTER TABLE vehicles ADD FOREIGN KEY (size_class_id) REFERENCES vehicle_size_class(size_class_id);
ALTER TABLE vehicles ADD FOREIGN KEY (atv_type_id) REFERENCES atv_type(atv_type_id);
ALTER TABLE vehicles ADD FOREIGN KEY (electric_motor_id) REFERENCES electric_motor(electric_motor_id);


ALTER TABLE vehicles
    DROP COLUMN electric_motor,
    DROP COLUMN atv_type,
    DROP COLUMN vehicle_size_class,
    DROP COLUMN drive,
    DROP COLUMN fuel_type,
    DROP COLUMN fuel_type1,
    DROP COLUMN fuel_type2,
    DROP COLUMN transmission,
    DROP COLUMN transmission_descriptor,
    DROP COLUMN t_charger,
    DROP COLUMN s_charger,
    DROP COLUMN cylinders,
    DROP COLUMN engine_displacement,
    DROP COLUMN engine_descriptor;


--  the vehicle table after script
CREATE TABLE vehicles (
    Make TEXT,
    Model TEXT,
    Annual_Petroleum_Consumption_FT1 NUMERIC,
    Annual_Petroleum_Consumption_FT2 NUMERIC,
    Time_to_charge_120V NUMERIC,
    Time_to_charge_240V NUMERIC,
    City_MPG_FT1 INTEGER,
    Unrounded_City_MPG_FT1 NUMERIC,
    City_MPG_FT2 INTEGER,
    Unrounded_City_MPG_FT2 NUMERIC,
    City_gasoline_consumption NUMERIC,
    City_electricity_consumption NUMERIC,
    EPA_city_utility_factor NUMERIC,
    CO2_FT1 NUMERIC,
    CO2_FT2 NUMERIC,
    CO2_Tailpipe_FT2 NUMERIC,
    CO2_Tailpipe_FT1 NUMERIC,
    Combined_MPG_FT1 INTEGER,
    Unrounded_Combined_MPG_FT1 NUMERIC,
    Combined_MPG_FT2 INTEGER,
    Unrounded_Combined_MPG_FT2 NUMERIC,
    Combined_electricity_consumption NUMERIC,
    Combined_gasoline_consumption NUMERIC,
    EPA_combined_utility_factor NUMERIC,
    EPA_model_type_index INTEGER,
    EPA_Fuel_Economy_Score INTEGER,
    Annual_Fuel_Cost_FT1 INTEGER,
    Annual_Fuel_Cost_FT2 INTEGER,
    GHG_Score INTEGER,
    GHG_Score_Alt_Fuel INTEGER,
    Highway_MPG_FT1 INTEGER,
    Unrounded_Highway_MPG_FT1 NUMERIC,
    Highway_MPG_FT2 INTEGER,
    Unrounded_Highway_MPG_FT2 NUMERIC,
    Highway_gasoline_consumption NUMERIC,
    Highway_electricity_consumption NUMERIC,
    EPA_highway_utility_factor NUMERIC,
    Hatchback_luggage_volume NUMERIC,
    Hatchback_passenger_volume NUMERIC,
    ID INTEGER,
    Two_door_luggage_volume NUMERIC,
    Four_door_luggage_volume NUMERIC,
    MPG_Data TEXT,
    PHEV_Blended BOOLEAN,
    Two_door_passenger_volume NUMERIC,
    Four_door_passenger_volume NUMERIC,
    Range_FT1 NUMERIC,
    Range_City_FT1 NUMERIC,
    Range_City_FT2 NUMERIC,
    Range_Highway_FT1 NUMERIC,
    Range_Highway_FT2 NUMERIC,
    Unadjusted_City_MPG_FT1 NUMERIC,
    Unadjusted_City_MPG_FT2 NUMERIC,
    Unadjusted_Highway_MPG_FT1 NUMERIC,
    Unadjusted_Highway_MPG_FT2 NUMERIC,
    Year INTEGER,
    You_Save_Spend INTEGER,
    Guzzler TEXT,
    EPA_Range_FT2 TEXT,
    MFR_Code TEXT,
    c240Dscr TEXT,
    charge240b TEXT,
    C240B_Dscr TEXT,
    Created_On DATE,
    Modified_On DATE,
    Start_Stop BOOLEAN,
    PHEV_City NUMERIC,
    PHEV_Highway NUMERIC,
    PHEV_Combined integer,
    baseModel TEXT
);