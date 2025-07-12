create user da with password '1234';
grant all privileges on database da to da;
grant all privileges on schema public to da;
grant all privileges on all tables in schema public to da;
grant all privileges on all sequences in schema public to da;
alter default privileges in schema public grant all on tables to da;
alter default privileges in schema public grant all on sequences to da;


drop table if exists images;
drop table if exists vehicles;
drop table if exists profiles;
drop table if exists generation_modifications;
drop table if exists colors;
drop table if exists services;
drop table if exists service_types;
drop table if exists regions;
drop table if exists cities;
drop table if exists fuel_types;
drop table if exists drivetrains;
drop table if exists engines;
drop table if exists transmissions;
drop table if exists generations;
drop table if exists body_types;
drop table if exists models;
drop table if exists brands;
drop table if exists users;
drop table if exists admins;
drop table if exists ownership_types;


create table users (
    "id" serial primary key,
    "username" varchar(100) not null,
    "email" varchar(100),
    "role_id" int not null default 1,
    "password" varchar(100) not null,
    "phone" varchar(100),
    "updated_at" timestamp default now(),
    "created_at" timestamp default now(),
    unique("email"),
    unique("phone")
);

insert into users (username, email, password, phone, created_at) 
    values ('user', 'user@gmail.com', '$2a$10$Cya9x0xSJSnRknBmJpW.Bu8ukZpVTqzwgrQgAYNPXdrX2HYGRk33W', '01234567890', now()); -- password: 12345678

insert into users (username, email, password, phone, created_at) 
    values ('user2', 'user2@gmail.com', '$2a$10$Cya9x0xSJSnRknBmJpW.Bu8ukZpVTqzwgrQgAYNPXdrX2HYGRk33W', '0111222222', now()); -- password: 12345678

create table profiles (
    "id" serial primary key, 
    "user_id" int not null,
    "notification" boolean default false,
    "last_active_date" timestamp default now(),
    "created_at" timestamp default now(),
    constraint profiles_user_id_fk 
    foreign key (user_id) references users(id) on delete cascade on update cascade
);

create table admins (
    "id" serial primary key,
    "username" varchar(255) not null,
    "email" varchar(255) not null,
    "password" varchar(255) not null,
    "last_active_date" timestamp default now(),
    "created_at" timestamp default now()
);

insert into admins (username, email, password, last_active_date, created_at) 
    values ('admin', 'admin@gmail.com', '$2a$10$wPb6//DXtLxpDjZgzVEMuOlqqHUtVmPQMbOhmBlkNlAkzzve..CZe', now(), now()); -- password: admin

create table brands (
    "id" serial primary key,
    "name" varchar(255) not null,
    "logo" varchar(255) not null,
    "car_count" int not null default 0,
    "popular" boolean default false,
    "updated_at" timestamp default now()
);

insert into brands (name, logo, car_count, popular) values ('Toyota', '/images/logo/toyota.png', 12, true);
insert into brands (name, logo, car_count, popular) values ('Honda', '/images/logo/honda.png', 8, false);



create table models (
    "id" serial primary key,
    "name" varchar(255) not null,
    "brand_id" int not null,
    "popular" boolean default false,
    "car_count" int not null default 0,
    "updated_at" timestamp default now(),
    constraint models_brand_id_fk 
        foreign key (brand_id) 
            references brands(id)
                on delete cascade
                on update cascade
);

-- toyota
insert into models (name, brand_id, popular, car_count) values ('Camry', 1, true, 7);
insert into models (name, brand_id, popular, car_count) values ('Corolla', 1, true, 41);
insert into models (name, brand_id, popular, car_count) values ('Rav4', 1, false, 73);
insert into models (name, brand_id, popular, car_count) values ('Land Cruiser', 1, false, 1);

-- honda
insert into models (name, brand_id, popular, car_count) values ('Civic', 2, true, 34);
insert into models (name, brand_id, popular, car_count) values ('Accord', 2, false, 23);
insert into models (name, brand_id, popular, car_count) values ('CR-V', 2, false, 65);



create table body_types (
    "id" serial primary key,
    "name" varchar(255) not null,
    "image" character varying(255) not null,
    "created_at" timestamp default now()
);

insert into body_types (name, image) values ('Sedan','/images/body/sedan.png');
insert into body_types (name, image) values ('Hatchback', '/images/body/hatchback.png');
insert into body_types (name, image) values ('SUV', '/images/body/suv.png');
insert into body_types (name, image) values ('Crossover', '/images/body/crossover.png');
insert into body_types (name, image) values ('Coupe', '/images/body/coupe.png');
insert into body_types (name, image) values ('Convertible', '/images/body/convertible.png');
insert into body_types (name, image) values ('Wagon', '/images/body/wagon.png');
insert into body_types (name, image) values ('Pickup Truck', '/images/body/pickup_truck.png');
insert into body_types (name, image) values ('Van', '/images/body/van.png');
insert into body_types (name, image) values ('Minivan', '/images/body/minivan.png');
insert into body_types (name, image) values ('Roadster', '/images/body/roadster.png');
insert into body_types (name, image) values ('Sports Car', '/images/body/sports_car.png');
insert into body_types (name, image) values ('Off-Road', '/images/body/off_road.png');
insert into body_types (name, image) values ('Limousine', '/images/body/limousine.png'); 
insert into body_types (name, image) values ('Utility', '/images/body/utility.png');




create table transmissions (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now()
);

insert into transmissions (name) values ('Automatic');
insert into transmissions (name) values ('Manual');
insert into transmissions (name) values ('Semi-Automatic');


create table engines (
    "id" serial primary key,
    "value" varchar(255) not null,
    "created_at" timestamp default now()
);

insert into engines (value) values ('1.0L');
insert into engines (value) values ('1.5L');
insert into engines (value) values ('2.0L');
insert into engines (value) values ('2.5L');
insert into engines (value) values ('3.0L');
insert into engines (value) values ('4.0L');
insert into engines (value) values ('5.0L');
insert into engines (value) values ('6.0L');
insert into engines (value) values ('7.0L');


create table drivetrains (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now()
);

insert into drivetrains (name) values ('Front-Wheel Drive');
insert into drivetrains (name) values ('Rear-Wheel Drive');
insert into drivetrains (name) values ('All-Wheel Drive');


create table fuel_types (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now()
);

insert into fuel_types (name) values ('Gasoline');
insert into fuel_types (name) values ('Diesel');
insert into fuel_types (name) values ('Electric');
insert into fuel_types (name) values ('Hybrid');

create table cities (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now()
);

insert into cities (name) values ('Dubai');
insert into cities (name) values ('Abu Dhabi');
insert into cities (name) values ('Sharjah');

create table regions (
    "id" serial primary key,
    "name" varchar(255) not null,
    "city_id" int not null,
    "created_at" timestamp default now(),
    constraint regions_city_id_fk
        foreign key (city_id)
            references cities(id)
                on delete cascade
                on update cascade
);

-- dubai
insert into regions (name, city_id) values ('Dubai Marina', 1);
insert into regions (name, city_id) values ('Dubai Mall', 1);
insert into regions (name, city_id) values ('Dubai Creek', 1);
insert into regions (name, city_id) values ('Dubai Creek Harbour', 1);
insert into regions (name, city_id) values ('Dubai Creek Golf Club', 1);

-- abu dhabi
insert into regions (name, city_id) values ('Abu Dhabi Marina', 2);
insert into regions (name, city_id) values ('Abu Dhabi Mall', 2);
insert into regions (name, city_id) values ('Abu Dhabi Creek', 2);
insert into regions (name, city_id) values ('Abu Dhabi Creek Harbour', 2);
insert into regions (name, city_id) values ('Abu Dhabi Creek Golf Club', 2);

-- sharjah
insert into regions (name, city_id) values ('Sharjah Marina', 3);
insert into regions (name, city_id) values ('Sharjah Mall', 3);
insert into regions (name, city_id) values ('Sharjah Creek', 3);

create table service_types (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now()
);

insert into service_types (name) values ('Car Wash');
insert into service_types (name) values ('Car Detailing');
insert into service_types (name) values ('Car Repair');

create table services (
    "id" serial primary key,
    "name" varchar(255) not null,
    "service_type_id" int not null,
    "created_at" timestamp default now(),
    constraint services_service_type_id_fk
        foreign key (service_type_id)
            references service_types(id)
                on delete cascade
                on update cascade
);

-- car wash
insert into services (name, service_type_id) values ('Rocket Wash', 1);
insert into services (name, service_type_id) values ('Premium Wash', 1);
insert into services (name, service_type_id) values ('Express Wash', 1);
insert into services (name, service_type_id) values ('Self-Service Wash', 1);

-- car detailing
insert into services (name, service_type_id) values ('Full Detail', 2);
insert into services (name, service_type_id) values ('Basic Detail', 2);
insert into services (name, service_type_id) values ('Premium Detail', 2);
insert into services (name, service_type_id) values ('Express Detail', 2);
insert into services (name, service_type_id) values ('Self-Service Detail', 2);

-- car repair
insert into services (name, service_type_id) values ('Oil Change', 3);
insert into services (name, service_type_id) values ('Brake Repair', 3);
insert into services (name, service_type_id) values ('Tire Repair', 3);
insert into services (name, service_type_id) values ('Engine Repair', 3);
insert into services (name, service_type_id) values ('Transmission Repair', 3);
insert into services (name, service_type_id) values ('Suspension Repair', 3);
insert into services (name, service_type_id) values ('Electrical Repair', 3);



create table generations (
    "id" serial primary key,
    "name" varchar(255) not null,
    "model_id" int not null,
    "start_year" int not null,
    "end_year" int not null,
    "wheel" boolean not null default true,
    "image" varchar(255) not null,
    "body_type_id" int not null,
    "created_at" timestamp default now(),
    constraint generations_model_id_fk
        foreign key (model_id)
            references models(id)
                on delete cascade,
    constraint generations_body_type_id_fk
        foreign key (body_type_id)
            references body_types(id)
                on delete cascade
);

insert into generations (
    name, model_id, start_year, end_year, image, body_type_id, wheel
) values (
    '1 generation', 1, 2010, 2020, '/images/gens/1.jpg', 1, true
);

insert into generations (
    name, model_id, start_year, end_year, image, body_type_id, wheel
) values (
    '2 generation', 1, 2005, 2025, '/images/gens/2.jpg', 2, true
);


insert into generations (
    name, model_id, start_year, end_year, image, body_type_id, wheel
) values (
    '2 generation', 1, 2005, 2025, '/images/gens/2.jpg', 2, false
);

insert into generations (
    name, model_id, start_year, end_year, image, body_type_id, wheel
) values (
    '3 generation', 2, 2020, 2025, '/images/gens/3.jpg', 2, true
);

insert into generations (
    name, model_id, start_year, end_year, image, body_type_id, wheel
) values (
    '3 generation', 2, 2020, 2025, '/images/gens/3.jpg', 2, false
);


create table generation_modifications (
    "id" serial primary key,
    "generation_id" int not null,
    "engine_id" int not null,
    "fuel_type_id" int not null, 
    "drivetrain_id" int not null,
    "transmission_id" int not null, 
    constraint generation_modifications_generation_id_fk
        foreign key (generation_id)
            references generations(id)
                on delete cascade
                on update cascade,
    constraint generation_modifications_engine_id_fk
        foreign key (engine_id)
            references engines(id)
                on delete cascade
                on update cascade,
    constraint generation_modifications_fuel_type_id_fk
        foreign key (fuel_type_id)
            references fuel_types(id)
                on delete cascade
                on update cascade,
    constraint generation_modifications_drivetrain_id_fk
        foreign key (drivetrain_id)
            references drivetrains(id)
                on delete cascade
                on update cascade,
    constraint generation_modifications_transmission_id_fk
        foreign key (transmission_id)
            references transmissions(id)
                on delete cascade
                on update cascade
);

insert into generation_modifications (
    generation_id, engine_id, fuel_type_id, drivetrain_id, transmission_id
) 
values 
    (1, 1, 1, 1, 1),
    (1, 3, 2, 2, 2),
    (2, 1, 1, 1, 1),
    (2, 3, 2, 2, 2),
    (1, 2, 2, 2, 2);


create table ownership_types (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now()
);

insert into ownership_types (name) values ('Dealership');
insert into ownership_types (name) values ('Private Owner');

create table colors (
    "id" serial primary key,
    "name" varchar(255) not null,
    "hex_code" varchar(7) not null,
    "created_at" timestamp default now()
);

insert into colors (name, hex_code) values ('Black', '#000000');
insert into colors (name, hex_code) values ('White', '#FFFFFF');
insert into colors (name, hex_code) values ('Red', '#FF0000');
insert into colors (name, hex_code) values ('Blue', '#0000FF');
insert into colors (name, hex_code) values ('Green', '#008000');
insert into colors (name, hex_code) values ('Yellow', '#FFFF00');
insert into colors (name, hex_code) values ('Silver', '#C0C0C0');
insert into colors (name, hex_code) values ('Gray', '#808080');
insert into colors (name, hex_code) values ('Orange', '#FFA500');
insert into colors (name, hex_code) values ('Purple', '#800080');


create table vehicles (
    "id" serial primary key,
    "user_id" int,
    "brand_id" int,
    "region_id" int,
    "city_id" int default 1,
    "model_id" int,
    "transmission_id" int,
    "engine_id" int,
    "drivetrain_id" int,
    "body_type_id" int,
    "fuel_type_id" int,
    "ownership_type_id" int not null default 1,
    "owners" int not null default 1,
    "view_count" int not null default 0,
    "year" int not null,
    "description" text,
    "exchange" boolean not null default false,
    "credit" boolean not null default false,
    "wheel" boolean not null default true,
    "crash" boolean not null default false,
    "odometer" int not null default 0,
    "vin_code" varchar(255) not null,
    "phone_numbers" varchar(255)[] not null,
    "price" int not null,
    "new" boolean not null default false,
    "color_id" int not null,
    "trade_in" int not null default 1, -- 1-v nalichi, 2-v put, 3-pod zakaz
    -- "generation_id" int,
    -- "announcement_type" int not null default 0,
    -- "modification_id" int,
    -- "door_count" int,
    -- "interior_color_id" int,
    -- "mileage_km" int,
    -- "negotiable" boolean not null default false,
    -- "credit_price" int,
    "status" int not null default 2, -- 1-pending, 2-not sale (my cars), 3-on sale,
    "updated_at" timestamp default now(),
    "created_at" timestamp default now(),
    constraint vehicles_color_id_fk
        foreign key (color_id)
            references colors(id)
                on delete set null
                on update cascade,
    constraint vehicles_ownership_type_id_fk
        foreign key (ownership_type_id)
            references ownership_types(id)
                on delete cascade
                on update cascade,
    constraint vehicles_user_id_fk
        foreign key (user_id)
            references users(id)
                on delete cascade
                on update cascade,
    constraint vehicles_brand_id_fk
        foreign key (brand_id)
            references brands(id)
                on delete cascade
                on update cascade,
    constraint vehicles_model_id_fk
        foreign key (model_id)
            references models(id)
                on delete cascade
                on update cascade,
    constraint vehicles_transmission_id_fk
        foreign key (transmission_id)
            references transmissions(id)
                on delete cascade
                on update cascade,
    constraint vehicles_engine_id_fk
        foreign key (engine_id)
            references engines(id)
                on delete cascade
                on update cascade,
    constraint vehicles_drivetrain_id_fk
        foreign key (drivetrain_id)
            references drivetrains(id)
                on delete cascade
                on update cascade,
    constraint vehicles_body_type_id_fk
        foreign key (body_type_id)
            references body_types(id)
                on delete cascade
                on update cascade,
    constraint vehicles_fuel_type_id_fk
        foreign key (fuel_type_id)
            references fuel_types(id)
                on delete cascade
                on update cascade,
    constraint vehicles_region_id_fk
        foreign key (region_id)
            references regions(id)
                on delete cascade
                on update cascade,
    constraint vehicles_city_id_fk
        foreign key (city_id)
            references cities(id)
                on delete cascade
                on update cascade
);

insert into vehicles (
    user_id, brand_id, city_id, model_id, body_type_id, fuel_type_id, ownership_type_id, 
    year, exchange, credit, wheel, odometer, vin_code, phone_numbers, price, new,
    status, crash, drivetrain_id, trade_in,
    color_id, view_count, region_id, transmission_id, engine_id
)  values 
(
    1, 1, 1, 1, 1, 1, 1, 
    2020, false, true, false, 74839782, '238748927', ARRAY['23487827397'], 39999, true, 
    3, true, 2, 2,
    1, 392, 1, 1, 1
),
(
   1, 1, 1, 2, 2, 2, 1, 
   2019, true, false, true, 273854, '26873900987', ARRAY['23748798273942'], 7828973, true,
    3, true, 3, 1,
    2, 234, 2, 2, 2
),
(
    2, 1, 1, 3, 2, 2, 1, 
    2010, false, false, false, 8373, '98987987987', ARRAY['89877683783'], 982739488, true,
    3, true, 3, 2,
    2, 487, 3, 3, 3
),
(
    1, 1, 1, 1, 2, 2, 1, 
    2006, true, false, false, 2784, '98987987987', ARRAY['89877683783'], 982739488, true,
    3, true, 3, 1,
    2, 438, 4, 1, 4
),
(
    2, 2, 1, 1, 2, 2, 1, 
    2022, false, true, false, 2739, '98987987987', ARRAY['89877683783'], 982739488, true,
    3, true, 3, 1,
    2, 3, 5, 2, 5
),
(
    1, 2, 1, 2, 2, 2, 1, 
    2021, true, false, true, 283847, '0987654321', ARRAY['01234567890'], 38742973, true,
    3, true, 3, 1,
    2, 45, 6, 3, 6
),
(
    1, 1, 1, 1, 3, 3, 1, 
    2022, true, false, false, 2837959, '1122334455', ARRAY['01234567890'], 3485, true,
    2, true, 1, 1,
    1, 12, 7, 1, 7
),
(
    2, 1, 1, 2, 3, 3, 1, 
    2022, false, false, false, 2348859, '1122334455', ARRAY['01234567890'], 1289397, true,
    2, true, 1, 2,
    3, 8, 8, 2, 8
),
(
    2, 1, 1, 2, 3, 3, 1, 
    2022, false, true, false, 234788, '1122334455', ARRAY['01234567890'], 23487, true,
    3, true, 1, 2,
    3, 97, 9, 3, 9
);
-- 1-pending, 2-not sale (my cars), 3-on sale,


create table images (
    "vehicle_id" int not null,
    "image" varchar(255) not null,
    constraint images_vehicle_id_fk
        foreign key (vehicle_id)
            references vehicles(id)
                on delete cascade
                on update cascade
);

insert into images (vehicle_id, image) values (1, '/images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (1, '/images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (1, '/images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (1, '/images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (2, '/images/cars/2/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (2, '/images/cars/2/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (2, '/images/cars/2/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (2, '/images/cars/2/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (3, '/images/cars/3/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (3, '/images/cars/3/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (3, '/images/cars/3/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (3, '/images/cars/3/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (4, '/images/cars/4/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (4, '/images/cars/4/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (4, '/images/cars/4/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (4, '/images/cars/4/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (5, '/images/cars/5/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (5, '/images/cars/5/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');    
insert into images (vehicle_id, image) values (5, '/images/cars/5/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (5, '/images/cars/5/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (6, '/images/cars/6/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (6, '/images/cars/6/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (6, '/images/cars/6/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (6, '/images/cars/6/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (7, '/images/cars/7/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (7, '/images/cars/7/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (7, '/images/cars/7/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (7, '/images/cars/7/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (8, '/images/cars/8/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (8, '/images/cars/8/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (8, '/images/cars/8/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (8, '/images/cars/8/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (9, '/images/cars/9/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (9, '/images/cars/9/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (9, '/images/cars/9/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (9, '/images/cars/9/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');
